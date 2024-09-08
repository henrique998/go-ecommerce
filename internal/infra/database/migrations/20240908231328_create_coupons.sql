
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS coupons(
  id UUID PRIMARY KEY,
  code VARCHAR(100) UNIQUE NOT NULL,
  discount_percentage INTEGER NOT NULL,
  expires_at TIMESTAMP NOT NULL,
  usage_limit INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS coupons;
-- +goose StatementEnd