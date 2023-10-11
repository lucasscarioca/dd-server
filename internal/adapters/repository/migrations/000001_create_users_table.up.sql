CREATE TABLE users (
  "id" VARCHAR NOT NULL,
  "name" VARCHAR NOT NULL,
  "avatar" VARCHAR,
  "email" VARCHAR NOT NULL,
  "password" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),

  CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "email" ON "users" ("email");