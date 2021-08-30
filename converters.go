package ytmusic

import (
	"math"
	"strconv"
	"strings"
)

// durationToInt converts the duration string ("4:20") to seconds (260).
func durationToInt(duration interface{}) int {
	items := strings.Split(duration.(string), ":")
	result := 0
	for i := 0; i < len(items); i++ {
		durationInt, _ := strconv.Atoi(items[i])
		result += durationInt * int(math.Pow(60, float64(len(items)-i-1)))
	}
	return result
}
