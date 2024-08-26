CREATE TABLE users(
  id VARCHAR(36) PRIMARY KEY NOT NULL,
  client_app_id VARCHAR(36) NOT NULL,
  name VARCHAR(255) NULL,
  password TEXT NOT NULL,
  email varchar(255) NOT NULL,
  username varchar(255) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NULL
);

