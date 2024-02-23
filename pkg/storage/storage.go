package storage

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 * 50 // 50MB

var defaultExtesions = []string{".png", ".jpeg", ".pdf", ".jpg", ".zip", ".rar"}

type File struct {
	BaseDir string
}

var (
	ErrFileDoesntExists       = errors.New("file doesn't exists")
	ErrFileNotAllowedExtesion = errors.New("file not allowed extension")
	ErrUniqueFileGeneration   = errors.New("unique file generation")
	ErrFilenameCollision      = errors.New("filename collision")
	ErrFileNotOpen            = errors.New("file not open")
	ErrFileNotCreated         = errors.New("file not created")
)

type UploadFile struct {
	File               multipart.File
	Header             *multipart.FileHeader
	MainDir            string
	SubDir             string
	GenerateUniqueName bool     // false originalname or true uniqueName
	AllowedExtensions  []string // default nil or custom  {".png",".jpeg"}
}

func NewFileManager(baseDir string) *File {

	err := os.MkdirAll(baseDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return &File{
		BaseDir: baseDir,
	}
}

func (f *File) Upload(
	fiilePart multipart.File,
	fileHeader *multipart.FileHeader,
	mainDir string,
	subDir string,
	generateUniqueName bool,
	allowedExtensions []string,
) (string, error) {
	var arg = UploadFile{
		File:               fiilePart,
		Header:             fileHeader,
		MainDir:            mainDir,
		SubDir:             subDir,
		GenerateUniqueName: generateUniqueName,
		AllowedExtensions:  allowedExtensions,
	}

	// /files/mainDir/subDir/fileName

	arg.Header.Filename = strings.ToLower(arg.Header.Filename)

	// dosya uzantısı whitelist'de mi diye kontrol et
	if err := f.acceptFile(filepath.Ext(arg.Header.Filename), arg.AllowedExtensions); err != nil {
		return "", err
	}

	fileName := arg.Header.Filename
	// benzersiz bir isimlendirme olacak mı?
	if arg.GenerateUniqueName {
		// benzeriz bir dosya adı olustur
		fileUUID, err := uuid.NewRandom()
		if err != nil {
			return "", ErrUniqueFileGeneration
		}
		fileName = fileUUID.String() + filepath.Ext(arg.Header.Filename)
	}

	err := os.MkdirAll(f.BaseDir+arg.MainDir+arg.SubDir, os.ModePerm)
	if err != nil {

		return "", ErrFilenameCollision
	}

	outLocation := filepath.Join(f.BaseDir, arg.MainDir, arg.SubDir, fileName)

	_, err = os.Stat(outLocation)
	if err == nil {
		err = os.RemoveAll(outLocation)
		if err != nil {
			return "", ErrFilenameCollision
		}
	}

	fileOs, err := os.Create(outLocation)
	if err != nil {

		return "", ErrFileNotOpen
	}
	if _, err := io.Copy(fileOs, arg.File); err != nil {

		return "", ErrFileNotCreated
	}

	return outLocation, nil
}

func (f *File) Get(mainDir, subDir, fileName string) (str string) {
	return filepath.Join(f.BaseDir, mainDir, subDir, fileName)
}

func (f *File) Delete(fileUUID string) (err error) {

	fileLocation := filepath.Join(f.BaseDir, fileUUID)

	if err := os.Remove(fileLocation); err != nil {

		return errors.New("file remove")
	}
	return nil
}

func (f *File) acceptFile(fileName string, extensions []string) error {

	// Dosya uzantısını kontrol et

	allowedExtensions := defaultExtesions // Kabul edilen uzantılar
	if extensions != nil {
		allowedExtensions = extensions // Kabul edilen uzantılar
	}

	fileExt := filepath.Ext(fileName)

	if !isValidExtension(fileExt, allowedExtensions) {

		return ErrFileNotAllowedExtesion
	}

	return nil

}

func isValidExtension(ext string, allowed []string) bool {
	for _, validExt := range allowed {
		if ext == validExt {
			return true
		}
	}
	return false
}
