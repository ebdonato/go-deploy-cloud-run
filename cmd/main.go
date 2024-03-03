package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ebdonato/go-deploy-cloud-run/pkg/server"
	"github.com/ebdonato/go-deploy-cloud-run/util"
)

func main() {
	r := server.NewWebServer()

	port := util.GetEnvVariable("PORT")

	log.Println("Starting web server on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
