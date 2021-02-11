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
    name VARCHAR (255) NOT NULL,
    country_id BIGINT NOT NULL,
    year_began INTEGER NULL,
    year_began_uncertain BOOLEAN NOT NULL,
    year_ended INTEGER NULL,
    year_ended_uncertain BOOLEAN NOT NULL,
    notes TEXT NOT NULL DEFAULT '',
    url VARCHAR (255) NOT NULL DEFAULT '',
    brand_count INTEGER NOT NULL DEFAULT 0,
    indicia_publisher_count INTEGER NOT NULL DEFAULT 0,
    series_count INTEGER NOT NULL DEFAULT 0,
    issue_count INTEGER NOT NULL DEFAULT 0,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    year_overall_began INTEGER DEFAULT NULL,
    year_overall_began_uncertain BOOLEAN NOT NULL DEFAULT FALSE,
    year_overall_ended INTEGER DEFAULT NULL,
    year_overall_ended_uncertain BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (country_id) REFERENCES countries(id)
);

--TODO: SERIES
--TODO: ISSUES
--TODO: OTHER IF NEED BE
