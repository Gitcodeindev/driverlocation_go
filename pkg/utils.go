package pkg

import "time"

// форматирует время в стандартный вид даты
func FormatAsDate(t time.Time) string {
	return t.Format("2006-01-02")
}
