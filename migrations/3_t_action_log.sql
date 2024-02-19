-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS t_action_log (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID,
    action_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    action_slug TEXT NOT NULL,
    ip_address TEXT NOT NULL,
    user_agent TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE t_action_log;
-- +goose StatementEnd
