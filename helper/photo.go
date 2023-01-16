package helper

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"simple-social-media-API/config"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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

func UploadProfilePhotoS3(file multipart.FileHeader, email string) (string, error) {
	s3Session := config.S3Config()
	uploader := s3manager.NewUploader(s3Session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	ext := filepath.Ext(file.Filename)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("socmedapibucket"),
		Key:    aws.String("files/user/" + email + "/profile-photo" + ext),
		Body:   src,
	})
	path := "files/user/" + email + "/profile-photo" + ext
	return path, nil
}

func UploadBackgroundPhotoS3(file multipart.FileHeader, email string) (string, error) {
	s3Session := config.S3Config()
	uploader := s3manager.NewUploader(s3Session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	ext := filepath.Ext(file.Filename)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("socmedapibucket"),
		Key:    aws.String("files/user/" + email + "/background-photo" + ext),
		Body:   src,
	})
	path := "files/user/" + email + "/background-photo" + ext
	return path, nil
}

func UploadPostPhotoS3(file multipart.FileHeader, userID int) (string, error) {
	s3Session := config.S3Config()
	uploader := s3manager.NewUploader(s3Session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	ext := filepath.Ext(file.Filename)

	cnv := strconv.Itoa(userID)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("socmedapibucket"),
		Key:    aws.String("files/user/" + cnv + "/post-photo" + ext),
		Body:   src,
	})
	path := "files/user/" + cnv + "/post-photo" + ext
	return path, nil
}
