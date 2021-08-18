package exceptions

type AppException struct {
	Code    string
	Message string
}

//func (e *AppException) Error() string {
//
//	return fmt.Sprintf("code: %s - message: %s", e.Code, e.Message)
//}

func NewAppException(code string, message string) AppException {
	return AppException{
		Code:    code,
		Message: message,
	}
}
