# Table: planetscale_deploy_request

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| branch | string | X | √ | Deploy request branch. | 
| into_branch | string | X | √ | Deploy request into branch. | 
| approved | bool | X | √ | True if the deploy request is approved. | 
| created_at | timestamp | X | √ | When the deploy request was created. | 
| updated_at | timestamp | X | √ | When the deploy request was updated. | 
| organization_name | string | X | √ | Name of the organization. | 
| database_name | string | X | √ | Name of the database. | 
| number | int | X | √ | Number for this deploy request. | 
| id | string | X | √ | Unique ID for the deplloy request. | 
| state | string | X | √ | State of the deploy request. | 
| notes | string | X | √ | Notes for the deploy request. | 
| deployment | json | X | √ | Details of the deployment. | 
| closed_at | timestamp | X | √ | When the deploy request was closed. | 


