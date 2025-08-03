package handler

import (
	"internship-mini-project/internal/helper"
	"internship-mini-project/internal/model"
	"internship-mini-project/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	ProductUseCase usecase.ProductUseCase
	PhotoUseCase   usecase.ProductPhotoUseCase
	Logger         *logrus.Logger
}

func (h *ProductHandler) FindAll(c *fiber.Ctx) error {
	query := &model.ProductQueryParams{
		Page:  1,
		Limit: 5,
	}

	if page, err := strconv.Atoi(c.Query("page")); err == nil && page > 0 {
		query.Page = page
	}

	if limit, err := strconv.Atoi(c.Query("limit")); err == nil && limit > 0 {
		query.Limit = limit
	}

	if nama := c.Query("nama_produk"); nama != "" {
		query.NamaProduk = &nama
	}

	if id_category, err := strconv.Atoi(c.Query("id_category")); err == nil && id_category > 0 {
		query.IDCategory = &id_category
	}

	if id_toko, err := strconv.Atoi(c.Query("id_toko")); err == nil && id_toko > 0 {
		query.IDToko = &id_toko
	}

	if min_harga := c.Query("min_harga"); min_harga != "" {
		query.MinHarga = &min_harga
	}

	if max_harga := c.Query("max_harga"); max_harga != "" {
		query.MaxHarga = &max_harga
	}

	products, err := h.ProductUseCase.FindAll(c.Context(), query)
	if err != nil {
		h.Logger.Error("Failed to fetch products:", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[[]*model.ProductResponse]{
		StatusCode: fiber.StatusOK,
		Message:    "Products fetched successfully",
		Data:       products,
		Metadata:   query,
	})

}

func (h *ProductHandler) FindById(c *fiber.Ctx) error {

	idProduct, err := helper.ParseToInt((c.Params("id_product")))
	if err != nil {
		h.Logger.Error("Invalid product ID:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid product ID")
	}

	product, err := h.ProductUseCase.FindByID(c.Context(), uint(idProduct))
	if err != nil {
		h.Logger.Error("Failed to fetch product by ID:", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[*model.ProductResponse]{
		StatusCode: fiber.StatusOK,
		Message:    "Product fetched successfully",
		Data:       product,
	})
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	auth := c.Locals("auth").(*model.Auth)

	NamaProduk := c.FormValue("nama_produk")
	IDCategory, err := helper.ParseToInt(c.FormValue("id_category"))
	if err != nil {
		h.Logger.Error("Invalid category ID:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid category ID")
	}
	HargeRaseller, err := helper.ParseToInt(c.FormValue("harga_reseller"))
	if err != nil {
		h.Logger.Error("Invalid reseller price:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid reseller price")
	}
	HargaKonsumen, err := helper.ParseToInt(c.FormValue("harga_konsumen"))
	if err != nil {
		h.Logger.Error("Invalid consumer price:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid consumer price")
	}
	Stok, err := helper.ParseToInt(c.FormValue("stok"))
	if err != nil {
		h.Logger.Error("Invalid stock value:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid stock value")
	}
	Deskripsi := c.FormValue("deskripsi")

	var req = model.CreateProductRequest{
		NamaProduk:    NamaProduk,
		IDToko:        auth.Toko.ID,
		IDCategory:    uint(IDCategory),
		HargaReseller: uint(HargeRaseller),
		HargaKonsumen: uint(HargaKonsumen),
		Stok:          uint(Stok),
		Slug:          helper.ParseToSlug(NamaProduk),
		Deskripsi:     Deskripsi,
	}

	if err := c.BodyParser(&req); err != nil {
		h.Logger.Error("Failed to parse request body:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	req.IDToko = auth.Toko.ID
	product, err := h.ProductUseCase.Create(c.Context(), &req)
	if err != nil {
		h.Logger.Error("Failed to create product:", err)
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(&model.WebResponse[*model.ProductResponse]{
		StatusCode: fiber.StatusCreated,
		Message:    "Product created successfully",
		Data:       product,
	})
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	id, err := helper.ParseToInt(c.Params("id_product"))
	if err != nil {
		h.Logger.Error("Invalid product ID:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid product ID")
	}

	auth := c.Locals("auth").(*model.Auth)

	NamaProduk := c.FormValue("nama_produk")
	IDCategory, err := strconv.Atoi(c.FormValue("id_category"))
	if err != nil {
		h.Logger.Error("Invalid category ID:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid category ID")
	}
	HargeRaseller, err := strconv.Atoi(c.FormValue("harga_reseller"))
	if err != nil {
		h.Logger.Error("Invalid reseller price:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid reseller price")
	}
	HargaKonsumen, err := strconv.Atoi(c.FormValue("harga_konsumen"))
	if err != nil {
		h.Logger.Error("Invalid consumer price:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid consumer price")
	}
	Stok, err := strconv.Atoi(c.FormValue("stok"))
	if err != nil {
		h.Logger.Error("Invalid stock value:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid stock value")
	}
	Deskripsi := c.FormValue("deskripsi")

	var req = model.UpdateProductRequest{
		ID:            uint(id),
		NamaProduk:    NamaProduk,
		IDToko:        auth.Toko.ID,
		IDCategory:    uint(IDCategory),
		HargaReseller: uint(HargeRaseller),
		HargaKonsumen: uint(HargaKonsumen),
		Stok:          uint(Stok),
		Slug:          helper.ParseToSlug(NamaProduk),
		Deskripsi:     Deskripsi,
	}

	if err := c.BodyParser(&req); err != nil {
		h.Logger.Error("Failed to parse request body:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	product, err := h.ProductUseCase.Update(c.Context(), &req)
	if err != nil {
		h.Logger.Error("Failed to update product:", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[*model.ProductResponse]{
		StatusCode: fiber.StatusOK,
		Message:    "Product updated successfully",
		Data:       product,
	})
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	id, err := helper.ParseToInt(c.Params("id_product"))
	if err != nil {
		h.Logger.Error("Invalid product ID:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid product ID")
	}

	if err := h.ProductUseCase.Delete(c.Context(), uint(id)); err != nil {
		h.Logger.Error("Failed to delete product:", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[any]{
		StatusCode: fiber.StatusOK,
		Message:    "Product deleted successfully",
	})
}

func NewProductHandler(productUseCase usecase.ProductUseCase, logger *logrus.Logger) *ProductHandler {
	return &ProductHandler{
		ProductUseCase: productUseCase,
		Logger:         logger,
	}
}
