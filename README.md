# k8s-operator-opt
## Generate Operator by KubeBuilder
### Init operator
```go
kubebuilder init --domain luojun96.io --owner "Jun Luo"
```
### Create operator
```go
kubebuilder create api --group apps --version v1beta1 --kind CustomDeamonset
```
