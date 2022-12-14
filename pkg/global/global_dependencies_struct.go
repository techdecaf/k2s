package global

import (
	"github.com/gin-gonic/gin"
	"github.com/reactivex/rxgo/v2"
	"github.com/sirupsen/logrus"
	"github.com/techdecaf/k2s/v2/pkg/config"
	"github.com/techdecaf/k2s/v2/pkg/kube"
)

// Server struct
type Server struct {
	Log    *logrus.Entry
	Gin    *gin.Engine
	Kube   *kube.Service
	Config *config.ConfigService
}

// OnModuleInit method
func (t *Server) OnModuleInit() rxgo.Observable {
	return rxgo.Just(t)()
}

// NewDependencies function description
func NewDependencies(
	Log *logrus.Entry,
	Gin *gin.Engine,
	Kube *kube.Service,
	Config *config.ConfigService,
) *Server {

	return &Server{
		Log:    Log,
		Gin:    Gin,
		Kube:   Kube,
		Config: Config,
	}
}
