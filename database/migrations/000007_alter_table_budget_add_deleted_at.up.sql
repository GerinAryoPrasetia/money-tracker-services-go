ALTER TABLE budget
ADD COLUMN deleted_by varchar(50),
ADD COLUMN deleted_at timestamp;