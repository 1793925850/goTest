package upload

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"blog-service/global"
	"blog-service/pkg/util"
)

/**
上传文件的工具库
*/

type FileType int

const TypeImage FileType = iota + 1

// GetFileName 获取文件名
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)

	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// GetFileExt 获取文件类型（后缀）
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetSavePath 获取文件保存路径
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// GetServerUrl 获取文件服务地址
func GetServerUrl() string {
	return global.AppSetting.UploadServerUrl
}

// CheckSavePath 检查保存目录是否存在
func CheckSavePath(dst string) bool {
	// Stat 获取文件的描述信息
	_, err := os.Stat(dst)

	return os.IsNotExist(err)
}

// CheckContainExt 检查文件后缀是否包含在约定的后缀配置项中
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)

	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}

	return false
}

// CheckMaxSize 检查文件大小是否超出限制，超出限制返回 true
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := io.ReadAll(f)
	size := len(content)

	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

// CheckPermission 检查文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)

	return os.IsPermission(err)
}

// CreateSavePath 创建保存上传文件的目录
func CreateSavePath(dst string, perm os.FileMode) error {
	// MkdirAll 通过传入的 perm（权限位），递归创建所需的所有目录结构
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

// SaveFile 保存上传的文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	// Open 打开源地址文件
	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	// Create 创建目标地址为 dst 的文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer out.Close()

	// Copy 进行两者之间文件内容的拷贝，即将 src 的内容写到 out 中
	_, err = io.Copy(out, src)
	return err
}
