package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "timestamp",
			log.FieldKeyLevel: "severity",
			log.FieldKeyMsg:   "message",
		},
	})

	log.SetLevel(log.DebugLevel)
	handler, err := NewHandler()
	if err != nil {
		log.WithError(err).Fatal("unable to initialize lambda handler")
	}
	log.Info("Starting handler...")

	lambda.Start(handler.Handle)
}
