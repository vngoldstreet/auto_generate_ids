package main

import (
	"vietvd/gennate_id/repository"
	"vietvd/gennate_id/service"
)

func init() {
	repository.ConnectToMongoDB()
	service.SaveMQLID()
}

func main() {

}
