package dto

type MetaResponse struct {
	TotalCount int    `json:"total_count,omitempty"`
	Limit      int    `json:"limit"`
	NextCursor string `json:"next_cursor"`
}

type GeneralListResponse struct {
	Message string        `json:"message"`
	Data    interface{}   `json:"data"`
	Meta    *MetaResponse `json:"meta"`
}

type GeneralResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
