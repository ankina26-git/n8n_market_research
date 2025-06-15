package utils

import "time"

// YYYY-MM-DD形式の日付を返す
func Today() string {
	return time.Now().Format("2006-01-02")
}

// 昨日の日付（文字列）を返す
func Yesterday() string {
	return time.Now().AddDate(0, 0, -1).Format("2006-01-02")
}

// 任意の差分日数でフォーマット返す
func DaysAgo(n int) string {
	return time.Now().AddDate(0, 0, -n).Format("2006-01-02")
}

// 指定の期間範囲を取得（過去7日間など）
func DateRange(days int) (start, end string) {
	end = Yesterday()
	start = time.Now().AddDate(0, 0, -days).Format("2006-01-02")
	return
}
