CREATE TABLE IF NOT EXISTS assists (
  "assistant_id" BIGSERIAL REFERENCES users(id),
  "user_id" BIGSERIAL REFERENCES users(id),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  PRIMARY KEY ("assistant_id", "user_id")
);

CREATE INDEX "assists_user_id" ON "assists" ("user_id");
CREATE INDEX "assists_assistant_id" ON "assists" ("assistant_id");