# Table: azure_storage_file_shares

This table shows data for Azure Storage File Shares.

https://learn.microsoft.com/en-us/rest/api/storagerp/file-shares/list?tabs=HTTP#fileshareitem

The primary key for this table is **id**.

## Relations

This table depends on [azure_storage_accounts](azure_storage_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|properties|JSON|
|etag|String|
|id (PK)|String|
|name|String|
|type|String|