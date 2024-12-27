package repository

import (
	"context"
	"e-wallet-ums/internal/models"
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var (
		user *models.User
		err  error
	)
	err = r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) InsertNewUserSession(ctx context.Context, session *models.UserSession) error {
	return r.DB.Create(session).Error
}

func (r *UserRepository) DeleteUserSession(ctx context.Context, token string) error {
	return r.DB.Exec("DELETE FROM user_sessions WHERE token = ?", token).Error
}

func (r *UserRepository) GetUserSessionByToken(ctx context.Context, token string) (*models.UserSession, error) {
	var (
		session *models.UserSession
		err     error
	)
	err = r.DB.First(&session, "token = ?", token).Error
	if err != nil {
		return session, err
	}

	if session.ID == 0 {
		return session, errors.New("user not found")
	}

	if session == nil || session.ID == 0 {
		return nil, errors.New("user not found")
	}
	return session, nil
}
