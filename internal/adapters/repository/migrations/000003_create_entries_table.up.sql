CREATE TABLE IF NOT EXISTS entries (
  "id" BIGSERIAL PRIMARY KEY,
  "title" VARCHAR NOT NULL,
  "content" TEXT,
  "user_id" BIGSERIAL REFERENCES users(id),
  "status" VARCHAR,
  "configs" JSON,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);