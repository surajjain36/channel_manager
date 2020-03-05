package web

import (
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/surajjain36/channel_manager/infra"
	"github.com/surajjain36/channel_manager/misc"
)

// Service HTTP server info
type Service struct {
	shutdownChan chan bool
	trackChannel []misc.ChannelData

	router    *gin.Engine
	wg        sync.WaitGroup
	mdb       *infra.MongoDB
	AppName   string
	Version   string
	BuildTime string
}

// NewService Create a new service
func NewService(conf *misc.Config) (*Service, error) {
	//Get connection with MongoDB.
	//var err error
	// var mDB *infra.MongoDB
	// mDB, err = infra.NewMongo(&conf.Mongo)
	// if err != nil {
	// 	log.WithError(err).Error("Failed to connect to MongoDB")
	// 	return nil, err
	// }

	s := &Service{
		router: gin.New(),
		//mdb:          mDB,
		shutdownChan: make(chan bool),
	}

	s.router.Use(gin.Logger())

	s.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	s.router.GET("/", s.index)
	v1 := s.router.Group("/v1")
	{
		v1.GET("/_create", s.CreateRoutine)
		v1.GET("/_check", s.CheckRoutine)
		v1.GET("/_pause", s.PauseRoutine)
		v1.GET("/_clear", s.StopRoutine)
	}

	return s, nil
}

// Start the web service
func (s *Service) Start(address string) error {
	return s.router.Run(address)
}

// Close all threads and free up resources
func (s *Service) Close() {
	close(s.shutdownChan)

	s.wg.Wait()

	//s.rc.Close()
}
