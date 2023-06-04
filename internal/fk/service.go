package fk

import "fmt"

type ForeignKeysService struct{}

func NewForeignKeysService() ForeignKeysService {
	return ForeignKeysService{}
}

func (s ForeignKeysService) Create() {
	fmt.Println("Creating FK's")
}

func (s ForeignKeysService) Delete() {
	fmt.Println("Deleting FK's")
}
