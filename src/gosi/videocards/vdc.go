package videocards

type DisplayController interface {
	Render()

	AttachRam(buf []byte)

	DasmTrace()
}
