-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_images (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  post_id uuid ,
  user_id uuid ,
  category uuid NOT NULL,
  url VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
