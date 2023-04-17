package service

import (
	"errors"
	"mime/multipart"
	"os"

	"blog-service/global"
	"blog-service/pkg/upload"
)

/**
业务层，上传文件服务
*/

// FileInfo 文件信息
type FileInfo struct {
	Name      string
	AccessUrl string
}

// UploadFile 上传文件
func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permission.")
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName

	return &FileInfo{
		Name:      fileName,
		AccessUrl: accessUrl,
	}, nil
}
