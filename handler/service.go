package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"servicemanager/model"
	"servicemanager/storage"
)

type Service struct {
	db storage.Storage
}

func NewService(db storage.Storage) *Service {
	return &Service{db: db}
}
func (s *Service) Mount(router *gin.Engine) {
	router.GET("/", HandlerErrorFunc(s.GetServices))
	router.POST("/", HandlerErrorFunc(s.StartService))
	router.GET("/:id", HandlerErrorFunc(s.GetServiceStatus))
	router.DELETE("/:id", HandlerErrorFunc(s.TerminateService))
}

func (s *Service) GetServiceStatus(c *gin.Context) error {
	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		return model.NotFound.New("pid must be declared")
	}
	service, err := s.db.Get(id)
	if err != nil {
		return err
	}
	return SuccessResult(c, http.StatusOK, service)
}

func (s *Service) GetServices(c *gin.Context) error {
	result, err := s.db.GetAll()
	if err != nil {
		return err
	}
	return SuccessResult(c, http.StatusOK, result)
}

func (s *Service) StartService(c *gin.Context) error {
	var cmdRequest model.Command
	if err := c.BindJSON(&cmdRequest); err != nil {
		return model.InvalidInput.WrapF(err, "invalid payload")
	}
	serv, err := model.StartCommand(cmdRequest.Command, cmdRequest.Args...)
	if err != nil {
		return err
	}

	if serv.PID == 0 {
		return model.Internal.New("wring PID")

	}
	if err = s.db.Set(serv.PID, serv); err != nil {
		return err
	}
	return SuccessResult(c, http.StatusCreated, serv)
}

func (s *Service) TerminateService(c *gin.Context) error {
	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		return model.NotFound.New("pid must be declared")
	}
	service, err := s.db.Get(id)
	if err != nil {
		return err
	}
	service.Cancel()
	return SuccessResult(c, http.StatusNoContent, service)
}
