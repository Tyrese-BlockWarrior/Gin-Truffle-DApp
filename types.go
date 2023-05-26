package main

type Details struct {
	From     string `json:"from"`
	Contract string `json:"contract"`
}

type InsertCertificate struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Course string `json:"course"`
	Grade  string `json:"grade"`
	Date   string `json:"date"`
}

type ReturnCertificate struct {
	Name   string
	Course string
	Grade  string
	Date   string
}
