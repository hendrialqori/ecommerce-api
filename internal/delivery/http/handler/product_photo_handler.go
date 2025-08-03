package handler

import (
	"internship-mini-project/internal/helper"
	"internship-mini-project/internal/model"
	"internship-mini-project/internal/usecase"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ProductPhotoHandler struct {
	PhotoUseCase usecase.ProductPhotoUseCase
	Logger       *logrus.Logger
	Config       *viper.Viper
}

func (h *ProductPhotoHandler) Create(c *fiber.Ctx) error {
	IDProduk, err := helper.ParseToInt(c.FormValue("id_produk"))
	if err != nil {
		h.Logger.Error("Invalid stock value:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid stock value")
	}

	form, err := c.MultipartForm()
	if err != nil {
		h.Logger.Error("Failed to parse multipart form:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid multipart form")
	}

	files := form.File["files"]

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Failed to open file")
		}
		defer src.Close()

		cloud, _ := cloudinary.NewFromURL(h.Config.GetString("CLOUDINARY_URL"))
		cloudinaryFile, err := cloud.Upload.Upload(c.Context(), src, uploader.UploadParams{
			PublicID: "product/" + file.Filename,
		})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to upload: "+file.Filename)
		}

		photo := &model.CreateProductPhotoRequest{
			IDProduct: uint(IDProduk),
			Url:       cloudinaryFile.SecureURL + "?w=500&h=500&c_fill",
		}

		if _, err := h.PhotoUseCase.Create(c.Context(), photo); err != nil {
			h.Logger.Error("Failed to upload photo product:", err)
			return err
		}
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[any]{
		StatusCode: fiber.StatusOK,
		Message:    "Upload photo product successfully",
	})
}

func (h *ProductPhotoHandler) Delete(c *fiber.Ctx) error {
	id, err := helper.ParseToInt(c.Params("id_foto"))
	if err != nil {
		h.Logger.Error("Invalid foto ID:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid foto ID")
	}

	if err := h.PhotoUseCase.Delete(c.Context(), uint(id)); err != nil {
		h.Logger.Error("Failed to delete foto:", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[any]{
		StatusCode: fiber.StatusOK,
		Message:    "Foto deleted successfully",
	})
}

func NewProductPhotoHandler(PhotoUseCase usecase.ProductPhotoUseCase, Logger *logrus.Logger, Config *viper.Viper) *ProductPhotoHandler {
	return &ProductPhotoHandler{
		PhotoUseCase: PhotoUseCase,
		Logger:       Logger,
		Config:       Config,
	}
}
