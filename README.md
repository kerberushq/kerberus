# Kerberus

## Kubernetes Kerberus Operator

### Extended Permission model for Kubernetes

`kerberus` is an operator which abstracts CRD objects via `CRDRequest` allowing you to remove `cluster-admin` from users and at the same time allow the to use `CRD` objects. More features are coming soon. Check `Roadmap` below

### Usage

Deploy Kerberus operator to your cluster:
```
kubectl create -f https://raw.githubusercontent.com/kerberushq/kerberus/master/data/namespace.yaml
kubectl create -f https://raw.githubusercontent.com/kerberushq/kerberus/master/data/service_account.yaml
kubectl create -f https://raw.githubusercontent.com/kerberushq/kerberus/master/data/role_binding.yaml
kubectl create -f https://raw.githubusercontent.com/kerberushq/kerberus/master/data/role.yaml
kubectl create -f https://raw.githubusercontent.com/kerberushq/kerberus/master/data/cluster_role.yaml
kubectl create -f https://raw.githubusercontent.com/kerberushq/kerberus/master/data/cluster_role_binding.yaml
kubectl create -f https://raw.githubusercontent.com/kerberushq/kerberus/master/data/operator.yaml
```

Deploy `CRDConfig` object defining `CRDRequest` behaviour:

* NOTE: for now you can have only one object of `CRDConfig` *

```
cat <<EOF | kubectl apply -n kerberus -f -
apiVersion: crd.kerberus.io/v1alpha1
kind: CRDConfig
metadata:
  name: crdconfig-dennyall
spec:
  denyAll: true
  allow:
  - "[a-zA-Z].openshift.io"
  - "[a-zA-Z].k8s.io"
EOF
```

Create example `CRDRequest` object:
```
cat <<EOF | kubectl apply -n kerberus -f -
apiVersion: crd.kerberus.io/v1alpha1
kind: CRDRequest
metadata:
  name: example-crdrequest
spec:
  owned: true
  crdList:
  - apiVersion: apiextensions.k8s.io/v1beta1
    kind: CustomResourceDefinition
    metadata:
      name: crontabs.stable.openshift.io
    spec:
      group: stable.openshift.io
      version: v1 
      scope: Namespaced 
      names:
        plural: crontabs 
        singular: crontab 
        kind: CronTab 
        shortNames:
        - ct 
  - apiVersion: apiextensions.k8s.io/v1beta1
    kind: CustomResourceDefinition
    metadata:
      name: appconfigs.stable.k8s.io
    spec:
      group: stable.k8s.io
      version: v1
      scope: Namespaced
      names:
        plural: appconfigs
        singular: appconfig
        kind: AppConfig
        shortNames:
        - ac
EOF
```

Check `CRD` objects being provisioned into the cluster:
`kubectl get crd`


## Roadmap

* Ability to create and control RBAC policies and controls them based on `CRDRequest` admitted objects.
* Ability to set `new namespace` level rolebindings
* Ability to set Cluster level rolebindings
* Ability to update Namespace level and Cluster level roles via `CRDRequest`
* Ability to use AdmissionController to prevent access to system object
