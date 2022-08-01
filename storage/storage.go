package storage

import "servicemanager/model"

type Storage interface {
	Set(pid int, service *model.Service) error
	Get(pid int) (*model.Service, error)
	GetAll() ([]*model.Service, error)
}
