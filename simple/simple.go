package simple

import "fmt"

type SimpleRepository struct {
	Error bool
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

func NewSimpleService(repo *SimpleRepository) (*SimpleService, error) {
	if repo.Error {
		return nil, fmt.Errorf("repository error")
	} else {
		return &SimpleService{SimpleRepository: repo}, nil
	}
}
