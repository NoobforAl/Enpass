package contract

type Validation interface {
	ValidatePassword
	ValidateService
	ValidateUser
}
