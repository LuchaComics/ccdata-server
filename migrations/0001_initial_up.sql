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

CREATE TABLE series (
    id BIGINT PRIMARY KEY,
    name VARCHAR (255) NOT NULL,
    sort_name VARCHAR(255) NOT NULL,
    format VARCHAR(255) NOT NULL DEFAULT '',
    year_began INTEGER NOT NULL,
    year_began_uncertain BOOLEAN NOT NULL DEFAULT FALSE,
    year_ended INTEGER DEFAULT NULL,
    year_ended_uncertain BOOLEAN NOT NULL DEFAULT FALSE,
    publication_dates VARCHAR(255) NOT NULL DEFAULT '',
    first_issue_id BIGINT DEFAULT NULL,
    last_issue_id BIGINT DEFAULT NULL,
    is_current BOOLEAN NOT NULL DEFAULT FALSE,
    publisher_id BIGINT NOT NULL,
    country_id BIGINT NOT NULL,
    language_id BIGINT NOT NULL,
    tracking_notes TEXT NOT NULL,
    notes TEXT NOT NULL,
    has_gallery BOOLEAN NOT NULL DEFAULT FALSE,
    issue_count INTEGER NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    has_indicia_frequency BOOLEAN NOT NULL DEFAULT TRUE,
    has_isbn BOOLEAN NOT NULL DEFAULT TRUE,
    has_barcode BOOLEAN NOT NULL DEFAULT TRUE,
    has_issue_title BOOLEAN NOT NULL DEFAULT FALSE,
    has_volume BOOLEAN NOT NULL DEFAULT TRUE,
    is_comics_publication BOOLEAN NOT NULL DEFAULT TRUE,
    color varchar(255) NOT NULL,
    dimensions varchar(255) NOT NULL,
    paper_stock varchar(255) NOT NULL,
    binding varchar(255) NOT NULL,
    publishing_format varchar(255) NOT NULL,
    has_rating BOOLEAN NOT NULL,
    publication_type_id INTEGER DEFAULT NULL,
    is_singleton BOOLEAN NOT NULL,
    has_about_comics BOOLEAN NOT NULL,
    has_indicia_printer BOOLEAN NOT NULL,
    FOREIGN KEY (publisher_id) REFERENCES publishers(id),
    FOREIGN KEY (country_id) REFERENCES countries(id)
);

CREATE TABLE issues (
    id BIGINT PRIMARY KEY,
    number VARCHAR(255) NOT NULL,
    volume VARCHAR(50) NOT NULL DEFAULT '',
    no_volume BOOLEAN NOT NULL DEFAULT FALSE,
    display_volume_with_number BOOLEAN NOT NULL DEFAULT FALSE,
    series_id BIGINT NOT NULL,
    indicia_publisher_id BIGINT DEFAULT NULL,
    indicia_pub_not_printed BOOLEAN NOT NULL,
    brand_id BIGINT DEFAULT NULL,
    no_brand BOOLEAN NOT NULL,
    publication_date VARCHAR(255) NOT NULL,
    key_date VARCHAR(10) NOT NULL,
    sort_code VARCHAR(11) NOT NULL,
    price VARCHAR(255) NOT NULL,
    page_count VARCHAR(11) DEFAULT NULL,
    page_count_uncertain BOOLEAN NOT NULL DEFAULT FALSE,
    indicia_frequency VARCHAR(255) NOT NULL DEFAULT '',
    no_indicia_frequency BOOLEAN NOT NULL DEFAULT FALSE,
    editing TEXT NOT NULL,
    no_editing BOOLEAN NOT NULL DEFAULT FALSE,
    notes TEXT NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    is_indexed BOOLEAN NOT NULL DEFAULT FALSE,
    isbn VARCHAR(32) NOT NULL DEFAULT '',
    valid_isbn VARCHAR(13) NOT NULL DEFAULT '',
    no_isbn BOOLEAN NOT NULL DEFAULT FALSE,
    variant_of_id BIGINT DEFAULT NULL,
    variant_name VARCHAR(255) NOT NULL DEFAULT '',
    FOREIGN KEY (series_id) REFERENCES series(id)
);

  `barcode` varchar(38) NOT NULL DEFAULT '',
  `no_barcode` tinyint(1) NOT NULL DEFAULT '0',
  `title` varchar(255) NOT NULL DEFAULT '',
  `no_title` tinyint(1) NOT NULL DEFAULT '0',
  `on_sale_date` varchar(10) NOT NULL,
  `on_sale_date_uncertain` tinyint(1) NOT NULL DEFAULT '0',
  `rating` varchar(255) NOT NULL,
  `no_rating` tinyint(1) NOT NULL,
  `volume_not_printed` tinyint(1) NOT NULL,
  `no_indicia_printer` tinyint(1) NOT NULL,
