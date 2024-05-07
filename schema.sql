create table if not exists users (
  id text not null primary key,
  username text not null,
  phone text not null unique,
  password text not null,
  account_type text not null,
  description text not null,
  country text not null default '',
  town text not null default ''
);

create table if not exists categories (
  id text not null primary key,
  label text not null unique
);
