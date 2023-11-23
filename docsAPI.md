# Shopping Store API [DOCS]
https://finalproject4-kelompok6.up.railway.app
### Users
##### [POST] /users/register
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
Response:
- status 201
- data :
```json
        {
        "id": "integer",
        "full_name": "string",
        "email": "string",
        "password": "string",
        "balance": "integer",
        "created_at": "date"
        }
```
#### [POST] /users/login
Request:
- description : Login 
- body :
  ``` json
        {
         "email" : "string",
         "password" : "string"
        }
  ```
Response :
- status 200
- data :
```json
    {
        "token" : "jwt-string",
    }
```
#### [PATCH] /users/topup
Request:
- description : Top up Balance
- headers : Authorization (Bearer Token String)
- body :
  ``` json
        {
         "balance" : "integer",
        }
  ```
Response :
- status 200
- data:
```json
    {
        "message": "Your balance has been successfully updated to Rp <balance>"
    }
```
### Categories
#### [POST] /categories
Request:
- description : Add Category
- headers : Authorization (Bearer Token String)
- body:
  ``` json
        {
         "type" : "string",
        }
  ```
Response :
- status 201
- data:
```json
        {
            "id": "integer",
            "name": "string",
            "created_at": "date"
        }
```
#### [GET] /categories
Request:
- description : Get Category
- headers : Authorization (Bearer Token String)

Response:
- status 200
- data :
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
#### [PATCH] /categories/:categoryID
Request:
- description : Update Category By ID
- headers : Authorization (Bearer Token String)
- body:
  ```json
        {
          "type" : "string"
        }
  ```
Response :
- status 200
- data : 
```json
    {
        "id": "integer",
        "type": "string",
        "sold_product_amount": "integer",
        "updated_at": "date"
    }
```
#### [DELETE] /categories/:categoryID
Request:
- description : Delete Category By ID
- headers : Authorization (Bearer Token String)
- params : categoryID(id)
Response:
- status 200
- data:
```json
    {
        "message": "Category has been successfully deleted"
    }
```

### Products
#### [POST] /products
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
Response : 
- status 201
- data:
```json
    {
        "id": "integer",
        "title": "string",
        "price": "integer",
        "stock": "integer",
        "category_id": "integer",
        "created_at": "date"
    }
```
#### [GET] /products
Request:
- description : Get All Products
- headers : Authorization (Bearer Token String)
Response:
- status 200
- data:
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
#### [PUT] /products/:productID
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
Response :
- status 200
- data : 
```json
    {
        "Product": {
            "id": "integer",
            "title": "string",
            "price": "rupiah money format",
            "stock": "integer",
            "Category_Id": "integer",
            "created_at": "date",
            "updated_at": "date"
        }
    }
```
##### [DELETE] /products/:productID
Request:
- description : Delete Product
- headers : Authorization (Bearer Token String)
- params : productID(integer)
Response: 
- status 200
- data :
```json
    {
        "message": "Product has been successfully deleted"
    }
```

### Transaction History
#### [POST] /transactions
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
Response:
- status 201
- data :
```json
    {
        "message": "You have successfully purchased the product",
        "transaction_bill": {
            "total_price": "integer",
            "quantity": "integer",
            "product_title": "string"
        }
    }
```

#### [GET] /transactions/my-transactions
Request:
- description : Get My Transaction
- headers : Authorization (Bearer Token String)
Response:
- status 200
- data :
```json
    [
        {
            "id": "integer",
            "product_id": "integer",
            "user_id": "integer",
            "quantity": "integer",
            "total_price": "integer",
            "Product": {
                "id": "integer",
                "title": "string",
                "price": "integer",
                "stock": "integer",
                "category_id": "integer",
                "created_at": "date",
                "updated_at": "date"
            }
        }
    ]
```

#### [GET] /transactions/user-transactions
Request:
- description : Get All Transactions
- headers : Authorization (Bearer Token String)
Response:
- status 200
- data :
```json
    [
    {
        "id": "integer",
        "product_id": "integer",
        "user_id": "integer",
        "quantity": "integer",
        "total_price": "integer",
        "Product": {
            "id": "integer",
            "title": "string",
            "price": "integer",
            "stock": "integer",
            "category_id": "integer",
            "created_at": "date",
            "updated_at": "date"
        },
        "User": {
            "id": "integer",
            "email": "string",
            "full_name": "string",
            "balance": "integer",
            "created_at": "date",
            "updated_at": "date"
        }
    }
    ]
```
