package sync

import "fmt"

type DatabasesSynchronizerService struct{}

func NewDatabasesSynchronizer() DatabasesSynchronizerService {
	return DatabasesSynchronizerService{}
}

func (s DatabasesSynchronizerService) Execute() {
	fmt.Println("Sync Database")
}
