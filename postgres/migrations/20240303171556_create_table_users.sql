-- +goose Up
-- +goose StatementBegin
create table users (
  id serial primary key,
  email varchar,
  name varchar,
  role varchar,
  password varchar,
  password_confirm varchar,
  created_at timestamp not null default now(),
  updated_at timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
