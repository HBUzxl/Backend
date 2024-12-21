package services

import (
	"backend/config"
	"backend/models"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// UploadAttachment 上传附件
func UploadAttachment(file *multipart.FileHeader, caseID string) (*models.Attachment, error) {

	// 1. 创建上传目录
	uploadDir := filepath.Join("uploads", "attachments", "case_"+caseID)
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, fmt.Errorf("创建上传目录失败: %w", err)
	}

	// 2. 生成文件名
	filename := file.Filename
	filePath := filepath.Join(uploadDir, filename)

	// 3. 保存文件
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("打开上传文件失败: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("创建目标文件失败: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		// 如果复制失败，删除已创建的文件
		os.Remove(filePath)
		return nil, fmt.Errorf("保存文件失败: %w", err)
	}

	// 4. 创建切片记录
	attachment := &models.Attachment{
		FileName: file.Filename,
		FilePath: filePath,
		FileSize: file.Size,
		CaseID:   caseID,
	}

	// 5. 保存到数据库
	if err := config.DB.Create(attachment).Error; err != nil {
		// 如果数据库保存失败，删除已上传的文件
		os.Remove(filePath)
		return nil, fmt.Errorf("保存附件记录失败: %w", err)
	}

	return attachment, nil
}
