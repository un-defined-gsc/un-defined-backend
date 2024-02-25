-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_submissions (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  roadmap_id uuid NOT NULL,
  user_id uuid NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE t_submissions;
-- +goose StatementEnd
