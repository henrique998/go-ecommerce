-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders(
  id UUID PRIMARY KEY,
  account_id UUID NOT NULL,
  total_price INTEGER NOT NULL,
  status payment_status NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  FOREIGN KEY (account_id) REFERENCES accounts(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
