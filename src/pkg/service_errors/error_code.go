package service_errors

const (
	// Token
	UnExpectedError = "Expected error"
	ClaimsNotFound  = " Claim not found"
	TokenRequired   = "Token required"
	TokenExpired    = "Token expired"
	TokenInvalid    = "Token invalid"
	//Otp
	OtpExists   = "Otp Exists"
	OtpUsed     = "Otp Used"
	OtpNotValid = "Otp invalid"
	//User
	EmailExists      = "Email exists"
	UsernameExists   = "Username exists"
	PermissionDenied = "Permission Denied"

	// DB
	RecordNotFound = "Record not found"
)
