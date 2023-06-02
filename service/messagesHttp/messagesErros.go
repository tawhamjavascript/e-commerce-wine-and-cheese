package messagesHttp

import (
	"reflect"
	"strings"
)

type MessageErro struct {
	Status  int
	Message string
}

var MessageErros map[string]*MessageErro

func InitializerMessage() {
	MessageErros = make(map[string]*MessageErro)
	MessageErros["UnprocessableEntityError"] = &MessageErro{
		Status:  422,
		Message: "Field error",
	}
	MessageErros["InvalidValidationError"] = &MessageErro{
		Status:  400,
		Message: "data not allowed",
	}
	MessageErros["ValidationErrors"] = &MessageErro{
		Status:  400,
		Message: "data not allowed",
	}
	MessageErros["ErrNoDocuments"] = &MessageErro{
		Status:  404,
		Message: "Not Found",
	}
	MessageErros["errorString"] = &MessageErro{
		Status:  401,
		Message: "data incorrect",
	}
	MessageErros["WriteException"] = &MessageErro{
		Status:  400,
		Message: "Email already exist",
	}
	MessageErros["ErrSession"] = &MessageErro{
		Status:  500,
		Message: "Internal Server Error",

	}
	MessageErros["ErrTransaction"] = &MessageErro{
		Status:  500,
		Message: "Internal Server Error",
	}
	MessageErros["ErrClientDisconnected"] = &MessageErro{
		Status:  500,
		Message: "Internal Server Error",
	}
	MessageErros["ErrTransactionInProgress"] = &MessageErro{
		Status:  500,
		Message: "Internal Server Error",
	}
	MessageErros["ErrTransactRetryable"] = &MessageErro{
		Status: 500,
		Message: "Internal Server Error",
	}
	MessageErros["ErrTransactWriteConcern"] = &MessageErro{
		Status:  500,
		Message: "Internal Server Error",
	}
	MessageErros["ErrTransactCallback"] = &MessageErro{
		Status:  500,
		Message: "Internal Server Error",
	}
	MessageErros["CommandError"] = &MessageErro{
		Status:  500,
		Message: "Internal Server Error",
	}
	MessageErros["CursorNotFoundError"] = &MessageErro{
		Status:  500,
		Message: "Internal Server Error",
	}
	MessageErros["ErrNilDocument"] = &MessageErro{
		Status:  404,
		Message: "Not Found",
	}
	MessageErros["WriteError"] = &MessageErro{
		Status:  500,
		Message: "Internal Server Error",
	}



}

func GetError(err error) *MessageErro {
	typeObj := reflect.TypeOf(err).String()
	typeSep := strings.Split(typeObj, ".")
	return MessageErros[typeSep[len(typeSep)-1]]

}
