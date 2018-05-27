package mb14241

type MB14241 struct{}

var count uint8
var data uint16

func (MB14241) WriteCount(v uint8) {
	dasmOpVal("WriteCount", v)
	count = v & 0x07
	traceState()
}

func (MB14241) WriteData(v uint8) {
	dasmOpVal("WriteData", v)
	data = (uint16(v)<<7) | (data>>8)
	traceState()
}

func (MB14241) ReadResult() uint8 {
	dasmOp("ReadResult")
	v := uint8(data >> (8 - count))
	traceStateAndResult(v)
	return v
}
