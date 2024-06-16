package persistance

import "github.com/vegadelalyra/go_microservices_the_exception_handler/employee/data"


type IEmployeeDbContext interface{
	GetAll() ([]data.Employee, *data.ErrorDetail)
	GetById(id int) (*data.Employee, *data.ErrorDetail)
	Update(emp *data.Employee) (*data.Employee, *data.ErrorDetail)
	Add(emp *data.Employee) (*data.Employee, *data.ErrorDetail)
	Delete(id int) (*data.Employee, *data.ErrorDetail)
}