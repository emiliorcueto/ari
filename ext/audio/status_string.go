// Code generated by "stringer --type=Status status.go"; DO NOT EDIT

package audio

import "fmt"

const _Status_name = "InProgressFinishedCanceledTimeoutHangupFailed"

var _Status_index = [...]uint8{0, 10, 18, 26, 33, 39, 45}

func (i Status) String() string {
	if i < 0 || i >= Status(len(_Status_index)-1) {
		return fmt.Sprintf("Status(%d)", i)
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
