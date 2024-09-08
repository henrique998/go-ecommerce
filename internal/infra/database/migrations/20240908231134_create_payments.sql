
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payments(
  id UUID PRIMARY KEY,
  order_id UUID NOT NULL,
  payment_method payment_method NOT NULL,
  amount INTEGER NOT NULL,
  payment_status payment_status NOT NULL,
  transaction_id TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,
  FOREIGN KEY (order_id) REFERENCES orders(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payments;
-- +goose StatementEnd