package middleware

// ErrorResponse 定义统一的错误响应格式
type ErrorResponse struct {
	Error string `json:"error" example:"错误信息"`
}
