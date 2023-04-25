package dbmanager

import (
	"errors"
	"fmt"
	config "github.com/bigrainnn/grpc_db_ops/internal/configer"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbManager struct {
	mu          sync.RWMutex
	connections map[string]*gorm.DB
}

func NewDbManager() *DbManager {
	return &DbManager{
		connections: make(map[string]*gorm.DB),
	}
}

func (dm *DbManager) AddConnection(businessID string, connInfo config.DatabaseConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		connInfo.Username, connInfo.Password, connInfo.IP, connInfo.Port, connInfo.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	dm.mu.Lock()
	defer dm.mu.Unlock()
	dm.connections[businessID] = db
	fmt.Println("businessID:", businessID)
	return nil
}

func (dm *DbManager) GetDB(businessID string) (*gorm.DB, error) {
	dm.mu.RLock()
	defer dm.mu.RUnlock()

	db, ok := dm.connections[businessID]
	if !ok {
		return nil, errors.New("unknown business_id")
	}
	return db, nil
}
