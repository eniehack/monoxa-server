-- migrate:up

CREATE TABLE shouts (
    ulid VARCHAR NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    notebook_id VARCHAR NOT NULL,
    user_id VARCHAR NOT NULL,
    script TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    FOREIGN KEY (notebook_id) REFERENCES notebooks (ulid),
    FOREIGN KEY (user_id) REFERENCES users (ulid)
);

-- migrate:down

DROP TABLE shouts;
