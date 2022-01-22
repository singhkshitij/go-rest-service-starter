package worker

import "github.com/singhkshitij/golang-rest-service-starter/src/worker/stream"

func Setup() {
	stream.CleanUpStreamRules()
	SetupStream()
}

func SetupStream() {
	streamSvc := stream.InitStream()
	stream.StartStreamJob(streamSvc)
	defer SetupStream()
	stream.ConsumeStreamData(streamSvc)
}
