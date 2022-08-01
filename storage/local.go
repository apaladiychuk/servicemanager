package storage

import (
	"sync"

	"servicemanager/model"
)

type local struct {
	mx *sync.RWMutex
	db map[int]*model.Service
}

func NewLocal() Storage {
	return &local{
		mx: new(sync.RWMutex),
		db: make(map[int]*model.Service),
	}
}

func (l *local) Set(pid int, service *model.Service) error {
	l.mx.Lock()
	l.db[pid] = service
	l.mx.Unlock()
	return nil
}

func (l *local) Get(pid int) (*model.Service, error) {
	l.mx.RLock()
	defer l.mx.RUnlock()
	service, ok := l.db[pid]
	if !ok {
		return nil, model.NotFound.New("service not found")
	}
	return service, nil
}

func (l *local) GetAll() ([]*model.Service, error) {
	result := make([]*model.Service, 0, len(l.db))
	l.mx.RLock()
	defer l.mx.RUnlock()
	for _, s := range l.db {
		result = append(result, s)
	}
	return result, nil
}
