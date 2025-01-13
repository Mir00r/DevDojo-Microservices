package dtos

type APIResponse struct {
	Error      bool        `json:"error"`
	Code       int         `json:"code"`
	CodeStatus string      `json:"codeStatus,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Metadata   interface{} `json:"metadata,omitempty"`
}

type PaginationMetadata struct {
	CurrentPage int `json:"currentPage"`
	TotalPages  int `json:"totalPages"`
}
