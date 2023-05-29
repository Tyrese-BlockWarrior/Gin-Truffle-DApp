package types

type JsonCertificate struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Course string `json:"course"`
	Grade  string `json:"grade"`
	Date   string `json:"date"`
}

type FormCertificate struct {
	ID     string `form:"id" binding:"required"`
	Name   string `form:"name" binding:"required"`
	Course string `form:"course" binding:"required"`
	Grade  string `form:"grade" binding:"required"`
	Date   string `form:"date" binding:"required"`
}

type ReturnCertificate struct {
	Name   string
	Course string
	Grade  string
	Date   string
}
