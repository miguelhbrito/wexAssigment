CREATE TABLE IF NOT EXISTS "transaction" (
  "id" varchar PRIMARY KEY,
  "description" varchar NOT NULL,
  "price" float NOT NULL,
  "country" varchar NOT NULL,
  "created_at" timestamp NOT NULL
);