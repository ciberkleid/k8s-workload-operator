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

```
# initialize go module
go mod init github.com/ciberkleid/k8s-workload-operator
# use kubebuilder init to initialize project 
kubebuilder init --domain k8s.coraiberkleid.xyz
# create api (say y & y to prompts to create Resource and Controller)
kubebuilder create api --group workloads --version v1alpha1 --kind WebSvc
```