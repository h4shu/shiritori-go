package models

type (
	WordchainGetLastRequest  struct{}
	WordchainGetLastResponse struct {
		Word string `json:"word"`
	}
	WordchainListRequest struct {
		Limit int `json:"limit"`
	}
	WordchainListResponse struct {
		Wordchain []string `json:"wordchain"`
		Len       int      `json:"len,string"`
	}
	WordchainAppendRequest struct {
		Word string `json:"word" form:"word"`
	}
	WordchainAppendResponse struct{}
)
