package get

import (
	"github.com/vegadelalyra/go_microservices_the_exception_handler/employee/data"
	"github.com/vegadelalyra/go_microservices_the_exception_handler/employee/persistance"
)

type Service struct {
	repository persistance.IEmployeeDbContext
}

func InitService(repo persistance.IEmployeeDbContext) *Service {
	return &Service{
		repository: repo,
	}
}

func (service *Service) GetAll() ([]data.Employee, *data.ErrorDetail) {
	return service.repository.GetAll()
}

func (service *Service) GetById(id int) (*data.Employee, *data.ErrorDetail) {
	return service.repository.GetById(id)
}