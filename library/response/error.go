package response

import "fmt"

type DataNotExistError struct {
	Code uint16
	Msg  string
}

func (err *DataNotExistError) Error() string {
	return fmt.Sprintf(" DataNoExistError code = %d, msg = %s", err.Code, err.Msg)
}

func NewDataNoExistError() *DataNotExistError {
	return &DataNotExistError{
		Code: DataNotExist,
		Msg:  errorInfo[DataNotExist],
	}
}

func IsDataNotExistError(err interface{}) bool {
	switch err.(type) {
	case *DataNotExistError:
		return true
	default:
		return false
	}
}
