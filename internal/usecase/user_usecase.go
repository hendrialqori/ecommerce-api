package usecase

import (
	"context"
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/exception"
	"internship-mini-project/internal/model"
	"internship-mini-project/internal/model/mapper"
	"internship-mini-project/internal/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type UserUseCase interface {
	Register(ctx context.Context, req *model.RegisterUserRequest) (*model.UserResponse, error)
	Login(ctx context.Context, req *model.LoginUserRequest) (*model.TokenResponse, error)
}

type UserUseCaseImpl struct {
	UserRepository repository.UserRepository
	Logger         *logrus.Logger
	Validate       *validator.Validate
	Config         *viper.Viper
}

// Login implements UserUseCase.
func (u *UserUseCaseImpl) Login(ctx context.Context, req *model.LoginUserRequest) (*model.TokenResponse, error) {
	if err := u.Validate.Struct(req); err != nil {
		u.Logger.WithError(err).Error("Validation failed for login request")

		return nil, err
	}

	user, err := u.UserRepository.FindByNoTelp(ctx, req.NoTelp)

	if err != nil {
		u.Logger.WithError(err).Error("failed to find user by notelp")
		return nil, exception.ErrUserNotFound
	}

	claims := jwt.MapClaims{
		"id":     user.ID,
		"nama":   user.Nama,
		"notelp": user.NoTelp,
		"exp":    time.Now().Add(2 * time.Hour).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(u.Config.GetString("JWT_SECRET_KEY")))
	if err != nil {
		u.Logger.WithError(err).Error("failed sign token")
		return nil, exception.ErrInternalServerError
	}

	tokenResponse := &model.TokenResponse{
		Token: token,
	}

	return tokenResponse, nil
}

// Register implements UserUseCase.
func (u *UserUseCaseImpl) Register(ctx context.Context, req *model.RegisterUserRequest) (*model.UserResponse, error) {
	if err := u.Validate.Struct(req); err != nil {
		u.Logger.WithError(err).Error("Validation failed for register request")

		return nil, err
	}

	// Check if user already exists

	existingUser, err := u.UserRepository.FindByNoTelp(ctx, req.NoTelp)

	if err != nil {
		u.Logger.WithError(err).Error("notelp already taken")
		return nil, exception.ErrInternalServerError
	}

	if existingUser != nil {
		u.Logger.Warn("user already exists with the provided notelp")
		return nil, exception.ErrUserAlreadyExists
	}

	user := &domain.User{
		Nama:         req.Nama,
		Email:        req.Email,
		KataSandi:    req.KataSandi,
		NoTelp:       req.NoTelp,
		TanggalLahir: req.TanggalLahir,
		JenisKelamin: req.JenisKelamin,
		Tentang:      req.Tentang,
		Pekerjaan:    req.Pekerjaan,
		IdProvinsi:   req.IdProvinsi,
		IdKota:       req.IdKota,
		IsAdmin:      true,
	}

	if err := u.UserRepository.Create(ctx, user); err != nil {
		u.Logger.WithError(err).Error("failed create user to database")
		return nil, exception.ErrInternalServerError
	}

	return mapper.ToUserResponse(user), nil
}

func NewUserUseCase(
	userRepo repository.UserRepository,
	logger *logrus.Logger,
	validate *validator.Validate,
	config *viper.Viper,
) UserUseCase {
	return &UserUseCaseImpl{
		UserRepository: userRepo,
		Logger:         logger,
		Validate:       validate,
		Config:         config,
	}
}
