package common

import (
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestResult(t *testing.T) {
	result := Result{
		Code:    Code_OK,
		Message: "Code_OK",
	}

	resultByte, _ := proto.Marshal(&result)
	t.Log(resultByte, len(resultByte))

	resultRef := &Result{}
	_ = proto.Unmarshal(resultByte, resultRef)
	t.Log(resultRef.Code, resultRef.Message)
}
