CREATE TABLE IF NOT EXISTS users (
	user_id varchar(100) PRIMARY KEY,
    name varchar(100) NOT NULL DEFAULT '',
    email varchar(60) NOT NULL DEFAULT '',
    password varchar(30) NOT NULL
);