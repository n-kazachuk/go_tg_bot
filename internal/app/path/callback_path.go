package path

import (
	"errors"
	"fmt"
	"strings"
)

type CallbackPath struct {
	CallbackName string
	CallbackData string
}

var ErrUnknownCallback = errors.New("unknown callback")

func ParseCallback(callbackData string) (CallbackPath, error) {
	callbackParts := strings.SplitN(callbackData, "__", 2)
	if len(callbackParts) != 2 {
		return CallbackPath{}, ErrUnknownCallback
	}

	return CallbackPath{
		CallbackName: callbackParts[0],
		CallbackData: callbackParts[1],
	}, nil
}

func (p CallbackPath) String() string {
	return fmt.Sprintf("%s__%s", p.CallbackName, p.CallbackData)
}
