-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS t_banned (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID,
    admin_id UUID,
    reason TEXT,
    permanent BOOLEAN DEFAULT FALSE,
    expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE t_banned;
-- +goose StatementEnd
