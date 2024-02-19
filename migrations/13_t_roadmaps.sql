-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_roadmaps (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  path_way TEXT NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE t_roadmaps;
-- +goose StatementEnd
