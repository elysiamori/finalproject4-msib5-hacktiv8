# Shopping Store API [DOCS]
https://finalproject4-kelompok6.up.railway.app
### Users
##### POST /users/register
Request:
- description : Register User
- body :
  ``` json
        {
         "full_name" : "string",
         "email" : "string",
         "password" : "string"
        }
  ```
#### POST /users/login
Request:
- description : Login 
- body :
  ``` json
        {
         "email" : "string",
         "password" : "string"
        }
  ```
#### PATCH /users/topup
Request:
- description : Top up Balance
- headers : Authorization (Bearer Token String)
- body :
  ``` json
        {
         "balance" : "integer",
        }
  ```
### Categories
#### POST /categories
Request:
- description : Add Category
- headers : Authorization (Bearer Token String)
- body:
  ``` json
        {
         "type" : "string",
        }
  ```
#### GET /categories
Request:
- description : Get Category
- headers : Authorization (Bearer Token String)
- body:
    ``` json
       [
        {
          "id": "integer",
          "type": "string",
          "sold_product_amount": "integer",
          "created_at": "date",
          "updated_at": "date",
          "Products": [
            {
              "id": "integer",
              "title": "string",
              "price": "integer",
              "stock": "integer",
              "created_at": "date",
              "updated_at": "date"
            },
          ]
        }
      ]
  ```
#### PATCH /categories/:categoryID
Request:
- description : Update Category By ID
- headers : Authorization (Bearer Token String)
- body:
  ```json
        {
          "type" : "string"
        }
  ```
#### DELETE /categories/:categoryID
Request:
- description : Delete Category By ID
- headers : Authorization (Bearer Token String)
- params : categoryID(id)

### Products
#### POST /products
Request:
- description : Add Products
- headers : Authorization (Bearer Token String)
- body:
 ```json
        {
          "title" : "string",
          "price" : "integer",
          "stock" : "integer",
          "category_id" : "integer"
        }
  ```
#### GET /products
Request:
- description : Get All Products
- headers : Authorization (Bearer Token String)
- body:
 ```json
        [
            {
                "id": "integer",
                "title": "string",
                "price": "integer",
                "stock": "integer",
                "category_id": "integer",
                "created_at": "date",
            },
        ]
 ```
#### PUT /products/:productID
Request:
- description : Get All Products
- headers : Authorization (Bearer Token String)
- params : productID(integer)
- body:
    ```json
        {
          "title" : "string",
          "price" : "integer",
          "stock" : "integer",
          "category_id" : "integer"
        }
    ```
##### DELETE /products/:productID
Request:
- description : Delete Product
- headers : Authorization (Bearer Token String)
- params : productID(integer)

### Transaction History
#### POST /transactions
Request:
- description : Create Transaction
- headers : Authorization (Bearer Token String)
- body:
    ```json
        {
            "product_id" : "integer",
            "quantity" : "integer"
        }
    ```
#### GET /transactions/my-transactions
Request:
- description : Get My Transaction
- headers : Authorization (Bearer Token String)

#### GET /transactions/user-transactions
Request:
- description : Get All Transactions
- headers : Authorization (Bearer Token String)
