package helper

import (
	"context"
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
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

func ParseQueryInt(c *fiber.Ctx, key string, defaultVal int) int {
	if valStr := c.Query(key); valStr != "" {
		if val, err := strconv.Atoi(valStr); err == nil && val > 0 {
			return val
		}
	}
	return defaultVal
}

func UploadFileToCloud(ctx context.Context, file multipart.File, filename, cloudinaryURL, folder string) (string, error) {
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return "", err
	}

	uploadParam := uploader.UploadParams{
		PublicID: folder + "/" + filename,
	}

	result, err := cld.Upload.Upload(ctx, file, uploadParam)
	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}
