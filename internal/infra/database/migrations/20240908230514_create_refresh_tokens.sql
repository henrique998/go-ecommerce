-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS refresh_tokens(
  id  UUID PRIMARY KEY,
  value TEXT UNIQUE NOT NULL,
  account_id UUID UNIQUE NOT NULL,
  expires_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS refresh_tokens;
-- +goose StatementEnd