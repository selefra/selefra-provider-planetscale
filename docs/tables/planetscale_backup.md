# Table: planetscale_backup

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ | When the backup was created. | 
| updated_at | timestamp | X | √ | When the backup was updated. | 
| organization_name | string | X | √ | Name of the organization. | 
| database_name | string | X | √ | Name of the database. | 
| branch_name | string | X | √ | Name of the database branch. | 
| name | string | X | √ | Name of the backup. | 
| size | int | X | √ | Size of the backup. | 
| id | string | X | √ | ID of the backup. | 
| state | string | X | √ | State of the backup. | 
| started_at | timestamp | X | √ | When the backup was started. | 
| completed_at | timestamp | X | √ | When the backup was completed. | 
| expires_at | timestamp | X | √ | When the backup expires. | 


