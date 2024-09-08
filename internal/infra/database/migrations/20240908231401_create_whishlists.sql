-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS whishlists(
  id UUID PRIMARY KEY,
  account_id UUID UNIQUE NOT NULL,
  FOREIGN KEY (account_id) REFERENCES accounts(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS whishlists;
-- +goose StatementEnd
