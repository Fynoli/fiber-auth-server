package data

import (
	"fmt"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type User struct {
	Id       int64
	Name     string
	Email    string
	Password string `json:"-"`
}

func CreateDBEngine() (*xorm.Engine, error) {
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "sopagorda", "authserver")
	engine, err := xorm.NewEngine("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	if err := engine.Ping(); err != nil {
		return nil, err
	}

	//Sincroniza las estructuras con las tablas para ver si coinciden antes de iniciarse
	if err := engine.Sync(new(User)); err != nil {
		return nil, err
	}

	return engine, nil
}