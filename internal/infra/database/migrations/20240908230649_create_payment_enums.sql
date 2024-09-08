-- +goose Up
-- +goose StatementBegin
CREATE TYPE payment_method AS ENUM ('debit_card', 'credit_card', 'money');
CREATE TYPE payment_status AS ENUM ('paid', 'pending', 'canceled');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS payment_method;
DROP TABLE IF EXISTS payment_status;
-- +goose StatementEnd