-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles (
  id UUID PRIMARY KEY,
  name VARCHAR(50) UNIQUE NOT NULL
);
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
