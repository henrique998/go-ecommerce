-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS categories(
  id UUID PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd
