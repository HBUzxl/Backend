package handlers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadAttachmentHandler 处理附件上传
// @Summary      上传附件
// @Description  上传附件
// @Tags         attachments
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "附件文件"
// @Param        caseID  formData  string  true  "病例ID"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/attachment/upload [post]
// @Security     Bearer
func UploadAttachmentHandler(c *gin.Context) {
	// 1. 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "获取上传文件失败: " + err.Error(),
		})
		return
	}

	// 2. 获取病例ID
	caseID := c.PostForm("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "病例ID不能为空",
		})
		return
	}

	// 3. 调用服务上传附件
	attachments, err := services.UploadAttachment(file, caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "上传附件失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "上传成功",
		"attachment": attachments,
	})
}
