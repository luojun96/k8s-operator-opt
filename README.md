# k8s-operator-opt
## Generate Operator by KubeBuilder
### Create a kubebuilder project, which requires an empty folder
```sh
kubebuilder init --domain luojun96.io
```
### Check project layout
```sh
cat PROJECT
domain: luojun96.io
layout:
- go.kubebuilder.io/v3
projectName: k8s-operator-opt
repo: github.com/luojun96/k8s-operator-opt
resources:
- api:
    crdVersion: v1
    namespaced: true
  controller: true
  domain: luojun96.io
  group: apps
  kind: CustomDeamonset
  path: github.com/luojun96/k8s-operator-opt/api/v1beta1
  version: v1beta1
version: "3"
```
### Create API, create resource[Y], create controller[Y]
```sh
kubebuilder create api --group apps --version v1beta1 --kind CustomDeamonset
```
### Open project with IDE and edit `api/v1alpha1/simplestatefulset_types.go`

### Finish custom logic
```go
// controllers/customdeamonset_controller.go
func (r *CustomDeamonsetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
  // TO-DO
	return ctrl.Result{}, nil
}
```
### Check Makefile

```makefile
Build targets:
    ### create code skeletion
    manifests: generate crd
    generate: generate api functions, like deepCopy

    ### generate crd and install
    run: Run a controller from your host.
    install: Install CRDs into the K8s cluster specified in ~/.kube/config.

    ### docker build and deploy
    docker-build: Build docker image with the manager.
    docker-push: Push docker image with the manager.
    deploy: Deploy controller to the K8s cluster specified in ~/.kube/config.
```

### Edit `controllers/mydaemonset_controller.go`, add permissions to the controller
```go
//+kubebuilder:rbac:groups=apps.cncamp.io,resources=mydaemonsets/finalizers,verbs=update
// Add the following
//+kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
```

### Generate crd

```sh
make manifests
```

### Build & install

```sh
make build
make docker-build
make docker-push
make deploy
```

## Enable webhooks

### Install cert-manager

```sh
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.11.1/cert-manager.yaml
```

### Create webhooks

```sh
kubebuilder create webhook --group apps --version v1beta1 --kind CustomDaemonset --defaulting --programmatic-validation
```

### Change code

### Enable webhook in `config/default/kustomization.yaml`

### Redeploy
### Enable webhook in `config/default/kustomization.yaml`
