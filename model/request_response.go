package model

type Request struct {
	Page     int    `json:"page,omitempty" form:"page"`
	PageSize int    `json:"page_size,omitempty" form:"page_size"`
	Status   string `json:"status,omitempty" form:"status"`         // Filter berdasarkan status
	Start    string `json:"start_date,omitempty" form:"start_date"` // Tanggal awal (format: YYYY-MM-DD)
	End      string `json:"end_date,omitempty" form:"end_date"`
}

type Response struct {
	Success bool        `json:"success`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ResponseView struct {
	Success bool        `json:"success`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Message string      `json:"message,omitempty"`
}

func NewSuccessResponse(message string, data any) Response {
	return Response{
		Success: true,
		Data:    data,
		Message: message,
	}
}
func NewSuccessResponseView(message string, data any, meta any) ResponseView {
	return ResponseView{
		Success: true,
		Data:    data,
		Meta:    meta,
		Message: message,
	}
}

func NewFailedResponse(message string) Response {
	return Response{
		Message: message,
	}
}
