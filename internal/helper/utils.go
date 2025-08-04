package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseToInt(input string) (int, error) {
	val, err := strconv.Atoi(input)
	if err != nil || val <= 0 {
		return val, err
	}

	return val, nil

}

func ParseToSlug(input string) string {
	// Convert to lowercase
	slug := strings.ToLower(input)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove special characters
	slug = regexp.MustCompile(`[^a-z0-9-]`).ReplaceAllString(slug, "")

	// Trim hyphens from start and end
	slug = strings.Trim(slug, "-")

	return slug
}

func GenerateInvoice() string {
	timestamp := time.Now().UnixMilli()
	return fmt.Sprintf("INVOICE-%d", timestamp)

}
