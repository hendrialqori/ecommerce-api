package usecase

import (
	"context"
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/exception"
	"internship-mini-project/internal/helper"
	"internship-mini-project/internal/model"
	"internship-mini-project/internal/model/mapper"
	"internship-mini-project/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TrxUseCase interface {
	FindAll(ctx context.Context) ([]*model.TrxResponse, error)
	FindById(ctx context.Context, id int) (*model.TrxResponse, error)
	Create(ctx context.Context, req *model.CreateTrxRequest) error
}

type TrxUseCaseImpl struct {
	TrxRepository        repository.TrxRepository
	ProductRepository    repository.ProductRepository
	ProductLogRepository repository.ProductLogRepository
	TrxDetailRepository  repository.TrxDetailRepository
	Logger               *logrus.Logger
	Validate             *validator.Validate
}

// FindById implements TrxUseCase.
func (t *TrxUseCaseImpl) FindById(ctx context.Context, id int) (*model.TrxResponse, error) {
	toko, err := t.TrxRepository.FindByID(ctx, id)
	if err != nil {
		t.Logger.WithError(err).Error("Failed to find transaction by ID")
		return nil, exception.ErrDataNotFound
	}

	return mapper.ToTrxResponse(toko), nil
}

// FindAll implements TrxUseCase.
func (t *TrxUseCaseImpl) FindAll(ctx context.Context) ([]*model.TrxResponse, error) {
	trxs, err := t.TrxRepository.FindAll(ctx)
	if err != nil {
		t.Logger.WithError(err).Error("Failed to find all transaction")
		return nil, err
	}

	var responses = make([]*model.TrxResponse, 0)
	for _, trx := range *trxs {
		responses = append(responses, mapper.ToTrxResponse(&trx))
	}

	return responses, nil
}

// Create implements TrxUseCase.
func (t *TrxUseCaseImpl) Create(ctx context.Context, req *model.CreateTrxRequest) error {
	if err := t.Validate.Struct(req); err != nil {
		t.Logger.WithError(err).Error("Validation failed for register request")
		return err
	}

	var trxDetails = []domain.TransactionDetail{}
	var hargaTotal int

	for _, p := range req.DetailTrx {

		product, err := t.ProductRepository.FindById(ctx, p.ProductID)
		if err != nil {
			t.Logger.WithError(err).Error("Cannot find product")
			return exception.ErrDataNotFound
		}

		// create product log
		log := &domain.ProductLog{
			NamaProduk:    product.NamaProduk,
			Slug:          product.Slug,
			HargaKonsumen: product.HargaKonsumen,
			HargaReseller: product.HargaReseller,
			Deskripsi:     product.Deskripsi,
			IDProduk:      product.ID,
			IDToko:        product.IDToko,
			IDCategory:    product.IDCategory,
		}

		if err := t.ProductLogRepository.Create(ctx, log); err != nil {
			t.Logger.WithError(err).Error("Failed create product log")
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		total := product.HargaKonsumen * uint(p.Kuantitas)

		trxDetails = append(trxDetails, domain.TransactionDetail{
			IDLogProduk: log.ID,
			IDToko:      product.IDToko,
			Kuantitas:   uint(p.Kuantitas),
			HargaTotal:  total,
		})

		hargaTotal += int(total)
	}

	trxReq := &domain.Transaction{
		IDUser:             req.IDUser,
		AlamatPengirimanID: req.AlamatKirimID,
		KodeInvoice:        helper.GenerateInvoice(),
		MethodBayar:        req.MethodBayar,
		HargaTotal:         hargaTotal,
	}

	err := t.TrxRepository.Create(ctx, trxReq)

	if err != nil {
		t.Logger.WithError(err).Error("Failed create new transations")
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	for _, trxDetail := range trxDetails {
		trxDetailReq := &domain.TransactionDetail{
			IDTrx:       trxReq.ID,
			IDLogProduk: trxDetail.IDLogProduk,
			IDToko:      trxDetail.IDToko,
			Kuantitas:   trxDetail.Kuantitas,
			HargaTotal:  trxDetail.HargaTotal,
		}
		if err := t.TrxDetailRepository.Create(ctx, trxDetailReq); err != nil {
			t.Logger.WithError(err).Error("Failed create new detail transations")
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	return nil
}

func NewTrxUseCase(
	TrxRepository repository.TrxRepository,
	ProductRepository repository.ProductRepository,
	ProductLogRepository repository.ProductLogRepository,
	TrxDetailRepository repository.TrxDetailRepository,
	Logger *logrus.Logger,
	Validate *validator.Validate,
) TrxUseCase {
	return &TrxUseCaseImpl{
		TrxRepository:        TrxRepository,
		ProductRepository:    ProductRepository,
		ProductLogRepository: ProductLogRepository,
		TrxDetailRepository:  TrxDetailRepository,
		Logger:               Logger,
		Validate:             Validate,
	}
}
