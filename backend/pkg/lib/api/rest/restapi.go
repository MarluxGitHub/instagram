package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type restApi interface {
	Init()
	AddEndpoint(endpoint *RestEndpoint)
	SetConfig(config *viper.Viper)
	Start()
}

// create new Instance of restApi
type restApiImpl struct {
    engine *gin.Engine
	endpoints []*RestEndpoint
	config *viper.Viper
}

func (r *restApiImpl) Init() {
	r.engine = gin.Default()
}

func (r *restApiImpl) Start() {
	for _, endpoint := range r.endpoints {
		(*endpoint).SetRouting(r.engine)
	}

	// Todo: Config
	r.engine.Run()
}

func (r *restApiImpl) AddEndpoint(endpoint *RestEndpoint) {
	r.endpoints = append(r.endpoints, endpoint)
}

func (r *restApiImpl) SetConfig(config *viper.Viper) {
	r.config = config
}

func NewRestApi() restApi {
    return &restApiImpl{
		engine: nil,
		endpoints: make([]*RestEndpoint, 0),
		config: nil,
	}
}