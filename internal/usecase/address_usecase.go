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
)

type AddressUsecase interface {
	FindAll(ctx context.Context, userId int) ([]*model.AddressResponse, error)
	FindById(ctx context.Context, id int) (*model.AddressResponse, error)
	Create(ctx context.Context, req *model.CreateAddressRequest) (*model.AddressResponse, error)
	Update(ctx context.Context, req *model.UpdateAddressRequest) (*model.AddressResponse, error)
	Delete(ctx context.Context, id int) error
}

type AddressUseCaseImpl struct {
	AddressRepository repository.AddressRepository
	Logger            *logrus.Logger
	Validate          *validator.Validate
}

// Create implements AddressUsecase.
func (a *AddressUseCaseImpl) Create(ctx context.Context, req *model.CreateAddressRequest) (*model.AddressResponse, error) {
	if err := a.Validate.Struct(req); err != nil {
		return nil, err
	}

	address := &domain.Address{
		IDUser:       req.IDUser,
		JudulAlamat:  req.JudulAlamat,
		NamaPenerima: req.NamaPenerima,
		NoTelp:       req.NoTelp,
		DetailAlamat: req.DetailAlamat,
	}

	if err := a.AddressRepository.Create(ctx, address); err != nil {
		a.Logger.WithError(err).Error("Failed to create address")
		return nil, err
	}

	return mapper.ToAddressResponse(address), nil
}

// Delete implements AddressUsecase.
func (a *AddressUseCaseImpl) Delete(ctx context.Context, id int) error {
	if err := a.AddressRepository.Delete(ctx, id); err != nil {
		a.Logger.WithError(err).Error("Failed to delete address")
		return err
	}
	return nil
}

// FindAll implements AddressUsecase.
func (a *AddressUseCaseImpl) FindAll(ctx context.Context, userId int) ([]*model.AddressResponse, error) {
	addresses, err := a.AddressRepository.FindAll(ctx, userId)
	if err != nil {
		a.Logger.WithError(err).Error("Failed to find all addresses")
		return nil, err
	}

	var addressResponses []*model.AddressResponse
	for _, address := range addresses {
		addressResponses = append(addressResponses, mapper.ToAddressResponse(address))
	}

	return addressResponses, nil
}

// FindById implements AddressUsecase.
func (a *AddressUseCaseImpl) FindById(ctx context.Context, id int) (*model.AddressResponse, error) {
	address, err := a.AddressRepository.FindById(ctx, id)
	if err != nil {
		a.Logger.WithError(err).Error("Failed to find address by ID")
		return nil, exception.ErrDataNotFound
	}

	if address == nil {
		a.Logger.Warn("Address not found")
		return nil, exception.ErrDataNotFound
	}

	return mapper.ToAddressResponse(address), nil
}

// Update implements AddressUsecase.
func (a *AddressUseCaseImpl) Update(ctx context.Context, req *model.UpdateAddressRequest) (*model.AddressResponse, error) {
	if err := a.Validate.Struct(req); err != nil {
		return nil, err
	}

	address := &domain.Address{
		ID:           req.ID,
		IDUser:       req.IDUser,
		JudulAlamat:  req.JudulAlamat,
		NamaPenerima: req.NamaPenerima,
		NoTelp:       req.NoTelp,
		DetailAlamat: req.DetailAlamat,
	}

	if err := a.AddressRepository.Update(ctx, address); err != nil {
		a.Logger.WithError(err).Error("Failed to update address")
		return nil, err
	}

	return mapper.ToAddressResponse(address), nil
}

func NewAddressUseCase(
	addressRepository repository.AddressRepository,
	logger *logrus.Logger,
	validate *validator.Validate,
) AddressUsecase {
	return &AddressUseCaseImpl{
		AddressRepository: addressRepository,
		Logger:            logger,
		Validate:          validate,
	}
}
