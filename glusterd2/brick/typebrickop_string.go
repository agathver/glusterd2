// Code generated by "stringer -type=TypeBrickOp"; DO NOT EDIT.

package brick

import "fmt"

const _TypeBrickOp_name = "OpBrickNullOpBrickTerminateOpBrickXlatorInfoOpBrickXlatorOpOpBrickStatusOpBrickOpOpBrickXlatorDefragOpNodeProfileOpNodeStatusOpVolumeBarrierOpOpBrickBarrierOpNodeBitrotOpBrickAttachOpDumpMetricsOpMaxValue"

var _TypeBrickOp_index = [...]uint8{0, 11, 27, 44, 59, 72, 81, 100, 113, 125, 142, 156, 168, 181, 194, 204}

func (i TypeBrickOp) String() string {
	if i >= TypeBrickOp(len(_TypeBrickOp_index)-1) {
		return fmt.Sprintf("TypeBrickOp(%d)", i)
	}
	return _TypeBrickOp_name[_TypeBrickOp_index[i]:_TypeBrickOp_index[i+1]]
}