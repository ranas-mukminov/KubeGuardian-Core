# KubeGuardian-Core

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-1.24+-blue.svg)](https://kubernetes.io/)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org/)

**Runtime Security Engine for Kubernetes**

🛡️ eBPF-based monitoring • 🤖 ML anomaly detection • 🚫 Policy enforcement • 📊 SIEM integration

---

## 🏢 About

**KubeGuardian** is a runtime security engine that provides continuous monitoring and threat protection for Kubernetes workloads. Using eBPF (Extended Berkeley Packet Filter) for kernel-level observability, machine learning for behavioral analysis, and Kubernetes admission webhooks for policy enforcement, KubeGuardian delivers comprehensive security without impacting application performance.

**Built by**: [@ranas-mukminov](https://github.com/ranas-mukminov)  
**Website**: [run-as-daemon.ru](https://run-as-daemon.ru)  
**Language**: [Русская версия](README.ru.md)

> **Note**: This is a showcase repository demonstrating the architecture and concept. The enterprise version with full eBPF instrumentation, production ML models, and 24/7 support is available through security consulting services.

---

## ✨ Key Features

### 🔍 Runtime Monitoring
- **eBPF-based syscall tracing** with sub-microsecond latency
- **Network activity monitoring**: TCP/UDP connections, DNS queries
- **File access tracking**: read/write operations, permission changes
- **Process execution monitoring**: full command-line and environment capture
- **Zero-copy data collection** for minimal overhead

### 🧠 Intelligent Threat Detection
- **ML-based anomaly detection**: learns normal behavior patterns per pod/namespace
- **Behavioral profiling**: statistical and neural network models
- **Multi-factor risk scoring**: aggregate threat assessment
- **Threat intelligence integration**: IP/domain reputation checks
- **Adaptive policies**: dynamic rule adjustment based on observed patterns

### 🚫 Policy Enforcement
- **Kubernetes admission webhooks**: validate pods before deployment
- **Real-time blocking**: prevent malicious actions at runtime
- **Automatic hardening**: mutate pod specs for security best practices
- **Compliance policies**: CIS Benchmarks, PCI-DSS, custom rules
- **Audit logging**: complete forensic trail

### 📊 Observability & Integration
- **Prometheus metrics**: performance and security metrics export
- **OpenTelemetry tracing**: distributed request tracking
- **SIEM integration**: Elasticsearch, Splunk, Loki
- **Alert routing**: PagerDuty, Slack, email, webhooks
- **Custom dashboards**: Grafana visualization templates

---

## 🏗️ Architecture

![KubeGuardian Architecture](docs/architecture.png)

### Components

1. **eBPF Agent** (DaemonSet)
   - Deployed on every Kubernetes node
   - Collects kernel-level events with eBPF/CO-RE
   - Aggregates and forwards to Policy Engine

2. **Policy Engine** (Deployment)
   - Processes events from all agents
   - Runs ML models for anomaly detection
   - Evaluates security policies
   - Exports metrics and alerts

3. **Admission Webhook** (Deployment)
   - Validates pod creation/updates
   - Enforces security constraints
   - Prevents non-compliant workloads

**📖 Detailed documentation**: [ARCHITECTURE.md](ARCHITECTURE.md)

---

## 🚀 Quick Start

### Prerequisites

- Kubernetes 1.24+ cluster
- `kubectl` configured with cluster admin access
- Helm 3.0+ (optional, for simplified deployment)

### Installation

#### 1. Create namespace
```bash
kubectl create namespace kubeguardian-system
```

#### 2. Deploy RBAC resources
```bash
kubectl apply -f deployments/kubernetes/rbac.yaml
```

#### 3. Deploy the agent
```bash
kubectl apply -f deployments/kubernetes/daemonset.yaml
```

#### 4. Deploy admission webhook
```bash
# Generate TLS certificates (requires cert-manager or manual creation)
# See docs/certificates.md for details

kubectl apply -f deployments/kubernetes/webhook.yaml
```

#### 5. Verify deployment
```bash
kubectl -n kubeguardian-system get pods
```

Expected output:
```
NAME                                    READY   STATUS    RESTARTS   AGE
kubeguardian-agent-xxxxx                1/1     Running   0          1m
kubeguardian-webhook-xxxxxxxxxx-xxxxx   1/1     Running   0          1m
```

### Configuration

Edit policies via ConfigMaps or CRDs (Custom Resource Definitions):

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: kubeguardian-policies
  namespace: kubeguardian-system
data:
  policies.yaml: |
    policies:
    - name: block-privilege-escalation
      description: Prevent containers from gaining additional privileges
      conditions:
      - field: securityContext.allowPrivilegeEscalation
        operator: equals
        value: true
      action: deny
```

---

## 📋 Use Cases

### 🎯 Runtime Threat Detection
Detect and block:
- Cryptocurrency miners
- Reverse shells and C2 connections
- Credential theft attempts
- Lateral movement
- Container breakout attempts

### 🔒 Compliance Enforcement
Automated validation against:
- CIS Kubernetes Benchmarks
- PCI-DSS container requirements
- Custom organizational policies
- Industry-specific regulations (HIPAA, GDPR, etc.)

### 🏢 Zero Trust Security
- Continuous verification of all workloads
- Least-privilege access enforcement
- Network segmentation validation
- Immutable infrastructure checks

### 📈 Security Operations
- Centralized security event aggregation
- Forensic investigation support
- Incident response automation
- Security posture dashboards

---

## 🛠️ Technology Stack

| Component | Technology |
|-----------|-----------|
| **Core Language** | Go 1.21+ |
| **Kernel Instrumentation** | eBPF, libbpf, CO-RE |
| **Machine Learning** | TensorFlow Lite, ONNX Runtime |
| **Kubernetes Integration** | client-go, admission webhooks, CRDs |
| **Observability** | Prometheus, OpenTelemetry |
| **Event Storage** | Elasticsearch, Loki |

---

## 📊 Performance

| Metric | Value |
|--------|-------|
| **Agent CPU overhead** | <2% per core |
| **Agent memory** | ~100MB per node |
| **Syscall latency impact** | <1μs |
| **Policy evaluation latency** | <10ms (p99) |
| **Webhook admission latency** | <50ms (p99) |
| **Event throughput** | >10,000 events/sec per controller |

Tested on: GKE, EKS, AKS, bare-metal Kubernetes

---

## 🔐 Security Considerations

- **Minimal privileges**: Agents run with only required capabilities (`CAP_SYS_ADMIN` for eBPF)
- **mTLS communication**: All inter-component traffic encrypted
- **No sensitive data storage**: Events processed in-memory, exports encrypted
- **RBAC integration**: Kubernetes-native access control
- **Audit logging**: Complete trail of all security decisions

---

## 🤝 Contributing

Contributions to documentation and architecture discussions are welcome!

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

**Note**: The core implementation is proprietary. This repository contains concept demonstrations and interfaces.

---

## 📜 License

MIT License - see [LICENSE](LICENSE) file.

**Note**: Enterprise version is available under commercial license.

---

## 💼 Commercial Support

This showcase demonstrates the KubeGuardian architecture. **Enterprise version** includes:

✅ Full eBPF instrumentation with CO-RE support  
✅ Production-hardened ML models  
✅ 24/7 support with SLA  
✅ Custom policy development  
✅ Compliance certifications  
✅ Multi-cluster management  
✅ Advanced threat intelligence  

### Services Available

🔍 **Security Audit**
- Kubernetes infrastructure assessment
- Vulnerability scanning and remediation
- Security posture evaluation
- Compliance gap analysis

⚙️ **DevSecOps Integration**
- CI/CD pipeline security
- Infrastructure as Code scanning
- Secret management
- Automated compliance checks

🛡️ **Monitoring & Incident Response**
- SIEM integration and tuning
- Real-time threat monitoring
- Incident response planning
- Forensic analysis support

📞 **Contact**: [https://run-as-daemon.ru](https://run-as-daemon.ru)  
👤 **GitHub**: [@ranas-mukminov](https://github.com/ranas-mukminov)

---

## 📚 Documentation

- [Architecture Overview](ARCHITECTURE.md)
- [Installation Guide](docs/installation.md) _(coming soon)_
- [Policy Configuration](docs/policies.md) _(coming soon)_
- [Troubleshooting](docs/troubleshooting.md) _(coming soon)_
- [API Reference](docs/api.md) _(coming soon)_

---

## 🗺️ Roadmap

- [x] Core architecture and concept
- [x] eBPF agent interfaces
- [x] Policy engine design
- [ ] Windows container support
- [ ] Service mesh integration (Istio, Linkerd)
- [ ] AI-powered policy generation
- [ ] Compliance reporting automation
- [ ] Multi-cloud deployment automation

---

## 🌟 Related Projects

- **CloudGuardian** - AI-powered cloud infrastructure auditor
- **AutoHarden-Toolkit** - Automated server hardening based on CIS Benchmarks

See more at: [https://github.com/ranas-mukminov](https://github.com/ranas-mukminov)

---

<p align="center">
  <strong>Built with ❤️ for Kubernetes Security</strong><br>
  <a href="https://run-as-daemon.ru">run-as-daemon.ru</a> • 
  <a href="https://github.com/ranas-mukminov">@ranas-mukminov</a>
</p>
