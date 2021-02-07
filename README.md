# 🛒🌐 eShop

The aim of this project is to be a proof-of-concept using `Hexagonal Architecture` (Ports-and-adapters) applying `DDD` in `Go` 🔥

The code will be reviewed over time to time while I will introduce new concepts about the Hexagonal Architecture or DDD and new features  to it; so for sure will be refactored several times.

## 👩‍💻 Project

The idea of this project is to support an E-Commerce called `eShop` (super original 🙄). It has to provide a user management system and an order and product catalogue.

The language chosen for this has been `Go` and the database `PostgreSQL`. For the simplicity of the case I've added a `docker-compose` to create the db automatically and a `Makefile` with a few commands described above ✨

## 📐 Architecture

As mentioned before, the project it's using the Hexagonal Architecture and the layers are 📁

- application - it contains the `application services` (the use cases).
- domain - it defines the core of our applications; it contains the `entities`, the `factories` and the `repositories interfaces`.
- infrastructure - it contains the implementation of the `repositories` defined in the domain and the `api` definition (router and handlers).

- cmd - the multiple commands offered by the project.

## 📝 Commands

Use `docker-compose up` to start the database.
Use `make run` to start the server.
Use `make migrations` to run the migrations.

### Sign In User

Endpoint: /users
Request: POST

```
curl -X POST http://localhost:8080/users -d '{"email":"charly@eshop.com", "name":"charly3pins"}'
```

### Add products to Order

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

## 🤝 Contributing

As with all my projects you are more than welcome to contribute. You can open an issue, a pull request or just send me a message and we can chat about the topic you want 🚀
