package model

type WebResponse[T any] struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message,omitempty"`
	Data       T      `json:"data"`
	Metadata   any    `json:"metadata,omitempty"`
}

type QueryParams struct {
	Page  int     `json:"page"`
	Limit int     `json:"limit"`
	Nama  *string `json:"nama,omitempty"`
}
