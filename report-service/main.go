package main

import (
	"log"
	"net/http"
)

func main() {

	ConnectDB()

	repo := MySQLReportRepository{db: DB}

	svc := NewReportService(repo)
	h := NewReportHandler(svc)

	http.HandleFunc("/report/daily", h.DailyReport)

	log.Println("Report service running on :8083")

	log.Fatal(http.ListenAndServe(":8083", nil))
}
