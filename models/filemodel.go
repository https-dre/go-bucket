package models

type FileMetaData struct {
	Id string `json:"id"`
	Filename string `json:"filename"`
	Path string `json:"path"`
	Filetype string `json:"filetype"`
	CreatedAtUTC string `json:"createdAtUTC"`
	Key string `json:"key"`
}
/* 
CREATE TABLE IF NOT EXISTS files (
    id TEXT PRIMARY KEY,
    filename TEXT NOT NULL,
    path TEXT NOT NULL,
    filetype TEXT NOT NULL,
    createdAtUTC TEXT NOT NULL
);
*/