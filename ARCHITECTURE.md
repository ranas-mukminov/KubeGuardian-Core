# KubeGuardian Architecture

## Overview

KubeGuardian is a runtime security engine for Kubernetes that provides continuous monitoring, threat detection, and policy enforcement at the pod level. The system uses eBPF (Extended Berkeley Packet Filter) for low-overhead kernel-level observability, machine learning for anomaly detection, and Kubernetes admission webhooks for policy enforcement.

## System Architecture

![KubeGuardian Architecture](architecture.png)

## Core Components

### 1. eBPF Monitoring Agent

**Deployment**: DaemonSet (runs on every node)

**Responsibilities**:
- Kernel-level syscall tracing using eBPF/CO-RE (Compile Once, Run Everywhere)
- Network connection monitoring and packet inspection
- File access and process execution tracking
- Container runtime event collection
- Zero-copy data transfer to userspace

**Technology Stack**:
- eBPF programs (C) compiled with LLVM
- libbpf for userspace interaction
- Go for agent orchestration and data aggregation

**Collected Data**:
```
- Syscalls: execve, open, connect, bind, etc.
- Network: TCP/UDP connections, DNS queries
- Files: read/write operations, permission changes
- Processes: execution tree, command lines, environment
```

### 2. Policy Engine & Controller

**Deployment**: Deployment (2+ replicas for HA)

**Responsibilities**:
- Aggregate data from all monitoring agents
- ML-based behavioral analysis and anomaly detection
- Rule evaluation and policy decision making
- Threat intelligence integration
- Metrics export to Prometheus

**ML Components**:
- **Baseline Learning**: Build normal behavior profiles per pod/namespace
- **Anomaly Detection**: Statistical and neural network models
- **Risk Scoring**: Multi-factor threat assessment
- **Adaptive Policies**: Dynamic rule adjustment based on learned patterns

**Policy Types**:
1. **Preventive**: Block actions before execution (via webhook)
2. **Detective**: Alert on suspicious behavior
3. **Audit**: Log all actions for forensics

### 3. Admission Webhook

**Deployment**: Deployment (2+ replicas)

**Responsibilities**:
- Intercept pod creation/update requests
- Validate against security policies
- Mutate pod specs for automatic hardening
- Reject non-compliant workloads

**Validation Checks**:
- Image provenance and signatures
- Security context constraints
- Resource limits enforcement
- Network policy requirements
- Behavioral risk assessment

## Data Flow

### Runtime Monitoring Flow

```
1. Container executes syscall
   ↓
2. eBPF program captures event in kernel
   ↓
3. Event buffered in perf/ring buffer
   ↓
4. Agent reads event in userspace
   ↓
5. Agent sends to Policy Engine via gRPC
   ↓
6. Engine evaluates against policies
   ↓
7. Decision: Allow / Alert / Block
   ↓
8. Export metrics & alerts
```

### Admission Control Flow

```
1. User creates Pod
   ↓
2. K8s API sends AdmissionReview
   ↓
3. Webhook validates against policies
   ↓
4. Query historical behavior from Policy Engine
   ↓
5. Calculate risk score
   ↓
6. Return decision: Allow / Deny / Mutate
```

## Technology Stack

### Core
- **Language**: Go 1.21+
- **eBPF**: libbpf, CO-RE, BTF
- **ML**: TensorFlow Lite / ONNX Runtime

### Kubernetes Integration
- **API**: client-go
- **CRDs**: Custom security policies
- **Webhooks**: Admission controllers

### Observability
- **Metrics**: Prometheus
- **Tracing**: OpenTelemetry
- **Logging**: Structured logs (JSON)

### Storage
- **Time-series**: Prometheus for metrics
- **Events**: Elasticsearch / Loki
- **Policies**: Persistent volume + ConfigMaps

## Security Considerations

### Agent Security
- Runs with `CAP_SYS_ADMIN` for eBPF (minimal required)
- Read-only access to `/sys` filesystem
- mTLS for agent-to-controller communication
- No data persisted on nodes

