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
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type UserUseCase interface {
	Register(ctx context.Context, req *model.RegisterUserRequest) (*model.UserResponse, error)
	Login(ctx context.Context, req *model.LoginUserRequest) (*model.TokenResponse, error)
	Update(ctx context.Context, req *model.UpdateUserRequest) (*model.UserResponse, error)
	Current(ctx context.Context, email string) (*model.UserResponse, error)
}

type UserUseCaseImpl struct {
	UserRepository repository.UserRepository
	TokoRepository repository.TokoRepository
	Logger         *logrus.Logger
	Validate       *validator.Validate
	Config         *viper.Viper
}

// Update implements UserUseCase.
func (u *UserUseCaseImpl) Update(ctx context.Context, req *model.UpdateUserRequest) (*model.UserResponse, error) {
	if err := u.Validate.Struct(req); err != nil {
		u.Logger.WithError(err).Error("Validation failed for update request")
		return nil, err
	}

	user, err := u.UserRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		u.Logger.WithError(err).Error("failed to find user by email")
		return nil, exception.ErrDataNotFound
	}

	if user == nil {
		u.Logger.Warn("user not found")
		return nil, exception.ErrDataNotFound
	}

	user.Nama = req.Nama
	user.Email = req.Email
	user.NoTelp = req.NoTelp
	user.TanggalLahir = req.TanggalLahir
	user.JenisKelamin = req.JenisKelamin
	user.Tentang = req.Tentang
	user.Pekerjaan = req.Pekerjaan
	user.IdProvinsi = req.IdProvinsi
	user.IdKota = req.IdKota

	if err := u.UserRepository.Update(ctx, user); err != nil {
		u.Logger.WithError(err).Error("failed to update user")
		return nil, exception.ErrInternalServerError
	}

	return mapper.ToUserResponse(user), nil
}

// Current implements UserUseCase.
func (u *UserUseCaseImpl) Current(ctx context.Context, email string) (*model.UserResponse, error) {
	user, err := u.UserRepository.FindByEmail(ctx, email)
	if err != nil {
		u.Logger.WithError(err).Error("user not found")
		return nil, exception.ErrDataNotFound
	}

	if user == nil {
		u.Logger.Warn("user not found")
		return nil, exception.ErrDataNotFound
	}

	return mapper.ToUserResponse(user), nil
}

// Login implements UserUseCase.
func (u *UserUseCaseImpl) Login(ctx context.Context, req *model.LoginUserRequest) (*model.TokenResponse, error) {
	if err := u.Validate.Struct(req); err != nil {
		u.Logger.WithError(err).Error("Validation failed for login request")

		return nil, err
	}

	user, err := u.UserRepository.FindByEmail(ctx, req.Email)

	if err != nil {
		u.Logger.WithError(err).Error("failed to find user")
		return nil, exception.ErrDataNotFound
	}

	if user.KataSandi != req.KataSandi {
		u.Logger.Warn("invalid credentials")
		return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid credentials")
	}

	claims := jwt.MapClaims{
		"id":       user.ID,
		"nama":     user.Nama,
		"email":    user.Email,
		"no_telp":  user.NoTelp,
		"is_admin": user.IsAdmin,
		"toko": map[string]any{
			"id":       user.Toko.ID,
			"nama":     user.Toko.NamaToko,
			"url_foto": user.Toko.UrlFoto,
		},
		"exp": time.Now().Add(2 * time.Hour).Unix(),
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
	existingUser, _ := u.UserRepository.FindByEmail(ctx, req.Email)

	if existingUser != nil {
		u.Logger.Warn("user already exists")
		return nil, exception.ErrDataAlreadyExists
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
		u.Logger.WithError(err).Error("failed create new user")
		return nil, exception.ErrInternalServerError
	}

	// Create toko for the user
	toko := &domain.Toko{
		IDUser:   user.ID,
		NamaToko: "Toko " + user.Nama,
	}

	if err := u.TokoRepository.Create(ctx, toko); err != nil {
		u.Logger.WithError(err).Error("failed create new toko for user")
		return nil, exception.ErrInternalServerError
	}

	return mapper.ToUserResponse(user), nil
}

func NewUserUseCase(
	userRepo repository.UserRepository,
	tokoRepo repository.TokoRepository,
	logger *logrus.Logger,
	validate *validator.Validate,
	config *viper.Viper,
) UserUseCase {
	return &UserUseCaseImpl{
		UserRepository: userRepo,
		TokoRepository: tokoRepo,
		Logger:         logger,
		Validate:       validate,
		Config:         config,
	}
}
