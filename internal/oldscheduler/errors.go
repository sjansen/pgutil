package oldscheduler

type EarlyTerminationError struct{}

func (e *EarlyTerminationError) Error() string {
	return "terminated early"
}
