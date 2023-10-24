CREATE TABLE users (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "avatar" VARCHAR,
  "email" VARCHAR NOT NULL,
  "password" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
);

CREATE UNIQUE INDEX "email" ON "users" ("email");