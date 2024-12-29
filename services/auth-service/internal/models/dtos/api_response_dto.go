package dtos

type APIResponse struct {
	Error    bool        `json:"error"`
	Code     int         `json:"code"`
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
}

type PaginationMetadata struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}
