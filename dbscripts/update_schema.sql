CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    username TEXT NOT NULL UNIQUE,
    user_password TEXT NOT NULL,
    user_role TEXT NOT NULL,
    access_token TEXT,
    CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE INDEX user_access_token ON users (access_token);

INSERT INTO users (username, user_password, user_role)
VALUES
    ('admin', crypt('admin', gen_salt('bf')), 'admin'),
    ('runner', crypt('runner', gen_salt('bf')), 'runner');
