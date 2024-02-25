-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_roadmap_categories (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  name TEXT UNIQUE NOT NULL,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE t_roadmap_categories;
-- +goose StatementEnd
