CREATE TABLE IF NOT EXISTS dinos (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "avatar" VARCHAR,
  "configs" JSON,
  "user_id" BIGSERIAL REFERENCES users(id),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);
