package utils_test

import (
	"testing"

	"github.com/tdl3/cloudflare-helper/utils"
)

func TestShouldReturnCorrectErrMsg(t *testing.T) {
	var code int64 = 1044
	msg := utils.GetErrMsg(code)
	if msg != "A record with these exact values already exists. Please modify or remove this record" {
		t.Errorf("ShouldReturnCorrectErrMsg(1044) = %v;", msg)
	}
	code = 1111
	msg = utils.GetErrMsg(code)
	if msg != "Unkown Error" {
		t.Errorf("ShouldReturnCorrectErrMsg(1044) = %v;", msg)
	}
}
