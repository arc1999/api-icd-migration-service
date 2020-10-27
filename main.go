package main

import (
	"api-icd-migration-service/service"
)

var s service.ICDService

func main() {
	s.Migrate()
}
