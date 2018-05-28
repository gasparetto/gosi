package i8085

func fetch() uint8 {
	v := memRead(regs.PC)
	regs.PC++
	return v
}

func fetch16() uint16 {
	v := memRead16(regs.PC)
	regs.PC += 2
	return v
}
