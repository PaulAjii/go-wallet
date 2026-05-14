package sysmsg

// Environment Variables
const (
	NoEnvFile = "No .env file found"
)

// Database
const (
	CannotConnect        = "Unable to connect to database"
	CannotPing           = "Unable to ping database"
	ConnectionSuccessful = "Database connection successful"
	ConnectionClosed     = "Database connection closed"
)

const (
	ErrBadReq                = "Invalid Request Body"
	ErrInvalidCredentials    = "Email Or Password Is Invalid"
	ErrInternalServerError   = "An Internal Server Error Occurred"
	ErrEmailAlreadyExists    = "Email Already Exists"
	ErrUsernameAlreadyExists = "Username Already Exists"
	ErrInvalidPasswordLength = "Password Must Be At Least 8 Characters"
)

const (
	CreationSuccess = "User created successfully"
	LoginSuccess    = "User Logged In Successfully"
)
