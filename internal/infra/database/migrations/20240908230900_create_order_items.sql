
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_items(
  id UUID PRIMARY KEY,
  order_id UUID NOT NULL,
  product_id UUID NOT NULL,
  quantity INTEGER NOT NULL,
  price_at_purchase INTEGER NOT NULL,
  FOREIGN KEY (order_id) REFERENCES orders(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;
-- +goose StatementEnd
