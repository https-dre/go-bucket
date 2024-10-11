package repositories

import (
	"database/sql"
	"go-bucket/models"
	"log"
)

type FileRepository struct {
	db *sql.DB
}

func CreateFileRepository(db *sql.DB) FileRepository {
	return FileRepository{db: db}
}

func (repo *FileRepository) PersistFile(file *models.FileMetaData) error {
	insertFile := "INSERT INTO files (id, filename, path, filetype, createdAtUTC, key) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := repo.db.Exec(insertFile, file.Id, file.Filename, file.Path, file.Filetype, file.CreatedAtUTC, file.Key)
	if (err != nil) {
		log.Println("Persist file error: ", err)
		return err
	}
	return nil
}