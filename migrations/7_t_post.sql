-- +goose Up
-- +goose StatementBegin
 CREATE TABLE IF NOT EXISTS t_post (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  category_id uuid NOT NULL,
  user_id uuid NOT NULL,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
  FOREIGN KEY (category_id) REFERENCES t_categories (id) ON DELETE CASCADE,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE t_post;
-- +goose StatementEnd
