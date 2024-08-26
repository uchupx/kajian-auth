CREATE TABLE refresh_tokens (
    user_id varchar(36) NOT NULL unique,
    client_app_id varchar(36) NOT NULL,
    tokens text NOT NULL,
    expired_at DATETIME NOT NULL
  );
