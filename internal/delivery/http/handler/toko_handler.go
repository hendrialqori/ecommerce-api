package handler

import (
	"internship-mini-project/internal/model"
	"internship-mini-project/internal/usecase"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type TokoHandler struct {
	TokoUseCase usecase.TokoUsecase
	logger      *logrus.Logger
	Config      *viper.Viper
}

func NewTokoHandler(tokoUseCase usecase.TokoUsecase, logger *logrus.Logger, config *viper.Viper) *TokoHandler {
	return &TokoHandler{
		TokoUseCase: tokoUseCase,
		logger:      logger,
		Config:      config,
	}
}
func (h *TokoHandler) FindAll(c *fiber.Ctx) error {
	query := &model.QueryParams{
		Page:  1,
		Limit: 5,
	}

	if page, err := strconv.Atoi(c.Query("page")); err == nil && page > 0 {
		query.Page = page
	}

	if limit, err := strconv.Atoi(c.Query("limit")); err == nil && limit > 0 {
		query.Limit = limit
	}

	if nama := c.Query("nama"); nama != "" {
		query.Nama = &nama
	}

	tokos, err := h.TokoUseCase.FindAll(c.Context(), query)
	if err != nil {
		h.logger.WithError(err).Error("Failed to find all tokos")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[[]*model.TokoResponse]{
		Data:     tokos,
		Metadata: query,
	})
}

func (h *TokoHandler) FindById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id_toko"))
	if err != nil || id <= 0 {
		h.logger.WithError(err).Error("Invalid toko ID")
		return err
	}

	toko, err := h.TokoUseCase.FindById(c.Context(), id)
	if err != nil {
		h.logger.WithError(err).Error("Failed to find toko by ID")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[*model.TokoResponse]{
		Data: toko,
	})
}

func (h *TokoHandler) Current(c *fiber.Ctx) error {
	auth := c.Locals("auth").(*model.Auth)

	toko, err := h.TokoUseCase.Current(c.Context(), int(auth.ID))
	if err != nil {
		h.logger.WithError(err).Error("Failed to find current toko")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[*model.TokoResponse]{
		Data: toko,
	})
}

func (h *TokoHandler) Update(c *fiber.Ctx) error {
	auth := c.Locals("auth").(*model.Auth)

	idToko, err := strconv.Atoi(c.Params("id_toko"))
	if err != nil || idToko <= 0 {
		h.logger.WithError(err).Error("Invalid toko ID")
		return err
	}

	namaToko := c.FormValue("nama_toko")
	file, err := c.FormFile("foto")
	if err != nil {
		h.logger.WithError(err).Error("Failed to get form file")
		return err
	}

	src, err := file.Open()
	if err != nil {
		h.logger.WithError(err).Error("Failed to open form file")
		return err
	}
	defer src.Close()

	cloud, err := cloudinary.NewFromURL(h.Config.GetString("CLOUDINARY_URL"))
	if err != nil {
		h.logger.WithError(err).Error("Failed to create Cloudinary client")
		return err
	}

	cloudinaryFile, err := cloud.Upload.Upload(c.Context(), src, uploader.UploadParams{
		PublicID: "toko/" + file.Filename,
	})

	if err != nil {
		h.logger.WithError(err).Error("Failed to upload file to Cloudinary")
		return err
	}

	req := &model.UpdateTokoRequest{}

	if err := c.BodyParser(req); err != nil {
		h.logger.WithError(err).Error("Failed to parse request body")
		return err
	}

	req.ID = uint(idToko)
	req.IDUser = uint(auth.ID)
	req.NamaToko = namaToko
	req.UrlFoto = cloudinaryFile.SecureURL + "?w=500&h=500&c_fill"

	toko, err := h.TokoUseCase.Update(c.Context(), req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to update toko")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[*model.TokoResponse]{
		Data: toko,
	})
}
