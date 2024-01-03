package main

func Beat(t int64) int {
	utc1 := t + 3600
	r := utc1 % 86400
	return int(float64(r) / 86.4)
}
