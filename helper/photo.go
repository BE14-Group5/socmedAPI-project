package helper

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"simple-social-media-API/config"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var ObjectURL string = "https://socmedapibucket.s3.ap-southeast-1.amazonaws.com/"

// func UploadProfilePhoto(file multipart.FileHeader, email string) (string, error) {
// 	src, _ := file.Open()
// 	defer src.Close()

// 	dir, _ := os.Getwd()
// 	os.MkdirAll(dir+"/files/user/"+email, 0770)
// 	dir = dir + "/files/user/" + email
// 	ext := filepath.Ext(file.Filename)
// 	dir = filepath.Join(dir, filepath.Base("profile-photo"+ext))

// 	dst, err := os.Create(dir)
// 	if err != nil {
// 		return "", errors.New("server problem")
// 	}
// 	defer dst.Close()
// 	if _, err = io.Copy(dst, src); err != nil {
// 		return "", errors.New("server problem")
// 	}
// 	return dir, nil
// }

// func UploadBackgroundPhoto(file multipart.FileHeader, email string) (string, error) {
// 	src, _ := file.Open()
// 	defer src.Close()

// 	dir, _ := os.Getwd()
// 	os.MkdirAll(dir+"/files/user/"+email, 0770)
// 	dir = dir + "/files/user/" + email
// 	ext := filepath.Ext(file.Filename)
// 	dir = filepath.Join(dir, filepath.Base("background-photo"+ext))

// 	dst, err := os.Create(dir)
// 	if err != nil {
// 		return "", errors.New("server problem")
// 	}
// 	defer dst.Close()
// 	if _, err = io.Copy(dst, src); err != nil {
// 		return "", errors.New("server problem")
// 	}
// 	return dir, nil
// }

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
	if err != nil {
		return "", errors.New("problem with upload profile photo")
	}
	path := ObjectURL + "files/user/" + email + "/profile-photo" + ext
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
	if err != nil {
		return "", errors.New("problem with upload profile photo")
	}
	path := ObjectURL + "files/user/" + email + "/background-photo" + ext
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
	// ext := filepath.Ext(file.Filename)

	cnv := strconv.Itoa(userID)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("socmedapibucket"),
		Key:    aws.String("files/post/" + cnv + "/" + file.Filename),
		Body:   src,
	})
	if err != nil {
		return "", errors.New("problem with upload post photo")
	}
	path := ObjectURL + "files/post/" + cnv + "/" + file.Filename
	return path, nil
}

// func DownloadProfilePhoto(path string) multipart.FileHeader {
// 	s3Session := config.S3Config()
// 	downloader := s3manager.NewDownloader(s3Session)

// 	profilePhoto, err := os.Create("profile_photo.jpg")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	_, err = downloader.Download(profilePhoto, &s3.GetObjectInput{
// 		Bucket: aws.String("socmedapibucket"),
// 		Key:    aws.String(path),
// 	})
// 	if err != nil {
// 		log.Println("error download profile photo: ", err.Error())
// 	}

// 	return profilePhoto
// }
