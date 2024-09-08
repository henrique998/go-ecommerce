-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS whishlist_items(
  id UUID PRIMARY KEY,
  whishlist_id UUID NOT NULL,
  product_id UUID NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  FOREIGN KEY (whishlist_id) REFERENCES whishlists(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS whishlist_items;
-- +goose StatementEnd