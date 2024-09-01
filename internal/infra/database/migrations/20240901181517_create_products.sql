-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products(
  id UUID PRIMARY KEY,
  name VARCHAR(70) NOT NULL,
  description TEXT NOT NULL,
  price INTEGER NOT NULL,
  stock_quantity INTEGER NOT NULL,
  image_url TEXT NOT NULL,  
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
