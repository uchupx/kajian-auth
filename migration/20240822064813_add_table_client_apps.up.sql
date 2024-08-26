CREATE TABLE client_apps(
    id varchar(36) NOT NULL UNIQUE,
    `key` text NOT NULL UNIQUE,
    name varchar(255) NOT NULL,
    secret text NOT NULL UNIQUE,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)
