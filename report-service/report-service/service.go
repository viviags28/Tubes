package main

import (
	"context"
	"database/sql"
)

type ReportService struct {
	db *sql.DB
}

func NewReportService(db *sql.DB) *ReportService {
	return &ReportService{
		db: db,
	}
}

func (s *ReportService) GetDailyReport(ctx context.Context, date string) (*DailyReport, error) {

	// sementara dummy data dulu
	return &DailyReport{
		TotalPaket:  10,
		Delivered:   8,
		Pending:     1,
		Terlambat:   1,
		RataRataETA: 2.5,
	}, nil
}