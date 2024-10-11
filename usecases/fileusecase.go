package usecases

import (
	"fmt"
	"go-bucket/models"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

type IFileRepository interface {
	PersistFile(file *models.File) error
}

type FileUseCase struct {
	repository IFileRepository
}

func CreateFileUseCase(repo IFileRepository) FileUseCase {
	return FileUseCase{repository: repo}
}

func (usecase *FileUseCase) UploadFileUseCase(file multipart.File, filename string) (*models.File, error) {
	fileId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	name := filename
	ext := ""
	if dot := lastDotIndex(filename); dot != -1 {
		name = filename[:dot]
		ext = filename[dot:] // mantém o ponto e a extensão
	}

	filenameFormatted := removeSpecialChars(name) + ext

	filetype, err := GetFileType(file)
	if err != nil {
		return nil, err
	}

	fileEntity := models.File{
		Id:          fileId.String(),
		Filename:    filenameFormatted,
		CreatedAtUTC: time.Now().UTC().String(),
		Path:        fmt.Sprintf("bucket/%s_%s", fileId.String(), filenameFormatted),
		Filetype:    filetype,
	}
	fmt.Println(fileEntity)

	err = saveFile(file, fileEntity.Path)
	if err != nil {
		return nil, err
	}
	err = usecase.repository.PersistFile(&fileEntity)
	if err != nil {
		return nil, err
	}

	return &fileEntity, nil
}

func lastDotIndex(s string) int {
	return strings.LastIndex(s, ".")
}

func saveFile(file multipart.File, filePath string) error {
	out, err := os.Create(filePath)
	if err != nil {
		log.Println("Error to save file, error: " + err.Error())
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}
	return nil
}

func removeSpecialChars(str string) string {
    re := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
    str = re.ReplaceAllString(str, "")
    
    str = regexp.MustCompile(`\s+`).ReplaceAllString(str, "_")

    return str
}

var mimeToExt = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/gif":  ".gif",
	"text/plain": ".txt",
	"application/pdf": ".pdf",
	"application/zip": ".zip",
	// Adicione mais tipos MIME e suas extensões conforme necessário
}

func GetFileType(file multipart.File) (string, error) {
	buf := make([]byte, 512)
	_, err := file.Read(buf)
	if err != nil {
		return "", err
	}

	filetype := http.DetectContentType(buf)

	// Retorne a extensão correspondente ou uma mensagem de erro se o tipo não for encontrado
	ext, ok := mimeToExt[filetype]
	if !ok {
		return "", fmt.Errorf("tipo de arquivo não reconhecido: %s", filetype)
	}

	// Retorne a extensão do arquivo
	file.Seek(0, io.SeekStart) // Reposiciona o ponteiro do arquivo para o início
	return ext, nil
}