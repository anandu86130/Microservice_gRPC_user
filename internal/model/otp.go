package model

type VerifyOTPs struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}
