CREATE TABLE IF NOT EXISTS users (
	id varchar(100) PRIMARY KEY,
    name varchar(100) NOT NULL DEFAULT '',
    email varchar(60) NOT NULL DEFAULT '',
    password varchar(30) NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    created_by varchar(60) NOT NULL DEFAULT '',
    updated_at timestamp DEFAULT NOW(),
    updated_by varchar(60)
);