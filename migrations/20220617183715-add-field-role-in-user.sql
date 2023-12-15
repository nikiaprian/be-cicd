
-- +migrate Up
ALTER TABLE Users
ADD role varchar(255) DEFAULT 'user' NOT NULL;
-- +migrate Down
ALTER TABLE Users
DROP Users role;