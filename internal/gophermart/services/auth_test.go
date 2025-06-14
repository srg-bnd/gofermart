package services_test

import (
	"context"
	"strconv"
	"testing"
	"ya41-56/internal/gophermart/models"
	"ya41-56/internal/gophermart/services"
	"ya41-56/internal/shared/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestBuildJWTString(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewGormRepository[models.User](db)

	authService := services.NewAuthService(repo, "secretKey")
	token, err := authService.BuildJWTString("login")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestLoginSuccess(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewGormRepository[models.User](db)
	ctx := context.Background()

	// Create test user
	passwordHash, err := passwordHash("password")
	require.NoError(t, err)
	testUser := models.User{
		Login:        "testuser",
		PasswordHash: passwordHash,
		Status:       models.UserStatusActive,
	}
	err = repo.Create(ctx, &testUser)
	require.NoError(t, err)
	require.NotZero(t, testUser.ID)

	// Login
	authService := services.NewAuthService(repo, "secretKey")
	token, err := authService.Login(ctx, testUser.Login, "password")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestParseAndValidateTokenSuccess(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewGormRepository[models.User](db)
	ctx := context.Background()

	// Create test user
	testUser := models.User{
		Login:    "testuser",
		Password: "password",
	}
	existUser, err := repo.FindByField(ctx, "login", testUser.Login)
	if err == nil {
		repo.Delete(ctx, strconv.Itoa(int(existUser.ID)))
	}

	// Login
	authService := services.NewAuthService(repo, "secretKey")
	token, err := authService.Register(ctx, &testUser)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	user, err := authService.ParseAndValidateToken(ctx, token)
	assert.NoError(t, err)
	assert.NotEmpty(t, user.Login)

}

func TestRegisterSuccess(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewGormRepository[models.User](db)
	ctx := context.Background()

	newUser := models.User{
		Login:    "testuser",
		Password: "password",
	}
	existUser, err := repo.FindByField(ctx, "login", newUser.Login)
	if err == nil {
		// TODO: existUser.ID as string
		repo.Delete(ctx, strconv.FormatUint(uint64(existUser.ID), 10))
	}

	authService := services.NewAuthService(repo, "secretKey")
	token, err := authService.Register(ctx, &newUser)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

// Helpers

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&models.User{})
	require.NoError(t, err)

	return db
}

func passwordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
