# Shopping cart microservice

Your mission, should you choose to accept it, is to create a shopping cart microservice that extracts only the relevant functionality out of a monolithic application.

We ask you to push your solution to Github and send the link of the repository in reply to the email.

## Current situation

You are given the following database schema of a monolithical order management system:

```sql
CREATE TABLE IF NOT EXISTS "products" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
   "title" varchar,
   "description" text,
   "image_url" varchar,
   "price" decimal(8,2),
   "created_at" datetime NOT NULL,
   "updated_at" datetime NOT NULL
);

CREATE TABLE IF NOT EXISTS "carts" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  "user_id" integer,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL,
  CONSTRAINT "fk_32e17d5e33" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE INDEX "index_cart_on_user_id" ON "carts" ("user_id");

CREATE TABLE IF NOT EXISTS "line_items" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  "product_id" integer,
  "cart_id" integer,
  "order_id" integer,
  "quantity" integer DEFAULT 1,
  "price" decimal,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL,
  CONSTRAINT "fk_11e15d5c6b" FOREIGN KEY ("product_id") REFERENCES "products" ("id"),
  CONSTRAINT "fk_af645e8e5f" FOREIGN KEY ("cart_id") REFERENCES "carts" ("id")
);

CREATE INDEX "index_line_items_on_product_id" ON "line_items" ("product_id");
CREATE INDEX "index_line_items_on_cart_id" ON "line_items" ("cart_id");

CREATE TABLE IF NOT EXISTS "orders" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  "user_id" integer,
  "name" varchar,
  "address" text,
  "email" varchar,
  "pay_type" integer,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL,
  CONSTRAINT "fk_22e24d5c6b" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE INDEX "index_orders_on_user_id" ON "orders" ("user_id");

CREATE TABLE IF NOT EXISTS "users" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  "name" varchar,
  "password_digest" varchar,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL
);
```

Note: if you want to load this schema use `sqlite3`:
```
sqlite3 --init schema.sql
```

## Instructions

* Extract the **shopping cart domain** and model it into a **backend** microservice in a structured way.
  * Migrate as much (or as little) of the monolithic database schema into your microservice as you see fit
  * Think what other microservices of the order management system could/should exist and how they interact with your shopping cart service.
* Expose an API to other microservices of an order management system with the following operations:
  * Create and persist a shopping cart
  * Add products to a shopping cart
  * Remove products from a shopping cart
  * Empty a shopping cart
  * Get the details of a cart

## Guidelines

* You can use a language or stack of your choice (we use **Ruby**/**Rails** or **Go**)
* You can choose to build a **RESTful API** using **JSON** or use **gRPC** using **Protocol Buffers** instead.
* You don't need to build a UI for the shopping cart microservice
* There is no need for overengineering, we just want to see readable and tested code.
* This take-home problem is timeboxed.
  * We don't want you to spend 3 full-time days on this, more like couple of evenings.
  * Don't worry if you can't include everything, pick your battles carefully.
  * Just make sure that the things you do include work.

## Bonus points (optional!)

* Including a way to consume your API
* Instrumenting your microservice
* a Dockerfile

## Out of scope

* Application deployment and scaling. We just want to run an instance of your service, not a Kubernetes cluster.
