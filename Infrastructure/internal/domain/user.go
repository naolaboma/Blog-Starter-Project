package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username       string             `bson:"username" json:"username" validate:"required,min=3,max=50"`
	Email          string             `bson:"email" json:"email" validate:"required,email"`
	Password       string             `bson:"password" json:"-" validate:"required,min=6"` // "-" means don't include in JSON
	Role           string             `bson:"role" json:"role"`
	ProfilePicture *Photo             `bson:"profile_picture,omitempty" json:"profile_picture,omitempty"`
	Bio            string             `bson:"bio,omitempty" json:"bio,omitempty"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

type Photo struct {
	Filename   string    `bson:"filename" json:"filename"`
	FilePath   string    `bson:"file_path" json:"file_path"`
	PublicID   string    `bson:"public_id" json:"public_id"`
	UploadedAt time.Time `bson:"uploaded_at" json:"uploaded_at"`
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id primitive.ObjectID) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	Update(user *User) error
	Delete(id primitive.ObjectID) error
	UpdateProfile(id primitive.ObjectID, updates map[string]interface{}) error
	UpdatePassword(id primitive.ObjectID, password string) error
	UpdateRole(id primitive.ObjectID, role string) error
	UploadProfilePicture(id primitive.ObjectID, photo *Photo) error
}

type UserUseCase interface {
	Register(username, email, password string) (*User, error)
	Login(email, password string) (*LoginResponse, error)
	GetByID(id primitive.ObjectID) (*User, error)
	UpdateProfile(id primitive.ObjectID, bio, profilePic, contactInfo *string) (*User, error)
	UpdateRole(id primitive.ObjectID, role string) error
	ValidatePassword(password string) error
	HashPassword(password string) (string, error)
	CheckPassword(password, hash string) bool
	RefreshToken(refreshToken string) (*LoginResponse, error)
	Logout(userID primitive.ObjectID) error
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UpdateProfileRequest struct {
	Bio          *string `json:"bio,omitempty"`
	ProfilePic   *string `json:"profile_pic,omitempty"`
	ContactInfo  *string `json:"contact_info,omitempty"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
} 