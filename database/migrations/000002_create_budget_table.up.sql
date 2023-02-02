CREATE TABLE IF NOT EXISTS budget (
	id varchar(100) PRIMARY KEY,
    user_id varchar(100) NOT NULL DEFAULT '',
    amount int NOT NULL,
    description varchar(100) NOT NULL DEFAULT '',
    created_at timestamp DEFAULT NOW()
);