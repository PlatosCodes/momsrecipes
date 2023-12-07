package util

// Constants for all supported cuisines
const (
	American = "American"
	Cuban    = "Cuban"
	Italian  = "Italian"
	Mexican  = "Mexican"
)

// IsSupportedCuisineType returns true if the cuisine type is supported
func IsSupportedCuisineType(cuisine_type string) bool {
	switch cuisine_type {
	case American, Cuban, Italian, Mexican:
		return true
	}
	return false
}
