package common

import (
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestResult(t *testing.T) {
	message := "Code_Ok"
	result := Result{
		Code:    Code_OK,
		Message: &message,
	}

	resultByte, _ := proto.Marshal(&result)
	t.Log(resultByte, len(resultByte))

	resultRef := &Result{}
	_ = proto.Unmarshal(resultByte, resultRef)
	t.Log(resultRef.Code, resultRef.Message)
}
