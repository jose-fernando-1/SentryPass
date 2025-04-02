-- Create initial schema.

DROP TABLE IF EXISTS Entries;
DROP TABLE IF EXISTS Logs;
DROP TABLE IF EXISTS Categories;
DROP TABLE IF EXISTS User;

CREATE TABLE IF NOT EXISTS User (
    id_user INTEGER PRIMARY KEY CHECK(id_user = 1),
    master_password_hash TEXT NOT NULL,
    master_password_salt TEXT NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Logs(
    id_log INTEGER PRIMARY KEY,
    action TEXT NOT NULL,
    timestamp TEXT DEFAULT CURRENT_TIMESTAMP,
    id_user INTEGER NOT NULL,
    FOREIGN KEY (id_user) REFERENCES User(id_user) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Categories(
    id_category INTEGER PRIMARY KEY,
    type TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Entries(
    id_entry INTEGER PRIMARY KEY,
    title TEXT UNIQUE NOT NULL,
    username TEXT NOT NULL,
    password_encrypted TEXT NOT NULL,
    url TEXT,
    details_encrypted TEXT,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TEXT DEFAULT NULL,
    id_category INTEGER NOT NULL,
    id_user INTEGER NOT NULL,
    FOREIGN KEY (id_user) REFERENCES User(id_user) ON DELETE CASCADE,
    FOREIGN KEY (id_category) REFERENCES Categories(id_category) ON DELETE CASCADE
);