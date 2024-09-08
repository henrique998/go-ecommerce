-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS addresses(
  id UUID PRIMARY KEY,
  account_id UUID NOT NULL,
  address_line_1 TEXT NOT NULL,
  address_line_2 TEXT NOT NULL,
  city VARCHAR(100) NOT NULL,
  state VARCHAR(20) NOT NULL,
  postal_code VARCHAR(50) NOT NULL,
  country VARCHAR(100) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,
  FOREIGN KEY (account_id) REFERENCES accounts(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS addresses;
-- +goose StatementEnd
