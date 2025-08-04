package usecase

import (
	"Blog-API/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionUseCase struct {
	sessionRepo domain.SessionRepository
}

func NewSessionUseCase(sessionRepo domain.SessionRepository) domain.SessionUseCase {
	return &SessionUseCase{
		sessionRepo: sessionRepo,
	}
}

func (s *SessionUseCase) CreateSession(userID primitive.ObjectID, username string, refreshToken string) (*domain.Session, error) {
	session := &domain.Session{
		UserID:       userID,
		Username:     username,
		Token:        refreshToken, // store the refreshToken
		IsActive:     true,
		CreatedAt:    time.Now(),
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour), // exp for 7 days
		LastActivity: time.Now(),
	}
	if err := s.sessionRepo.Create(session); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SessionUseCase) GetSessionByUserID(userID primitive.ObjectID) (*domain.Session, error) {
	return s.sessionRepo.GetByUserID(userID)
}

func (s *SessionUseCase) DeleteSession(userID primitive.ObjectID) error {
	return s.sessionRepo.DeleteByUserID(userID)
}

func (s *SessionUseCase) CleanupExpiredSessions() error {
	return s.sessionRepo.DeleteExpired()
}

func (s *SessionUseCase) UpdateSessionActivity(userID primitive.ObjectID) error {
	session, err := s.sessionRepo.GetByUserID(userID)
	if err != nil {
		return err
	}
	return s.sessionRepo.UpdateLastActivity(session.ID)
}
