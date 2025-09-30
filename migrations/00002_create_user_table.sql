-- +goose Up
-- +goose StatementBegin

-- Create roles enum
CREATE TYPE user_role AS ENUM ('admin', 'manager', 'user');

-- Create users table with role
CREATE TABLE users(
                      id SERIAL PRIMARY KEY,
                      email VARCHAR(255) UNIQUE NOT NULL,
                      password VARCHAR(255) NOT NULL,
                      role user_role NOT NULL DEFAULT 'user',
                      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create index on email for faster lookups
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);

-- Create SCS sessions table (replacing old sessions table)
CREATE TABLE sessions (
                          token TEXT PRIMARY KEY,
                          data BYTEA NOT NULL,
                          expiry TIMESTAMPTZ NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS sessions_expiry_idx;
DROP TABLE IF EXISTS sessions;

DROP INDEX IF EXISTS idx_users_role;
DROP INDEX IF EXISTS idx_users_email;
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS user_role;

-- +goose StatementEnd