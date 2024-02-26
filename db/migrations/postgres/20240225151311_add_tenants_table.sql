-- migrate:up
CREATE TABLE notebooks (
    ulid VARCHAR PRIMARY KEY,
    alias_id VARCHAR(15) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE UNIQUE INDEX idx__notebooks__alias_id ON notebooks (alias_id);

CREATE TABLE notebooks_users (
    notebook_id VARCHAR NOT NULL,
    user_id VARCHAR NOT NULL,
    is_admin BOOLEAN DEFAULT False,
    FOREIGN KEY (notebook_id) REFERENCES notebooks(ulid),
    FOREIGN KEY (user_id) REFERENCES users(ulid)
);

-- migrate:down
DROP TABLE notebooks_users;
DROP INDEX idx__notebooks__alias_id ON notebooks;
DROP TABLE notebooks;
