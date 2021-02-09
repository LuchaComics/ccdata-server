CREATE TABLE countries (
    id BIGINT PRIMARY KEY,
    code VARCHAR (10) UNIQUE NOT NULL,
    name VARCHAR (255) NOT NULL
);
CREATE UNIQUE INDEX idx_country_id
ON countries (id);
CREATE UNIQUE INDEX idx_country_name
ON countries (name);

--TODO: PUBLISHERS
--TODO: SERIES
--TODO: ISSUES
--TODO: OTHER IF NEED BE
