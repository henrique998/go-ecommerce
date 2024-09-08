-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS shopping_carts(
  id UUID PRIMARY KEY,
  account_id UUID UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  FOREIGN KEY (account_id) REFERENCES accounts(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shopping_carts;
-- +goose StatementEnd
