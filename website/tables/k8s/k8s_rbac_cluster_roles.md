# Table: k8s_rbac_cluster_roles

This table shows data for Kubernetes (K8s) Role-Based Access Control (RBAC) Cluster Roles.

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|context|String|
|kind|String|
|api_version|String|
|name|String|
|namespace|String|
|uid (PK)|String|
|resource_version|String|
|generation|Int|
|deletion_grace_period_seconds|Int|
|labels|JSON|
|annotations|JSON|
|owner_references|JSON|
|finalizers|StringArray|
|rules|JSON|
|aggregation_rule|JSON|