### Controller Security
- RBAC with least-privilege service accounts
- API authentication via K8s RBAC
- Encrypted storage for sensitive policy data
- Audit logging for all policy decisions

### Webhook Security
- TLS certificates with cert-manager integration
- Failure policy: configurable (fail-open vs fail-closed)
- Timeout protection to prevent cluster lockout
- Version skew handling for API compatibility

## Performance Characteristics

### eBPF Agent
- **CPU**: <2% per core under normal load
- **Memory**: ~100MB per node
- **Latency**: <1μs overhead per syscall

### Policy Engine
- **Throughput**: >10,000 events/second per replica
- **Latency**: <10ms p99 for policy evaluation
- **Memory**: ~500MB per replica

### Admission Webhook
- **Latency**: <50ms p99 for pod validation
- **Availability**: 99.9% with 2+ replicas

## Deployment Architecture

### Namespace: kubeguardian-system

```
Components:
├── DaemonSet: kubeguardian-agent (1 per node)
├── Deployment: kubeguardian-controller (2 replicas)
├── Deployment: kubeguardian-webhook (2 replicas)
├── Service: kubeguardian-controller (internal)
├── Service: kubeguardian-webhook (webhook)
└── ConfigMaps: policies, ML models
```

### Resource Requirements

**Minimum** (3-node cluster):
- CPU: 1.5 cores total
- Memory: 1.5GB total

**Recommended** (production):
- CPU: 3+ cores total
- Memory: 3+ GB total
- Persistent storage: 10GB+ for event retention

## Integration Points

### Upstream (Data Collection)
- Kubernetes API (pod metadata)
- Container runtime (containerd/CRI-O)
- Kernel (eBPF hooks)

### Downstream (Alerting & SIEM)
- **Prometheus**: Metrics export
- **AlertManager**: Alert routing
- **Elasticsearch**: Event forwarding
- **Splunk/SIEM**: Security event export
- **Slack/PagerDuty**: Incident notifications

### External Services
- **Threat Intelligence**: IP/domain reputation
- **Image Registries**: Vulnerability scanning
- **Certificate Authority**: mTLS certificates

## Scalability

### Horizontal Scaling
- **Agents**: Automatic (DaemonSet)
- **Controllers**: Scale based on event rate
- **Webhooks**: Scale for admission latency

### Vertical Scaling
- Controller memory scales with policy complexity
- Agent CPU scales with syscall rate

### Multi-Cluster
- Federated policy management
- Cross-cluster threat correlation
- Centralized dashboard

## Comparison with Alternatives

| Feature | KubeGuardian | Falco | Tetragon | Aqua |
|---------|--------------|-------|----------|------|
| eBPF-based | ✅ | ✅ | ✅ | ❌ |
| ML Anomaly Detection | ✅ | ❌ | ❌ | ✅ |
| Admission Control | ✅ | ❌ | ❌ | ✅ |
| Open Core | ✅ | ✅ | ✅ | ❌ |
| K8s Native | ✅ | Partial | ✅ | ✅ |

## Future Roadmap

### Phase 1 (Q1 2025)
- [ ] CiliumNetworkPolicy integration
- [ ] Istio service mesh support
- [ ] Windows container support

### Phase 2 (Q2 2025)
- [ ] Auto-remediation actions
- [ ] Compliance reporting (CIS, PCI-DSS)
- [ ] Multi-tenancy isolation

### Phase 3 (Q3 2025)
- [ ] AI-powered policy generation
- [ ] Zero-trust network policies
- [ ] Cloud provider integrations (EKS, GKE, AKS)

---

## Commercial Support

This is a showcase/concept version demonstrating the architecture.

**Enterprise Implementation** includes:
- Full eBPF instrumentation with CO-RE support
- Production-ready ML models
- 24/7 support and SLA
- Custom policy development
- Compliance certifications

### Services Available
- Infrastructure security audit
- DevSecOps pipeline integration
- Kubernetes hardening consulting
- On-premise and cloud deployments

**Contact**: [https://run-as-daemon.ru](https://run-as-daemon.ru)  
**Author**: [@ranas-mukminov](https://github.com/ranas-mukminov)
