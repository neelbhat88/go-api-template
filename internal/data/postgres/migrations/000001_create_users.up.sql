create table users (
  id serial primary key,
  email text not null unique,
  password text not null
);