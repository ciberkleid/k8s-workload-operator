# k8s-workload-operator

General purpose workload operator.

Purpose: provide convenient abstractions and sensible defaults for applications using common patterns when deploying to K8s-based platforms.

domain: k8s.coraiberkleid.xyz

api group: workloads

version: v1alpha1

api type: WebSvc

## Data Model

spec:

- image
- deploymentTier
    'dev' = replicas:2
    'test' = replicas:3
    'prod' = replicas:5
    
## Notes

go mod init github.com/ciberkleid/k8s-workload-operator
kubebuilder init --domain k8s.coraiberkleid.xyz