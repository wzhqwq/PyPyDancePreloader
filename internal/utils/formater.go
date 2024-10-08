package utils

import (
	"fmt"
	"math"
)

func PrettyByteSize(b int64) string {
	bf := float64(b)
	for _, unit := range []string{"", "K", "M", "G", "T", "P", "E", "Z"} {
		if math.Abs(bf) < 1024.0 {
			return fmt.Sprintf("%3.1f%sB", bf, unit)
		}
		bf /= 1024.0
	}
	return fmt.Sprintf("%.1fYiB", bf)
}

func PrettyTime(s float64) string {
	return fmt.Sprintf("%02d:%02d", int(s)/60, int(s)%60)
}
