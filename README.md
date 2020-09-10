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
# create cluster using kind
kind create cluster --name dev --config kind-config.yaml
# install CRD
make install
# check CRD created
k get crd # returns: websvcs.workloads.k8s.coraiberkleid.xyz
# code controller (first pass), then create and start it:
make run
# Create sample WebSvc
k apply -f config/samples/workloads_v1alpha1_websvc.yaml   # expect output websvc.workloads.k8s.coraiberkleid.xyz/websvc-sample created
# List the new resource
k get websvc
# List all
k get all   # Notice 2 pods for deploymentTier "dev"
# clean up
###k delete deploy websvc-sample # not needed since controller defines websvc as parent to deployment
###k delete websvc websvc-sample # not needed since controller defines websvc as parent to deployment using SetControllerReference
k delete websvc websvc-sample


# Add owner reference to make future cleanups easier
```