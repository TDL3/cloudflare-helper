package utils_test

import (
	"testing"

	"github.com/tdl3/cloudflare-helper/config"
	"github.com/tdl3/cloudflare-helper/utils"
)

func TestAssembleUrl(t *testing.T) {
	cfg := config.LoadConfig("../config.yml")
	got := utils.AssembleUrl(cfg, utils.UpdateDns)
	if got == "" {
		t.Errorf("AssembleUrl(UpdateDns) = %v;", got)
	}
}
