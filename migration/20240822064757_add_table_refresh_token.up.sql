CREATE TABLE refresh_tokens (
    id varchar(36) NOT NULL PRIMARY KEY,
    user_id varchar(36) NOT NULL unique,
    client_app_id varchar(36) NOT NULL,
    access_token text NOT NULL,
    refresh_token text NOT NULL,
    expired_at DATETIME NOT NULL
  );
