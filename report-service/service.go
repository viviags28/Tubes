package main

import "context"

type ReportRepository interface {
	GetDailyReport(ctx context.Context, date string) (*DailyReport, error)
}

type ReportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetDailyReport(ctx context.Context, date string) (*DailyReport, error) {

	// sementara masih dummy (boleh tetap gini buat functional test)
	return &DailyReport{
		TotalPaket:  10,
		Delivered:   8,
		Pending:     1,
		Terlambat:   1,
		RataRataETA: 2.5,
	}, nil
}