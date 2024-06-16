package add

import (
	"github.com/vegadelalyra/go_microservices_the_exception_handler/task/data"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/task/persistance"
)

type Service struct {
	repository persistance.ITaskDbContext
}

func InitService(repo persistance.ITaskDbContext) *Service {
	return &Service{
		repository: repo,
	}
}

func (service *Service) Add(emp *data.Task) (*data.Task, *data.ErrorDetail) {
	return service.repository.Add(emp)
}