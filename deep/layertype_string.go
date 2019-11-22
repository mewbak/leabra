// Code generated by "stringer -type=LayerType"; DO NOT EDIT.

package deep

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Deep_-4]
	_ = x[TRC_-5]
	_ = x[LayerTypeN-6]
}

const _LayerType_name = "Deep_TRC_LayerTypeN"

var _LayerType_index = [...]uint8{0, 5, 9, 19}

func (i LayerType) String() string {
	i -= 4
	if i < 0 || i >= LayerType(len(_LayerType_index)-1) {
		return "LayerType(" + strconv.FormatInt(int64(i+4), 10) + ")"
	}
	return _LayerType_name[_LayerType_index[i]:_LayerType_index[i+1]]
}

func StringToLayerType(s string) (LayerType, error) {
	for i := 0; i < len(_LayerType_index)-1; i++ {
		if s == _LayerType_name[_LayerType_index[i]:_LayerType_index[i+1]] {
			return LayerType(i + 4), nil
		}
	}
	return 0, errors.New("String: " + s + " is not a valid option for type: LayerType")
}