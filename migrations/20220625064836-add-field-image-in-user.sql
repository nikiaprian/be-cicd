
-- +migrate Up
ALTER TABLE Users
ADD photo VARCHAR(255);
-- +migrate Down
ALTER TABLE Users
DROP Users photo;