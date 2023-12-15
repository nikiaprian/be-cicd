
-- +migrate Up
CREATE TABLE IF NOT EXISTS GoogleAccounts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  google_id VARCHAR(255) NOT NULL,
  user_id INTEGER NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,

  FOREIGN KEY (google_id)
  REFERENCES Users (id)
);

CREATE INDEX IF NOT EXISTS idx_user_google_account ON GoogleAccounts(user_id);
-- +migrate Down
DROP TABLE IF EXISTS GoogleAccounts;