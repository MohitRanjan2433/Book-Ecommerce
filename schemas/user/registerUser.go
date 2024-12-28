package user


type RegisterUserSchema struct{
	Username				string `json:"username" validate:"required"`
	Email					string `json:"email" validate:"required,email"`
	Password				string `json:"password" validate:"required"`
	PhoneNumber       string `json:"phone_number" validate:"required"`
	Role					string `json:"role" validate:"required,oneof=user admin vendor"`
}