package chartofaccounts

import (
	"munshee/internal/generic_event"
	"reflect"
	"testing"
)

func TestValidateAccount(t *testing.T) {

	var tests = []struct {
		name string
		args Account
		want *generic_event.GenericEvent
	}{
		{"Testing Correct Account", Account{
			Code:     "1000",
			Name:     "CorrectAccount",
			ParentId: nil,
			Children: nil,
		}, generic_event.GenericEventSuccess(true)},
		{"Testing Incorrect Account", Account{
			Code:     "ASDV",
			Name:     "123123",
			ParentId: nil,
			Children: nil,
		}, generic_event.GenericEventSuccess(false)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}