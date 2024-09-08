-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS reviews(
  id UUID PRIMARY KEY,
  account_id UUID NOT NULL,
  product_id UUID NOT NULL,
  rating INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,
  FOREIGN KEY (account_id) REFERENCES accounts(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reviews;
-- +goose StatementEnd
