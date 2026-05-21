package main

import (
	"log"
	"net/http"
)

type MySQLReportRepository struct {
	db *sql.DB
}

func (r MySQLReportRepository) GetDailyReport(ctx context.Context, date string) (*DailyReport, error) {
	return &DailyReport{
		TotalPaket:  10,
		Delivered:   8,
		Pending:     1,
		Terlambat:   1,
		RataRataETA: 2.5,
	}, nil
}