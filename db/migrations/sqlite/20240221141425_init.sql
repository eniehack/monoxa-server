-- migrate:up
PRAGMA foreign_keys=true;

CREATE TABLE users (
    ulid VARCHAR PRIMARY KEY,
    uid VARCHAR NOT NULL UNIQUE,
    alias_id VARCHAR(15) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE UNIQUE INDEX idx__users__alias_id ON users (alias_id);
CREATE UNIQUE INDEX idx__users__ulid ON users (ulid);

-- migrate:down

DROP INDEX idx__users__alias_id ON users;
DROP INDEX idx__users__ulid ON users;
DROP TABLE users;
