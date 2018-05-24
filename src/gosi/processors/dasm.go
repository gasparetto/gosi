package processors

import (
	"fmt"
)

var debug = true

func dasmOpReg(op string, r string) {
	if debug {
		print(fmt.Sprintf("%s\t%s", op, r))
	}
}

func dasmOpReg1Reg2(op string, r1 string, r2 string) {
	if debug {
		print(fmt.Sprintf("%s\t%s, %s", op, r1, r2))
	}
}

func dasmOpRegVal(op string, r string, v uint8) {
	if debug {
		print(fmt.Sprintf("%s\t%s, 0x%02x", op, r, v))
	}
}

func dasmOpRegVal1Val2(op string, r string, v1 uint8, v2 uint8) {
	if debug {
		print(fmt.Sprintf("%s\t%s, 0x%02x%02x", op, r, v1, v2))
	}
}

func dasmOpRegAddr(op string, r string, addr uint16) {
	if debug {
		print(fmt.Sprintf("%s\t%s, 0x%04x", op, r, addr))
	}
}

func dasmOpAddr(op string, addr uint16) {
	if debug {
		print(fmt.Sprintf("%s\t0x%04x", op, addr))
	}
}

func dasmOpAddrReg(op string, addr uint16, r string) {
	if debug {
		print(fmt.Sprintf("%s\t0x%04x, %s", op, addr, r))
	}
}

func dasmOpAddrVal(op string, addr uint16, v uint8) {
	if debug {
		print(fmt.Sprintf("%s\t0x%04x, 0x%02x", op, addr, v))
	}
}

func dasmOp(op string) {
	if debug {
		print(op)
	}
}

func print(message string) {
	fmt.Printf("0x%04x   %s\n", regs.PC, message)
}

func trace() {
	fmt.Printf("> SP = %04x   mem = %02x %02x [%02x] %02x %02x\n",
		regs.SP, memRead(regs.SP-2), memRead(regs.SP-1), memRead(regs.SP), memRead(regs.SP+1), memRead(regs.SP+2))
	fmt.Printf("> PC = %04x   mem = %02x %02x [%02x] %02x %02x\n",
		regs.PC, memRead(regs.PC-2), memRead(regs.PC-1), memRead(regs.PC), memRead(regs.PC+1), memRead(regs.PC+2))
}
