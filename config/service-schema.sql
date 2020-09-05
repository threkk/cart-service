CREATE TABLE IF NOT EXISTS "carts" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  "user_id" integer NOT NULL,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL
);
CREATE INDEX "index_cart_on_user_id" ON "carts" ("user_id");
CREATE TABLE IF NOT EXISTS "line_items" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  "product_id" integer,
  "cart_id" integer NOT NULL,
  "order_id" integer,
  "quantity" integer DEFAULT 1,
  "price" decimal,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL,
  CONSTRAINT "fk_af645e8e5f" FOREIGN KEY ("cart_id") REFERENCES "carts" ("id")
);
