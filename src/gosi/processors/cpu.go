package processors

type Cpu interface {
	AttachRam(buf []byte, offset uint16)

	AttachRom(buf []byte, offset uint16)

	Step() int

	Interrupt(op uint8) int

	DebugPrint()

	DebugBreakpoint() bool
}
