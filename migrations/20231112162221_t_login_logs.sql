-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS t_login_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID,
    login_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    login_ip VARCHAR(15),
    login_user_agent TEXT,
    is_success BOOLEAN DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE t_login_logs;
-- +goose StatementEnd
