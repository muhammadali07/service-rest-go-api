CREATE DATABASE rest_middleware_dev;

CREATE TABLE IF NOT EXISTS user_groups (
    user_groups_id SERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    role_id VARCHAR(30) NOT NULL,
    create_at TIMESTAMP NOT NULL
);
