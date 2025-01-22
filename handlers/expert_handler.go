package handlers

import (
	"backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetExpertsHandler handles getting experts
// @Summary      获取专家列表
// @Description  获取所有专家
// @Tags         experts
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string][]models.Expert
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/expert [get]
// @Security     Bearer
func GetExpertsHandler(c *gin.Context) {
	experts, err := services.GetExperts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get Experts " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"experts": experts,
	})
}

// GetPendingCasesByExpertUsernameHandler handles getting pending cases by expert username
// @Summary      获取专家待处理的病例
// @Description  根据专家用户名获取所有待处理的病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        username  path      string  true  "专家用户名"
// @Success      200  {object}  map[string][]models.Case
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/case/pending/{username} [get]
// @Security     Bearer
func GetPendingCasesByExpertUsernameHandler(c *gin.Context) {
	username := c.Param("username")
	fmt.Println(username)
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing expertID"})
		return
	}
	cases, err := services.GetPendingCasesByExpertUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get Pending Cases By ExpertID " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cases": cases})
}

// GetDiagnosedCasesByExpertUsernameHandler handles getting diagnosed cases by expert username
// @Summary      获取专家已诊断的病例
// @Description  根据专家用户名获取所有已诊断的病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        username  path      string  true  "专家用户名"
// @Success      200  {object}  map[string][]models.Case
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/case/diagnosed/{username} [get]
// @Security     Bearer
func GetDiagnosedCasesByExpertUsernameHandler(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing expertID"})
		return
	}
	cases, err := services.GetDiagnosedCasesByExpertUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get Diagnosed Cases By ExpertID " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cases": cases})
}

// GetReturnedCasesByExpertUsernameHandler handles getting returned cases by expert username
// @Summary      获取专家已退回的病例
// @Description  根据专家用户名获取所有已退回的病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        username  path      string  true  "专家用户名"
// @Success      200  {object}  map[string][]models.Case
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/case/returned/{username} [get]
// @Security     Bearer
func GetReturnedCasesByExpertUsernameHandler(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing expertID"})
		return
	}
	cases, err := services.GetReturnedCasesByExpertUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get Returned Cases By ExpertID " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cases": cases})
}

// GetAllCasesByExpertUsernameHandler handles getting all cases by expert username
// @Summary      获取专家所有病例
// @Description  根据专家用户名获取所有病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        username  path      string  true  "专家用户名"
// @Success      200  {object}  map[string][]models.Case
// @Failure      500  {object}  map[string]string
// @Router       /api/expert/all/{username} [get]
// @Security     Bearer
func GetAllCasesByExpertUsernameHandler(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing expertID"})
		return
	}
	cases, err := services.GetAllCasesByExpertUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get All Cases By ExpertID " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cases": cases})
}

// GetWithdrawCasesByExpertUsernameHandler handles getting withdraw cases by expert username
// @Summary      获取专家已撤回的病例
// @Description  根据专家用户名获取所有已撤回的病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        username  path      string  true  "专家用户名"
// @Success      200  {object}  map[string][]models.Case
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/case/withdraw/{username} [get]
// @Security     Bearer
func GetWithdrawCasesByExpertUsernameHandler(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing expertID"})
		return
	}
	cases, err := services.GetWithdrawCasesByExpertUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get Withdraw Cases By ExpertID " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cases": cases})
}

// ExportExcelCasesByUsernameHandler handles exporting excel cases by username
// @Summary      导出专家所有病例
// @Description  根据专家用户名导出所有病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        username  path      string  true  "专家用户名"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/expert/excel/{username} [get]
// @Security     Bearer
func ExportExcelCasesByUsernameHandler(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing expertID"})
		return
	}
	excelData, err := services.ExportExcelCasesByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Export Excel Cases By ExpertID " + err.Error()})
		return
	}
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=统计报表.xlsx")
	c.Data(http.StatusOK, "application/octet-stream", excelData)
}

// GetAppointmentsByUsernameHandler handles getting appointments by username
// @Summary      获取专家所有预约
// @Description  根据专家用户名获取所有预约
// @Tags         appointments
// @Accept       json
// @Produce      json
// @Param        username  path      string  true  "专家用户名"
// @Success      200  {object}  map[string][]models.Appointment
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/expert/{username}/appointments [get]
// @Security     Bearer
func GetAppointmentsByUsernameHandler(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing expertID"})
		return
	}
	appointments, err := services.GetAppointmentsByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get Appointments By ExpertID " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"appointments": appointments})
}

type DiagnoseCaseRequest struct {
	CaseID                 string `json:"caseID"`
	ExpertDiagnosisOpinion string `json:"expertDiagnosisOpinion"`
	DiagnosisContent       string `json:"diagnosisContent"`
	DiagnosisRemarks       string `json:"diagnosisRemarks"`
	MirrorDescription      string `json:"mirrorDescription"`
}

// DiagnoseCaseHandler handles diagnosing a case
// @Summary      诊断病例
// @Description  诊断病例
// @Tags         cases
// @Accept       json
// @Produce      json
// @Param        caseID  path      string  true  "病例ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/expert/diagnose [post]
// @Security     Bearer
func DiagnoseCaseHandler(c *gin.Context) {
	var req DiagnoseCaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 打印接收到的请求数据
	fmt.Printf("Received request data: %+v\n", req)

	err := services.DiagnoseCase(req.CaseID, req.ExpertDiagnosisOpinion, req.DiagnosisContent, req.DiagnosisRemarks, req.MirrorDescription)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Diagnose Case " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Diagnose Case Success"})
}

// GetCaseCounts 获取专家的各类病例数量
func GetCaseCounts(c *gin.Context) {
	expertID := c.GetString("userID")
	if expertID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	counts, err := services.GetCaseCounts(expertID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, counts)
}
