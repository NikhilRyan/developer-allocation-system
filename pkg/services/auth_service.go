package services

import (
    "developer-allocation-system/pkg/models"
    "developer-allocation-system/pkg/repositories"
    "errors"
    "time"

    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
    Register(user *models.User) error
    Login(credentials models.Credentials) (string, error)
    RefreshToken(refreshToken string) (string, error)
    VerifyToken(tokenString string) (int, error)
}

type authService struct {
    userRepo  repositories.UserRepository
    jwtSecret string
}

func NewAuthService(userRepo repositories.UserRepository, jwtSecret string) AuthService {
    return &authService{
        userRepo:  userRepo,
        jwtSecret: jwtSecret,
    }
}

func (s *authService) Register(user *models.User) error {
    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    user.Role = "user" // Default role

    // Create user in the database
    return s.userRepo.Create(user)
}

func (s *authService) Login(credentials models.Credentials) (string, error) {
    // Retrieve user by username
    user, err := s.userRepo.GetByUsername(credentials.Username)
    if err != nil {
        return "", errors.New("invalid username or password")
    }

    // Compare password
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
    if err != nil {
        return "", errors.New("invalid username or password")
    }

    // Generate JWT token
    tokenString, err := s.generateToken(user)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func (s *authService) RefreshToken(refreshToken string) (string, error) {
    // Implement refresh token logic if applicable
    return "", errors.New("not implemented")
}

func (s *authService) VerifyToken(tokenString string) (int, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(s.jwtSecret), nil
    })
    if err != nil || !token.Valid {
        return 0, errors.New("invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return 0, errors.New("invalid token")
    }

    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        return 0, errors.New("invalid token")
    }

    return int(userIDFloat), nil
}

func (s *authService) generateToken(user *models.User) (string, error) {
    claims := jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.jwtSecret))
}
