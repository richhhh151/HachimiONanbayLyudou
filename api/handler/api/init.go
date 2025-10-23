package api

import (
	"github.com/FantasyRL/HachimiONanbayLyudou/pkg/base"
	"github.com/FantasyRL/HachimiONanbayLyudou/pkg/constant"
)

var clientSet *base.ClientSet

func Init() {
	clientSet = base.NewClientSet(base.WithMCPClient([]string{constant.ServiceNameMCPLocal, constant.ServiceNameMCPRemote}), base.WithAiProviderClient())
}
