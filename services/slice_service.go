package services

import (
	"backend/config"
	"backend/models"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// UploadSlice 上传切片
func UploadSlice(file *multipart.FileHeader, sliceID string, caseID string) (*models.Slice, error) {
	// 1. 创建上传目录结构
	caseDir := filepath.Join("uploads", "slices", "case_"+caseID)
	originalDir := filepath.Join(caseDir, "original")
	if err := os.MkdirAll(originalDir, 0755); err != nil {
		return nil, fmt.Errorf("创建上传目录失败: %w", err)
	}

	// 2. 保持原始文件名，这样便于后续处理
	filename := file.Filename

	// 获取绝对路径用于文件存储
	absOriginalDir, err := filepath.Abs(originalDir)
	if err != nil {
		return nil, fmt.Errorf("获取绝对路径失败: %w", err)
	}

	// 原始文件的绝对存储路径
	filePath := filepath.Join(absOriginalDir, filename)

	// URL路径（使用正斜杠，不使用系统路径分隔符）
	fileUrl := fmt.Sprintf("/uploads/slices/case_%s/original/%s", caseID, filename)

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
	slice := &models.Slice{
		SliceID:  sliceID,
		FileName: file.Filename,
		FilePath: filePath, // 存储原始文件的绝对路径
		FileSize: file.Size,
		CaseID:   caseID,
		FileUrl:  fileUrl,   // 存储原始文件的URL路径
		Status:   "pending", // 设置初始状态为待处理
	}

	// 5. 保存到数据库
	if err := config.DB.Create(slice).Error; err != nil {
		// 如果数据库保存失败，删除已上传的文件
		os.Remove(filePath)
		return nil, fmt.Errorf("保存切片记录失败: %w", err)
	}

	// 6. 异步转换为DZI格式
	go func() {
		if err := ConvertSVStoDZI(slice); err != nil {
			// 转换失败，更新状态
			slice.Status = "convert_failed"
			config.DB.Save(slice)
			fmt.Printf("切片转换失败: %v\n", err)
		}
	}()

	return slice, nil
}

// ConvertSVStoDZI 将SVS格式切片转换为DZI格式
func ConvertSVStoDZI(slice *models.Slice) error {
	// 1. 更新状态为转换中
	slice.Status = "converting"
	if err := config.DB.Save(slice).Error; err != nil {
		return fmt.Errorf("更新状态失败: %w", err)
	}

	// 2. 创建DZI目录结构
	caseDir := filepath.Join("uploads", "slices", "case_"+slice.CaseID)
	dziDir := filepath.Join(caseDir, "dzi")
	if err := os.MkdirAll(dziDir, 0755); err != nil {
		return fmt.Errorf("创建DZI目录失败: %w", err)
	}

	// 3. 获取DZI输出目录的绝对路径
	absDZIDir, err := filepath.Abs(dziDir)
	if err != nil {
		return fmt.Errorf("获取DZI目录绝对路径失败: %w", err)
	}

	// 4. 准备文件名
	baseFileName := strings.TrimSuffix(slice.FileName, filepath.Ext(slice.FileName))
	dziFileName := baseFileName
	dziPath := filepath.Join(absDZIDir, dziFileName)

	// 5. 执行转换
	cmd := exec.Command("vips", "dzsave", slice.FilePath, dziPath)
	fmt.Println(cmd.Args)
	if output, err := cmd.CombinedOutput(); err != nil {
		slice.Status = "convert_failed"
		config.DB.Save(slice)
		return fmt.Errorf("转换失败: %s, %w", string(output), err)
	}

	// 6. 更新切片记录
	dziUrl := fmt.Sprintf("/uploads/slices/case_%s/dzi/%s", slice.CaseID, dziFileName)
	slice.DZIPath = dziPath
	slice.DZIUrl = dziUrl
	slice.Status = "converted"

	// 7. 保存更新到数据库
	if err := config.DB.Save(slice).Error; err != nil {
		return fmt.Errorf("更新切片记录失败: %w", err)
	}

	return nil
}

// // GetSlicesByCaseID 获取病例的所有切片
// func GetSlicesByCaseID(caseID uint) ([]models.Slice, error) {
// 	var slices []models.Slice
// 	err := config.DB.Where("case_id = ?", caseID).Find(&slices).Error
// 	if err != nil {
// 		return nil, fmt.Errorf("获取切片列表失败: %w", err)
// 	}
// 	return slices, nil
// }

// // DeleteSlice 删除切片
// func DeleteSlice(sliceID uint) error {
// 	var slice models.Slice
// 	if err := config.DB.First(&slice, sliceID).Error; err != nil {
// 		return fmt.Errorf("切片不存在: %w", err)
// 	}

// 	// 删除文件
// 	if err := os.Remove(slice.FilePath); err != nil {
// 		return fmt.Errorf("删除文件失败: %w", err)
// 	}

// 	// 删除数据库记录
// 	if err := config.DB.Delete(&slice).Error; err != nil {
// 		return fmt.Errorf("删除切片记录失败: %w", err)
// 	}

// 	return nil
// }
