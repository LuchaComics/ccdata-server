CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    uuid VARCHAR (36) UNIQUE NOT NULL,
    first_name VARCHAR (50) NULL,
    last_name VARCHAR (50) NULL,
    email VARCHAR (255) UNIQUE NOT NULL,
    password_hash VARCHAR (511) NOT NULL,
    state SMALLINT NOT NULL,
    timezone VARCHAR (63) NOT NULL DEFAULT 'utc',
    created_time TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    session_uuid VARCHAR (36) UNIQUE NOT NULL
);
CREATE UNIQUE INDEX idx_user_uuid
ON users (uuid);
CREATE UNIQUE INDEX idx_user_session_uuid
ON users (session_uuid);

CREATE TABLE countries (
    id BIGINT PRIMARY KEY,
    code VARCHAR (10) UNIQUE NOT NULL,
    name VARCHAR (255) NOT NULL
);
CREATE UNIQUE INDEX idx_country_id
ON countries (id);
CREATE UNIQUE INDEX idx_country_name
ON countries (name);

CREATE TABLE publishers (
    id BIGINT PRIMARY KEY,
    name VARCHAR (255) NOT NULL
);

--TODO: PUBLISHERS
--TODO: SERIES
--TODO: ISSUES
--TODO: OTHER IF NEED BE
