-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_mfa_settings (
   user_id uuid NOT NULL UNIQUE,
   key TEXT NOT NULL UNIQUE,
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
  
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE t_mfa_settings;
-- +goose StatementEnd
