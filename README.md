# Go-Cart
Mini Service Shopping Cart with Gin-Gonic HTTP Respons Framework, GORM for Object relation model, PostgreSQL for database.

## 🔗 Description

This Backend Service Shopping Cart is used for simple order product.
There are 3 main routers :
1. Add Cart (Can add product if product is no exist also update quantity product if name product exist)
> Attributes ProductCode, ProductName, Quantity
2. Delete Product (Delete by prodcut code)
3. Get Product (Get all product, Get by product name, Get by quantity)

## Notes :
1. I'am using UUID for cart_id, don't forget to create extenxions in SQL console after create the database with this query below :
```bash
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```
2. For answers no. 01 & 02, i put it in a folder soal_01 & soal_02, change directory to folder and run
```bash
go run .
```
3. For answer no. 03 - 05 i put in a folder soal_03-05, open the .pdf file / .svg

## Several command you must know in this app :
```bash
1. go run . serve //to run the app / server
2. go run . migrate -u //for database migration
# or
go run . migrate -d //for rollback
```

## 🛠️ Installation Steps

1. Clone the repository

```bash
https://github.com/adiet95/go-cart.git
```

2. Install dependencies

```bash
go mod tidy
```
> Wait a minute, if still error run :

```bash
go mod vendor
```

3. Add Env File

```sh
  USER = Your DB User
  HOST = Your DB Host
  DB   = Your DB Name
  PASS = Your DB Password
  PORT = Your Port
```

4. Database Migration and Rollback

```bash
go run main.go migrate --up //for database migration table
# or
go run main.go migrate --down //for rollback the database
```

5. Run the app

```bash
go run . serve
```

### 🚀 You are all set

## 🔗 REST API Endpoints

### GET /cart
> Get Data Cart

_Request Body_
```
not needed
```
_Request Query Params_
```
not needed
```

### POST /cart
> Post Data Cart

_Request Body_
```
{
    "product_code" : "a-03",
    "product_name" : "kabel-1",
    "quantity" : 2
}
```
_Request Query Params_
```
no need
```

### DELETE /cart
> Delete Data Cart by Product Code

_Request Body_
```
no need
```
_Request Query Params_
```
code = (input product code)
```

### GET /cart/name

> Search Data Cart by product name

_Request Body_
```
no need
```
_Request Query Params_
```
name = (input product name)
```

### GET /cart/qty

> Search Data Cart by quantity

_Request Body_
```
no need
```
_Request Query Params_
```
qty = (input quantity)
```

## 💻 Built with

- [Golang](https://go.dev/): Go Programming Language
- [Gin-Gonic](https://gin-gonic.com/): for handle http request
- [Postgres](https://www.postgresql.org/): for DBMS


## 🚀 About Me

- Linkedin : [Achmad Shiddiq](https://www.linkedin.com/in/achmad-shiddiq-alimudin/)
