package model

type WebResponse[T any] struct {
	Data     T            `json:"data"`
	Metadata *QueryParams `json:"metadata,omitempty"`
}

type QueryParams struct {
	Page  int     `json:"page"`
	Limit int     `json:"limit"`
	Nama  *string `json:"nama,omitempty"`
}
