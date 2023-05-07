package enum

type ErrorEnum struct {
	Code    string
	Message string
}

func ValidationFailedErrorEnum() ErrorEnum {
	return ErrorEnum{
		Code:    "general_68b329da",
		Message: "Validation failed",
	}
}
