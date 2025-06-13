package test

import (
	"fmt"
	"programmerzamannow/belajar-golang-restful-api/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, _ := simple.InitializeSimpleService()
	fmt.Println(simpleService.SimpleRepository)
}

func TestSimpleServiceError(t *testing.T) {
	simpleRepository := &simple.SimpleRepository{Error: true}
	simpleService, err := simple.NewSimpleService(simpleRepository)
	if err == nil {
		t.Fatal("Expected error but got nil")
	}
	if simpleService != nil {
		t.Fatalf("Expected simple service to be nil, got %v", simpleService)
	}
	fmt.Println(err)
}
