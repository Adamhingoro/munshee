package generic_event

type EventStatus int

const (
	INITIALIZED EventStatus = iota
	PENDING
	INPROGRESS
	SUCCESSFUL
	ABORTED
	INVALID_DATA // this will occur when data validation is failed
	FAILED
)

type GenericEvent struct {
	Status EventStatus
	Payload interface{}
	Error error
}

func GenericEventSuccess(payload interface{}) *GenericEvent{
	return &GenericEvent{
		Status:  SUCCESSFUL,
		Payload: payload,
		Error:   nil,
	}
}

func GenericEventError(err error) *GenericEvent{
	return &GenericEvent{
		Status:  FAILED,
		Payload: nil,
		Error:   err,
	}
}

func GenericEventInvalidData(err error) *GenericEvent{
	return &GenericEvent{
		Status:  INVALID_DATA,
		Payload: nil,
		Error:   err,
	}
}
// TODO Use this generic event to have communicate between services