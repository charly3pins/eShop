# eShop

Example of application using DDD in Go

## Commands

Use `docker-compose up` to start the database.
Use `make run` to start the server.
Use `make migrations` to run the migrations.

## Sign In User

Endpoint: /users
Request: POST

```
curl -X POST http://localhost:8080/users -d '{"email":"charly@eshop.com", "name":"charly3pins"}'
```

## Add products to Order

Endpoint: /orders
Request: POST

Creating an Order (empty `order_id`)
```
curl -X POST http://localhost:8080/orders -d '{"user_id":"48c20124-824a-4cf6-a377-d2efdc789bf3", "products":[{"name":"DDD book", "unit_price": 35, "currency": "EUR", "quantity": 1}, {"name":"post-it", "unit_price": 2, "currency": "EUR", "quantity": 3}, {"name":"pen", "unit_price": 1, "currency": "EUR", "quantity": 5}]}'
```

Adding a new product to an existing Order (with `order_id`)
```
curl -X POST http://localhost:8080/orders -d '{"user_id":"48c20124-824a-4cf6-a377-d2efdc789bf3", "order_id": "47fafa06-a73a-4733-966b-60914d1dd8ff", "products":[{"name":"IDDD book", "unit_price": 55, "currency": "EUR", "quantity": 1}]}'
```

Updating an existing product to an existing Order (with `order_id` and `product_id`)
```
curl -X POST http://localhost:8080/orders -d '{"user_id":"48c20124-824a-4cf6-a377-d2efdc789bf3", "order_id": "47fafa06-a73a-4733-966b-60914d1dd8ff", "products":[{"product_id":"c92ca476-adcd-4843-a5a3-8c442c2f7d4e", "name":"post-it", "unit_price": 2, "currency": "EUR", "quantity": 10}]}'
```