package pk

import "fmt"

type PrimaryKeysService struct{}

func NewPrimaryKeysService() PrimaryKeysService {
	return PrimaryKeysService{}
}

func (s PrimaryKeysService) Create() {
	fmt.Println("Creating PK's")
}

func (s PrimaryKeysService) Delete() {
	fmt.Println("Deleting PK's")
}
