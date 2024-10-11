CREATE TABLE IF NOT EXISTS files (
    id TEXT PRIMARY KEY,
    filename TEXT NOT NULL,
    path TEXT NOT NULL,
    filetype TEXT NOT NULL,
    createdAtUTC TEXT NOT NULL,
    key TEXT UNIQUE NOT NULL
);