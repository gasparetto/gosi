package mb14241

import (
	"fmt"
)

var debug = false
var trace = false

func dasmOp(op string) {
	if debug {
		printLine(fmt.Sprintf("%s", op))
	}
}

func dasmOpVal(op string, v uint8) {
	if debug {
		printLine(fmt.Sprintf("%s\t0x%02x", op, v))
	}
}

func printLine(message string) {
	fmt.Printf("MB14241  %s\n", message)
}

func traceState() {
	if trace {
		fmt.Printf("> (count = %02x   data = %04x)\n", count, data)
	}
}

func traceStateAndResult(result uint8) {
	if trace {
		fmt.Printf("> (count = %02x   data = %04x   result = %02x)\n", count, data, result)
	}
}
