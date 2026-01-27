package runtimes

import (
	"github.com/cruise-org/cruise/pkg/config"
)

type RuntimeService struct {
	Runtimes map[string]Runtime
}

func NewRuntimeService() {
	runtimeNames := config.Cfg.Global.Runtimes

	runtimes := map[string]Runtime{}
	for _, v := range runtimeNames {
		runtimes[v] = runtimeMap[v]()
	}
}
