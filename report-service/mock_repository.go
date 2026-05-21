package main

type MockReportRepo struct{}

func (m MockReportRepo) GetDailyReport(date string) (DailyReport, error) {
	return DailyReport{
		TotalPaket: 10,
	}, nil
}
