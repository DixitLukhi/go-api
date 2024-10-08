package model

import "time"

var (
	LogLevel        = "log-level"
	LogLevelInfo    = "info"
	LogLevelError   = "error"
	LogLevelDebug   = "debug"
	LogLevelWarning = "warn"
)

var (
	ApiPackage        = "api"
	StorePackage      = "store"
	ControllerPackage = "controller"
	ModelPackage      = "model"
	UtilPackage       = "util"
	MainPackage       = "main"
)
var (
	Controller = "controller"
	Store      = "store"
	Api        = "api"
	Main       = "main"
)

var (
	NewServer  = "new-server"
	NewStore   = "new-store"
	CreateUser = "create-user"
	GetUsers   = "get-users"
	GetUser    = "get-user"
	SignUp     = "sign-up"
	SignIn     = "sign-in"
)

var TokenExpiration = time.Hour * 24

var SecretKey = []byte("dixitlukhi-golang-secret-key")
