ALTER TABLE budget
ADD COLUMN created_by varchar(50) NOT NULL DEFAULT '',
ADD COLUMN updated_at timestamp NOT NULL DEFAULT NOW(),
ADD COLUMN updated_by varchar(50) NOT NULL DEFAULT '';