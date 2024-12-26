package handlers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadSliceHandler 处理切片上传
// @Summary      上传切片
// @Description  上传切片
// @Tags         slices
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "切片文件"
// @Param        caseID  formData  string  true  "病例ID"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/slice/upload [post]
// @Security     Bearer
func UploadSliceHandler(c *gin.Context) {
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

	// 3. 获取切片号
	SliceID := file.Filename
	if SliceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "切片号不能为空",
		})
		return
	}
	SliceID = caseID + "_" + SliceID

	// 4. 调用服务上传切片
	slice, err := services.UploadSlice(file, SliceID, caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "上传切片失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "上传成功",
		"slice":   slice,
	})
}

// // GetSlicesHandler 获取病例的所有切片
// func GetSlicesHandler(c *gin.Context) {
// 	// 获取病例ID
// 	caseIDStr := c.Param("caseID")
// 	if caseIDStr == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "病例ID不能为空",
// 		})
// 		return
// 	}
// 	caseID, err := strconv.ParseUint(caseIDStr, 10, 32)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "病例ID格式错误",
// 		})
// 		return
// 	}

// 	// 获取切片列表
// 	slices, err := services.GetSlicesByCaseID(uint(caseID))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "获取切片列表失败: " + err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"slices": slices,
// 	})
// }

// // DeleteSliceHandler 删除切片
// func DeleteSliceHandler(c *gin.Context) {
// 	// 获取切片ID
// 	sliceIDStr := c.Param("sliceID")
// 	if sliceIDStr == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "切片ID不能为空",
// 		})
// 		return
// 	}
// 	sliceID, err := strconv.ParseUint(sliceIDStr, 10, 32)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "切片ID格式错误",
// 		})
// 		return
// 	}

// 	// 删除切片
// 	if err := services.DeleteSlice(uint(sliceID)); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "删除切片失败: " + err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "删除成功",
// 	})
// }
