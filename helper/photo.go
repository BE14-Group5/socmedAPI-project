package helper

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadProfilePhoto(file multipart.FileHeader, email string) (string, error) {
	src, _ := file.Open()
	defer src.Close()

	dir, _ := os.Getwd()
	os.MkdirAll(dir+"/files/user/"+email, 0770)
	dir = dir + "/files/user/" + email
	ext := filepath.Ext(file.Filename)
	dir = filepath.Join(dir, filepath.Base("profile-photo"+ext))

	dst, err := os.Create(dir)
	if err != nil {
		return "", errors.New("server problem")
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return "", errors.New("server problem")
	}
	return dir, nil
}

func UploadBackgroundPhoto(file multipart.FileHeader, email string) (string, error) {
	src, _ := file.Open()
	defer src.Close()

	dir, _ := os.Getwd()
	os.MkdirAll(dir+"/files/user/"+email, 0770)
	dir = dir + "/files/user/" + email
	ext := filepath.Ext(file.Filename)
	dir = filepath.Join(dir, filepath.Base("background-photo"+ext))

	dst, err := os.Create(dir)
	if err != nil {
		return "", errors.New("server problem")
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return "", errors.New("server problem")
	}
	return dir, nil
}
