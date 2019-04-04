// Code generated by "stringer -type=opcode"; DO NOT EDIT.

package websocket

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[opContinuation-0]
	_ = x[opText-1]
	_ = x[opBinary-2]
	_ = x[opClose-11]
	_ = x[opPing-12]
	_ = x[opPong-13]
}

const (
	_opcode_name_0 = "opContinuationopTextopBinary"
	_opcode_name_1 = "opCloseopPingopPong"
)

var (
	_opcode_index_0 = [...]uint8{0, 14, 20, 28}
	_opcode_index_1 = [...]uint8{0, 7, 13, 19}
)

func (i opcode) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _opcode_name_0[_opcode_index_0[i]:_opcode_index_0[i+1]]
	case 11 <= i && i <= 13:
		i -= 11
		return _opcode_name_1[_opcode_index_1[i]:_opcode_index_1[i+1]]
	default:
		return "opcode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}