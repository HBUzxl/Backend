package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUnsubmitCasesHandler 获取未提交的病例
// @Summary      获取未提交的病例
// @Description  获取所有未提交的病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Success      200      {object}  map[string][]models.Case
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/unsubmitted [post]
// @Security     Bearer
func GetUnsubmitCasesHandler(c *gin.Context) {
	cases, err := services.GetUnsubmitCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Unsubmit Cases " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cases": cases,
	})
}

// GetPendingDiagnosisCasesHandler 获取待诊断的病例
// @Summary      获取待诊断的病例
// @Description  获取所有待诊断的病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Success      200      {object}  map[string][]models.Case
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/pendingdiagnosis [post]
// @Security     Bearer
func GetPendingDiagnosisCasesHandler(c *gin.Context) {
	cases, err := services.GetPendingDiagnosisCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Pending Diagnosis Cases " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cases": cases,
	})
}

// GetDiagnosedCasesHandler 获取已诊断的病例
// @Summary      获取已诊断的病例
// @Description  获取所有已诊断的病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Success      200      {object}  map[string][]models.Case
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/diagnosed [post]
// @Security     Bearer
func GetDiagnosedCasesHandler(c *gin.Context) {
	cases, err := services.GetDiagnosedCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Diagnosed Cases " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cases": cases,
	})
}

// GetReturnedCasesHandler 获取已退回的病例
// @Summary      获取已退回的病例
// @Description  获取所有已退回的病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Success      200      {object}  map[string][]models.Case
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/returned [post]
// @Security     Bearer
func GetReturnedCasesHandler(c *gin.Context) {
	cases, err := services.GetReturnedCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Returned Cases " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cases": cases,
	})
}

// GetWithdrawCasesHandler 获取已撤回的病例
// @Summary      获取已撤回的病例
// @Description  获取所有已撤回的病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Success      200      {object}  map[string][]models.Case
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/withdraw [post]
// @Security     Bearer
func GetWithdrawCasesHandler(c *gin.Context) {
	cases, err := services.GetWithdrawCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Withdraw Cases " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cases": cases,
	})
}

// GetAllCasesHandler 获取所有病例
// @Summary      获取所有病例
// @Description  获取所有病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Success      200      {object}  map[string][]models.Case
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/all [get]
// @Security     Bearer
func GetAllCasesHandler(c *gin.Context) {
	cases, err := services.GetAllCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get All Cases " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cases": cases,
	})
}

// GetCaseByCaseIDHandler 根据病例ID获取病例
// @Summary      根据病例ID获取病例
// @Description  根据病例ID获取病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        caseID  path      string  true  "病例ID"
// @Success      200      {object}  map[string]models.Case
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/{caseID} [get]
// @Security     Bearer
func GetCaseByCaseIDHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}

	caseData, err := services.GetCaseByCaseID(caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get Case By CaseID " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"case": caseData})
}

// UpdatePendingCaseHandler 更新状态：到待诊断
// @Summary      更新状态：到待诊断
// @Description  更新状态：到待诊断
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        caseID  path      string  true  "病例ID"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/toPendingdiagnosis/{caseID} [post]
// @Security     Bearer
func UpdatePendingCaseHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}
	err := services.UpdateCaseToPendingDiagnosis(caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update Case Status " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Case Status Success"})
}

// UpdateDiagnosedCaseHandler 更新状态：到已诊断
// @Summary      更新状态：到已诊断
// @Description  更新状态：到已诊断
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        caseID  path      string  true  "病例ID"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/toDiagnosed/{caseID} [post]
// @Security     Bearer
func UpdateDiagnosedCaseHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}
	err := services.UpdateCaseToDiagnosed(caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update Case Status " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Case Status Success"})
}

// UpdateReturnedCaseHandler 更新状态：到被退回
// @Summary      更新状态：到被退回
// @Description  更新状态：到被退回
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        caseID  path      string  true  "病例ID"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/toReturned/{caseID} [post]
// @Security     Bearer
func UpdateReturnedCaseHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}
	err := services.UpdateCaseToReturned(caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update Case Status " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Case Status Success"})
}

// UpdateWithdrawCaseHandler 更新状态：到撤回
// @Summary      更新状态：到撤回
// @Description  更新状态：到撤回
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        caseID  path      string  true  "病例ID"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/toWithdraw/{caseID} [post]
// @Security     Bearer
func UpdateWithdrawCaseHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}
	err := services.UpdateCaseToWithdraw(caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update Case Status " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Case Status Success"})
}

// IncreasePrintCountHandler 增加打印次数
// @Summary      增加打印次数
// @Description  增加打印次数
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        caseID  path      string  true  "病例ID"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/{caseID}/print [post]
// @Security     Bearer
func IncreasePrintCountHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}
	err := services.IncreasePrintCount(caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Increase Print Count " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Increase Print Count Success"})
}

// SubmitCaseHandler 提交病例
// @Summary      提交病例
// @Description  提交病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        caseData  body      models.Case  true  "病例数据"
// @Success      200      {object}  map[string]models.Case
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/submit [post]
// @Security     Bearer
func SubmitCaseHandler(c *gin.Context) {
	caseData := models.Case{}
	err := c.BindJSON(&caseData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bind Case Data Error: " + err.Error()})
		return
	}
	err = services.SubmitCase(&caseData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Submit Case Error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Submit Case Success", "case": caseData})

}

// ExportExcelHandler 导出病例Excel
// @Summary      导出病例Excel
// @Description  导出病例Excel
// @Tags         cases
// @Accept       json
// @Produce      json
// @Success      200      {object}  map[string]string
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/excel [get]
// @Security     Bearer
func ExportExcelHandler(c *gin.Context) {
	excelData, err := services.ExportExcel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Export Excel Error: " + err.Error()})
		return
	}
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=统计报表.xlsx")
	c.Data(http.StatusOK, "application/octet-stream", excelData)
}

// DeleteCaseHandler 删除病例
// @Summary      删除病例
// @Description  删除病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        caseID  path      string  true  "病例ID"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  middleware.ErrorResponse "错误响应"
// @Failure      500      {object}  middleware.ErrorResponse "错误响应"
// @Router       /api/case/{caseID} [delete]
// @Security     Bearer
func DeleteCaseHandler(c *gin.Context) {
	caseID := c.Param("caseID")
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing caseID"})
		return
	}
	err := services.DeleteCase(caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete Case Error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete Case Success"})
}
