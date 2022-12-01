# Table: planetscale_database_branch

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| database_name | string | X | √ | Name of the database. | 
| parent_branch | string | X | √ | Parent of this branch. | 
| region_slug | string | X | √ | Region where the database is located. | 
| production | bool | X | √ | True if this branch is in production. | 
| created_at | timestamp | X | √ | When the branch was created. | 
| updated_at | timestamp | X | √ | When the branch was updated. | 
| organization_name | string | X | √ | Name of the organization. | 
| name | string | X | √ | Name of the branch. | 
| ready | bool | X | √ | True if the branch is ready. | 
| access_host_url | string | X | √ | Host name to access the database. | 


