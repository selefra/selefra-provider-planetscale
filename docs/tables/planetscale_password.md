# Table: planetscale_password

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| deleted_at | timestamp | X | √ | When the password was deleted. | 
| connection_strings | json | X | √ | Connection strings for the branch. | 
| organization_name | string | X | √ | Name of the organization. | 
| id | string | X | √ | ID of the password. | 
| role | string | X | √ | Role for the password. | 
| created_at | timestamp | X | √ | When the password was created. | 
| database_name | string | X | √ | Name of the database. | 
| branch_name | string | X | √ | Name of the database branch. | 
| name | string | X | √ | Name of the password. | 


