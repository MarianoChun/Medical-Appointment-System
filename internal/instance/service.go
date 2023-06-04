package instance

import "fmt"

type DatabaseInstantiatorService struct{}

func NewDatabaseInstantiator() DatabaseInstantiatorService {
	return DatabaseInstantiatorService{}
}

func (s DatabaseInstantiatorService) Execute() {
	fmt.Println("Instantiating Database")
}
