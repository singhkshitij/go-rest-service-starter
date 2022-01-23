package worker

import (
	"github.com/singhkshitij/golang-rest-service-starter/src/worker/stream"
)

func Setup() {
	svc := stream.CleanUpStreamRules()
	stream.CreateStreamRules(svc)
	SetupStream()
}

func SetupStream() {
	createdStream := stream.StartStreamJob()
	defer SetupStream()
	stream.ConsumeStreamData(createdStream)
}
