-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_users (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    password TEXT NOT NULL, 
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    lang TEXT NOT NULL DEFAULT 'en',
    gender TEXT NOT NULL DEFAULT 'other',
    appeal TEXT,
    email TEXT NOT NULL UNIQUE,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    master_admin BOOLEAN NOT NULL DEFAULT FALSE,
    banned BOOLEAN NOT NULL DEFAULT FALSE,
    mfa_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    disabled BOOLEAN NOT NULL DEFAULT FALSE,
    disabled_at TIMESTAMP,
    PRIMARY KEY (id)
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE t_users;
-- +goose StatementEnd
