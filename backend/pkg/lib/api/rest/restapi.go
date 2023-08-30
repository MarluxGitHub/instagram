package rest

import "github.com/gin-gonic/gin"

type restApi interface {
	Init()
	AddEndpoint(endpoint *RestEndpoint)
	Start()
}

// create new Instance of restApi
type restApiImpl struct {
    engine *gin.Engine
	endpoints []*RestEndpoint
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

func NewRestApi() restApi {
    return &restApiImpl{
		engine: nil,
		endpoints: make([]*RestEndpoint, 0),
	}
}