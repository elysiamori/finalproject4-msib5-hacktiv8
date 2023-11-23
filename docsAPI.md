# Shopping Store API [DOCS]
https://finalproject4-kelompok6.up.railway.app
### Users
POST /users/register
- description : Register User
- body :
  ``` json
        {
         "full_name" : "string",
         "email" : "string",
         "password" : "string"
        }
  ```
POST /users/login
- description : Login 
- headers : Authorization (Bearer Token String)
- body :
  ``` json
        {
         "email" : "string",
         "password" : "string"
        }
  ```
PATCH /users/topup
- description : Top up Balance
- body :
  ``` json
        {
         "balance" : "integer",
        }
  ```
### Categories
POST 
  
