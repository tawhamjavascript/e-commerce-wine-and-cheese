package messagesHttp

import (
	"fmt"
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
		Message: "Email or Password incorrect",
	}
	MessageErros["WriteException"] = &MessageErro{
		Status:  400,
		Message: "Email already exist",
	}

}

func GetError(err error) *MessageErro {
	typeObj := reflect.TypeOf(err).String()
	typeSep := strings.Split(typeObj, ".")
	fmt.Println(typeSep)
	return MessageErros[typeSep[len(typeSep)-1]]

}
