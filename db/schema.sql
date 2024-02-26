CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
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
    is_admin BOOLEAN DEFAULT 0,
    FOREIGN KEY (notebook_id) REFERENCES notebooks(ulid),
    FOREIGN KEY (user_id) REFERENCES users(ulid)
);
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
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20240221141425'),
  ('20240225151311'),
  ('20240225151435');
