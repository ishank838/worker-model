package worker

type Job interface {
	Process()
}
