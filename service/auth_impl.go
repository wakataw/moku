package service

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/wakataw/moku/config"
	"github.com/wakataw/moku/entity"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/pkg"
	"github.com/wakataw/moku/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

type authService struct {
	userRepo     repository.UserRepository
	tokenManager pkg.TokenManager
	ldapConfig   config.Ldap
	ldapMapping  config.LdapAttributeMapping
}

func (a *authService) LocalLogin(request model.LoginRequest) (*entity.User, error) {
	user, exist := a.userRepo.FindByUsername(request.Username)

	if user.AccountType == "ldap" {
		return &user, ErrNotLocalUser
	}

	if !exist {
		return &entity.User{}, ErrUserNotExists
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return &entity.User{}, ErrWrongUsernamePassword
	}

	return &user, nil
}

func (a *authService) LdapLogin(request model.LoginRequest) (*entity.User, error) {
	ldapRepo, err := repository.NewLdapRepository(&a.ldapConfig, &a.ldapMapping)

	if err != nil {
		return &entity.User{}, err
	}

	user, err := ldapRepo.Authenticate(request.Username, request.Password)

	if err != nil {
		return &entity.User{}, err
	}

	return user, nil

}

func (a *authService) Login(request model.LoginRequest) (*model.LoginResponse, error) {
	// first try local login
	user, err := a.LocalLogin(request)

	// if user is not a local user or user doesn't exist auth using ldap
	if errors.Is(err, ErrNotLocalUser) || err == ErrUserNotExists {
		newUser := errors.Is(err, ErrUserNotExists)
		err = nil
		ldapUserProfile, err := a.LdapLogin(request)

		if err != nil {
			return &model.LoginResponse{}, err
		}

		// if user doesn't exist, insert new one
		if newUser {
			user = ldapUserProfile
			err := a.userRepo.Insert(user)

			if err != nil {
				return &model.LoginResponse{}, err
			}
		}
	}

	if err != nil {
		log.Println(err.Error())
		return &model.LoginResponse{}, err
	}

	// construct role slice
	var roles []string
	if user.IsAdmin {
		roles = append(roles, "admin")
	}
	if user.IsManager {
		roles = append(roles, "manager")
	}
	if user.IsTeacher {
		roles = append(roles, "teacher")
	}

	token, err := a.tokenManager.GenerateToken(user.ID, roles)

	if err != nil {
		return &model.LoginResponse{}, errors.New("user does not exist")
	}

	response := model.LoginResponse{
		AcessToken:   token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return &response, nil

}

func (a *authService) RefreshToken(request model.RefreshTokenRequest) (*model.LoginResponse, error) {
	var response model.LoginResponse

	// validate token
	token, err := a.tokenManager.Validate(request.RefreshToken)

	if err != nil {
		return &response, err
	}
	claims := token.Claims.(jwt.MapClaims)

	if claims["typ"] != "refresh" {
		return &response, errors.New("provided token is not a valid refresh token")
	}

	// get user roles
	var roles []string

	userId, err := strconv.Atoi(claims["sub"].(string))

	if err != nil {
		return &response, err
	}

	roles = a.userRepo.FindRoles(userId)

	// generate new token
	newToken, err := a.tokenManager.GenerateToken(userId, roles)

	if err != nil {
		return &response, err
	}

	response.AcessToken = newToken.AccessToken
	response.RefreshToken = newToken.RefreshToken

	return &response, nil
}

func NewAuthService(userRepo *repository.UserRepository, tokenManager *pkg.TokenManager, ldapConfig *config.Ldap, ldapMapping *config.LdapAttributeMapping) AuthService {
	return &authService{
		userRepo:     *userRepo,
		tokenManager: *tokenManager,
		ldapConfig:   *ldapConfig,
		ldapMapping:  *ldapMapping,
	}
}
