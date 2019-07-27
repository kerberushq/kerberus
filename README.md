# Kerberus

## Kubernetes Kerberus Operator

### Extended Permission model for Kubernetes

`kerberus` is an operator which abstracts CRD objects via `CRDRequest` allowing you to remove `cluster-admin` from users and at the same time allow the to use `CRD` objects.


### Usage

Create `CRDConfig` object defining `CRDRequest` behaviour:



### Close roadMap

* Ability to create and control RBAC policies and controls them based on `CRDRequest`
* Ability to set Namespace level rolebindings
* Ability to set Cluster level rolebindings
* Ability to update Namespace level and Cluster level roles via `CRDRequest`

### Future RoadMap

* Ability to use AdmissionController to prevent access to system object
