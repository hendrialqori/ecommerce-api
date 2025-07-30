package usecase

import (
	"context"
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/exception"
	"internship-mini-project/internal/model"
	"internship-mini-project/internal/model/mapper"
	"internship-mini-project/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type TokoUsecase interface {
	FindAll(ctx context.Context, query *model.QueryParams) ([]*model.TokoResponse, error)
	FindById(ctx context.Context, id int) (*model.TokoResponse, error)
	Current(ctx context.Context, userId int) (*model.TokoResponse, error)
	Create(ctx context.Context, req *model.CreateTokoRequest) (*model.TokoResponse, error)
	Update(ctx context.Context, req *model.UpdateTokoRequest) (*model.TokoResponse, error)
}

type TokoUseCaseImpl struct {
	TokoRepository repository.TokoRepository
	Logger         *logrus.Logger
	Validate       *validator.Validate
	Config         *viper.Viper
}

// Current implements TokoUsecase.
func (t *TokoUseCaseImpl) Current(ctx context.Context, userId int) (*model.TokoResponse, error) {
	toko, err := t.TokoRepository.Current(ctx, userId)
	if err != nil {
		t.Logger.WithError(err).Error("Failed to find current toko")
		return nil, exception.ErrDataNotFound
	}
	return mapper.ToTokoResponse(toko), nil
}

// Create implements TokoUsecase.
func (t *TokoUseCaseImpl) Create(ctx context.Context, req *model.CreateTokoRequest) (*model.TokoResponse, error) {
	if err := t.Validate.Struct(req); err != nil {
		t.Logger.WithError(err).Error("Validation failed for create request")
		return nil, err
	}

	toko := &domain.Toko{
		IDUser:   req.IDUser,
		NamaToko: req.NamaToko,
	}

	if err := t.TokoRepository.Create(ctx, toko); err != nil {
		t.Logger.WithError(err).Error("Failed to create toko")
		return nil, err
	}

	return mapper.ToTokoResponse(toko), nil
}

// FindAll implements TokoUsecase.
func (t *TokoUseCaseImpl) FindAll(ctx context.Context, query *model.QueryParams) ([]*model.TokoResponse, error) {
	tokos, err := t.TokoRepository.FindAll(ctx, query)
	if err != nil {
		t.Logger.WithError(err).Error("Failed to find all tokos")
		return nil, err
	}

	var responses []*model.TokoResponse
	for _, toko := range tokos {
		responses = append(responses, &model.TokoResponse{
			ID:        toko.ID,
			IDUser:    toko.IDUser,
			NamaToko:  toko.NamaToko,
			UrlFoto:   toko.UrlFoto,
			CreatedAt: toko.CreatedAt,
			UpdatedAt: toko.UpdatedAt,
		})
	}

	return responses, nil
}

// FindById implements TokoUsecase.
func (t *TokoUseCaseImpl) FindById(ctx context.Context, id int) (*model.TokoResponse, error) {
	toko, err := t.TokoRepository.FindById(ctx, id)
	if err != nil {
		t.Logger.WithError(err).Error("Failed to find toko by ID")
		return nil, exception.ErrDataNotFound
	}

	return mapper.ToTokoResponse(toko), nil
}

// Update implements TokoUsecase.
func (t *TokoUseCaseImpl) Update(ctx context.Context, req *model.UpdateTokoRequest) (*model.TokoResponse, error) {
	if err := t.Validate.Struct(req); err != nil {
		t.Logger.WithError(err).Error("Validation failed for update request")
		return nil, err
	}

	toko := &domain.Toko{
		ID:       req.ID,
		IDUser:   req.IDUser,
		NamaToko: req.NamaToko,
		UrlFoto:  req.UrlFoto,
	}

	if err := t.TokoRepository.Update(ctx, toko); err != nil {
		t.Logger.WithError(err).Error("Failed to update toko")
		return nil, err
	}

	return mapper.ToTokoResponse(toko), nil
}

func NewTokoUseCase(
	tokoRepository repository.TokoRepository,
	logger *logrus.Logger,
	validate *validator.Validate,
	config *viper.Viper,
) TokoUsecase {
	return &TokoUseCaseImpl{
		TokoRepository: tokoRepository,
		Logger:         logger,
		Validate:       validate,
		Config:         config,
	}
}
