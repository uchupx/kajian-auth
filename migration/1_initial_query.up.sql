CREATE TABLE users(
  id VARCHAR(36) PRIMARY KEY NOT NULL,
  name VARCHAR(255) NULL,
  password varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  username varchar(255) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NULL
);