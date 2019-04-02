package loader

type InternalError struct {
	Original error
}

func (e *InternalError) Error() string {
	return e.Original.Error()
}
