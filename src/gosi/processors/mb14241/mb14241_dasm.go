package mb14241

import (
	"fmt"
)

var debug = false

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

func trace() {
	if debug {
		fmt.Printf("> (count = %02x   data = %04x)\n", count, data)
	}
}

func traceResult(result uint8) {
	if debug {
		fmt.Printf("> (count = %02x   data = %04x   result = %02x)\n", count, data, result)
	}
}
