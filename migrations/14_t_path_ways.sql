-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_path_ways (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  roadmap_id  uuid NOT NULL ,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  parent_id uuid ,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE t_path_ways;
-- +goose StatementEnd
