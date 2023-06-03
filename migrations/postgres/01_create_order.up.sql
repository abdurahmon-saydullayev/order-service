CREATE TABLE IF NOT EXISTS "order" (
    "id" UUID PRIMARY KEY,
    "product_id" VARCHAR(30) NOT NULL,
    "user_id" VARCHAR(100) NOT NULL,
    "status" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
); 