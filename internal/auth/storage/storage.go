package storage

import (
	"context"
	"crm_system/internal/auth/entity"
	"crm_system/internal/auth/repository"
	"crm_system/pkg/auth/logger"
	"fmt"
	"sync"
	"time"
)

type DataStorage struct {
	interval time.Duration
	service  *repository.AuthRepo
	users    map[string]*[]entity.User
	mu       sync.RWMutex
	logger   *logger.Logger
}

func NewDataStorage(interval time.Duration, service *repository.AuthRepo, logger *logger.Logger) *DataStorage {
	return &DataStorage{
		interval: interval,
		service:  service,
		users:    make(map[string]*[]entity.User),
		mu:       sync.RWMutex{},
		logger:   logger,
	}
}

func (ds *DataStorage) Run() {
	ticker := time.NewTicker(ds.interval)
	defer ticker.Stop()

	ctx := context.Background()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			startTime := time.Now()
			fmt.Println(startTime)

			ds.LoadData()

			elapsedTime := time.Since(startTime)

			timeToNextTick := ds.interval - elapsedTime

			time.Sleep(timeToNextTick)
		}
	}

}

func (ds *DataStorage) LoadData() {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	allUsers, err := ds.service.GetAllUsers()
	if err != nil {
		ds.logger.Error("failed to GetAllUsers err: %v", err)
	}

	method := make(map[string]*[]entity.User)

	method["users"] = allUsers

	ds.users = method
}

func (ds *DataStorage) GetAllUsers() (*[]entity.User, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	users, ok := ds.users["users"]
	if !ok {
		return nil, fmt.Errorf("get all users from storage err")
	}
	return users, nil
}
