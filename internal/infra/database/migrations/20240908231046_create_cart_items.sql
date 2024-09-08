
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cart_items(
  id UUID PRIMARY KEY,
  cart_id UUID NOT NULL,
  product_id UUID NOT NULL,
  quantity INTEGER NOT NULL,
  FOREIGN KEY (cart_id) REFERENCES shopping_carts(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cart_items;
-- +goose StatementEnd
