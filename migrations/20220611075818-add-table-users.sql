
-- +migrate Up
CREATE TABLE IF NOT EXISTS Users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  email VARCHAR(255) UNIQUE NOT NULL ,
  username VARCHAR(255) UNIQUE NOT NULL ,
  password VARCHAR(255) DEFAULT '',
  status VARCHAR(255) CHECK( status IN ('ACTIVE', 'DELETED') ) DEFAULT 'ACTIVE',
  provider VARCHAR(255) CHECK( provider IN ('local', 'google') )  DEFAULT 'local',
  created_at DATETIME,
  updated_at DATETIME
);

CREATE INDEX IF NOT EXISTS idx_email ON Users(email);
CREATE INDEX IF NOT EXISTS idx_username ON Users(username);
-- +migrate Down
DROP TABLE IF EXISTS Users;