-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_categories (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
-- +goose StatementEnd
INSERT INTO t_categories (name) VALUES ('story');
INSERT INTO t_categories (name) VALUES ('problem');
INSERT INTO t_categories (name) VALUES ('question');
INSERT INTO t_categories (name) VALUES ('jobadvert');

-- +goose Down
-- +goose StatementBegin
DROP TABLE t_categories;
-- +goose StatementEnd


