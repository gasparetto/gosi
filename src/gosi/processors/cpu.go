package processors

type Cpu interface {
	AttachRam(buf []byte, offset uint16)

	AttachRom(buf []byte, offset uint16)

	AttachPortIn(fp func() uint8, port uint8)

	AttachPortOut(fp func(v uint8), port uint8)

	Step() int

	Interrupt(op uint8) int

	DebugPrint()

	DebugBreakpoint() bool
}
