-- +goose Up
-- +goose StatementBegin
CREATE TABLE account_roles (
  account_id UUID REFERENCES accounts(id),
  role_id UUID REFERENCES roles(id),
  PRIMARY KEY(account_id, role_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account_roles;
-- +goose StatementEnd
