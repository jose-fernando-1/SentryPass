-- Create initial schema.

DROP TABLE IF EXISTS Entry;

CREATE TABLE IF NOT EXISTS Entry (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- Unique identifier for each entry
    username TEXT NOT NULL,               -- Username or email
    password TEXT NOT NULL,               -- Encrypted password
    notes TEXT,                           -- Optional notes
    date_of_creation DATETIME DEFAULT CURRENT_TIMESTAMP, -- Timestamp of creation
    date_of_last_password_update DATETIME DEFAULT CURRENT_TIMESTAMP -- Timestamp of last update
);

-- Timestamp is stored in UTC timezone for best practices.
-- On display to the user, it should be converted to the user's timezone.

CREATE TRIGGER update_date_of_last_password_update
AFTER UPDATE OF password ON Entry
FOR EACH ROW
WHEN NEW.password != OLD.password
BEGIN
    UPDATE Entry
    SET date_of_last_password_update = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;
