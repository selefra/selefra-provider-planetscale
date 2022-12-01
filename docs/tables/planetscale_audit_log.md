# Table: planetscale_audit_log

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| action | string | X | √ | Short action for this audit record, e.g. created. | 
| remote_ip | string | X | √ | IP address the action was requested from. | 
| target_id | string | X | √ | ID of the resource type for this audit record. | 
| target_display_name | string | X | √ | Display name for the target resoruce, e.g. test_db. | 
| created_at | timestamp | X | √ | When the audit record was created. | 
| actor_type | string | X | √ | Type of the actor, e.g. User. | 
| auditable_type | string | X | √ | Resource type the audit entry is for, e.g. Branch. | 
| auditable_display_name | string | X | √ | Display name of the resource for this audit entry, e.g. test_branch. | 
| actor_display_name | string | X | √ | Display name of the actor. | 
| target_type | string | X | √ | Resource type for this audit record, e.g. Database. | 
| metadata | json | X | √ | Metadata for the audit record. | 
| actor_id | string | X | √ | Unique ID of the actor. | 
| audit_action | string | X | √ | Full action for this audit record, e.g. deploy_request.created. | 
| location | string | X | √ | Geographic location the action was requested from. | 
| updated_at | timestamp | X | √ | When the audit record was updated. | 
| id | string | X | √ | Unique ID of the log entry. | 
| type | string | X | √ | Type of log entry, e.g. AuditLogEvent. | 
| auditable_id | string | X | √ | Unique ID for the resource type of the audit entry. | 


