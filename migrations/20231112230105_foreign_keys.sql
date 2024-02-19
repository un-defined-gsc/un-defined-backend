-- +goose Up

ALTER TABLE t_action_log  ADD CONSTRAINT fk_action_logs_users FOREIGN KEY (user_id) REFERENCES t_users (id) ON DELETE NO ACTION;
ALTER TABLE t_login_logs ADD CONSTRAINT fk_login_logs_users FOREIGN KEY (user_id) REFERENCES t_users (id) ON DELETE NO ACTION;
ALTER TABLE t_banned ADD CONSTRAINT fk_banned_users FOREIGN KEY (user_id) REFERENCES t_users (id) ON DELETE NO ACTION;
ALTER TABLE t_mfa_settings ADD CONSTRAINT fk_mfa_settings_users FOREIGN KEY (user_id) REFERENCES t_users (id) ON DELETE CASCADE;

-- +goose Down

ALTER TABLE t_action_log DROP CONSTRAINT fk_action_logs_users;
ALTER TABLE t_login_logs DROP CONSTRAINT fk_login_logs_users;
ALTER TABLE t_banned DROP CONSTRAINT fk_banned_users;
ALTER TABLE t_mfa_settings DROP CONSTRAINT fk_mfa_settings_users;
