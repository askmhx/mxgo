package config

import (
	"testing"
	"fmt"
)

func TestConfig(*testing.T){
	config := NewConfig("/Users/MengHX/Desktop/app.conf")
	fmt.Println(config.String("database.dialect"))
}
