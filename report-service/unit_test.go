package main

import (
	"context"
	"database/sql"
	"testing"
)

func TestGetDailyReport_ShouldReturnReport(t *testing.T) {

	svc := NewReportService(&sql.DB{})

	report, err := svc.GetDailyReport(
		context.Background(),
		"2026-04-25",
	)

	if report == nil {
		t.Errorf("Expected report not nil, got nil")
	}

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
