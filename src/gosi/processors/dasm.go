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

func dasmOpVal(op string, v uint8) {
	if debug {
		print(fmt.Sprintf("%s\t0x%02x", op, v))
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
	fmt.Printf("> PC = %04x   mem = %s\n", regs.PC, traceMem(regs.PC))
	fmt.Printf("> SP = %04x   mem = %s\n", regs.SP, traceMem(regs.SP))
	fmt.Printf("> regs A:%02x B:%02x C:%02x D:%02x E:%02x H:%02x L:%02x\n",
		regs.A, regs.B, regs.C, regs.D, regs.E, regs.H, regs.L)
	fmt.Printf("> flags Z:%t S:%t P:%t AC:%t CY:%t\n", flags.Z, flags.S, flags.P, flags.AC, flags.CY)
}

func traceMem(addr uint16) string {
	return fmt.Sprintf("%02x %02x [%02x] %02x %02x", memReadUnsafe(addr-2), memReadUnsafe(addr-1),
		 memReadUnsafe(addr), memReadUnsafe(addr+1), memReadUnsafe(addr+2))
}
func memReadUnsafe(addr uint16) uint8 {
	switch {
	case inBetween(addr, ramStart, ramEnd):
		return ram[addr-ramStart]
	case inBetween(addr, romStart, romEnd):
		return rom[addr-romStart]
	default:
		return 0
	}
}
