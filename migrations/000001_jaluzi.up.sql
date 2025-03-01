-- Admin table migration
CREATE TABLE IF NOT EXISTS "admin"(
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(25) NOT NULL,
    "password" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
);

-- Product table migration
CREATE TABLE IF NOT EXISTS "product"(
     "id" UUID PRIMARY KEY,
     "name" VARCHAR(250) NOT NULL,
     "code" VARCHAR(20) NOT NULL,
     "price" DECIMAL(10, 2) NOT NULL,
     "product_image" VARCHAR(255),
     "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     "updated_at" TIMESTAMP,
     "deleted_at" TIMESTAMP
);

-- Order table migration
CREATE TABLE IF NOT EXISTS "order"(
     "id" UUID PRIMARY KEY,
     "order_name" VARCHAR(250) NOT NULL,
     "address" VARCHAR(500) NOT NULL,
     "phone_number" VARCHAR(100) NOT NULL,
     "finished_data" DATE NOT NULL,
     "total_all_summ" DECIMAL(10, 2) NOT NULL,
     "date" DATE NOT NULL,
     "final_sum" DECIMAL(10, 2) NOT NULL,
     "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     "updated_at" TIMESTAMP,
     "deleted_at" TIMESTAMP
);

-- Order_items table migration
CREATE TABLE IF NOT EXISTS "order_items" (
     "id" UUID PRIMARY KEY,
     "order_id" UUID REFERENCES "order"("id"),
     "product_id" UUID REFERENCES "product"("id"),
     "height" DECIMAL(10, 2) NOT NULL,
     "width" DECIMAL(10, 2) NOT NULL,
     "price" DECIMAL(10, 2) NOT NULL,
     "total_summ" DECIMAL(10, 2) NOT NULL,
     "square_metrs" DECIMAL(10, 2) NOT NULL,
     "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     "updated_at" TIMESTAMP,
     "deleted_at" TIMESTAMP
);
