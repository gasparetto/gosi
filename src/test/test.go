package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"gosi/processors/i8085"
)

func main() {

	fmt.Println("*********************************")
	buffer := new(bytes.Buffer)
	//loadFile("res/test/CPUTEST.COM", buffer)
	//loadFile("res/test/TEST.COM", buffer)
	//loadFile("res/test/8080PRE.COM", buffer)
	loadFile("res/test/8080EX1.COM", buffer)
	fmt.Printf("File \"%s\" loaded, size %d\n", "TEST.COM", buffer.Len())

	cpu := new(i8085.I8085)

	ram := make([]byte, 0x10000)
	// copy program to ram
	for i, b := range buffer.Bytes() {
		ram[i+0x0100] = b
	}
	ram[0x0005] = 0xC9 // Inject RET at 0x0005 to handle "CALL 5".
	cpu.AttachRam(ram, 0x0000)

	cpu.SetProgramCounter(0x0100)

	for {

		pc := cpu.GetProgramCounter()

		if pc == 0x069f {
			log.Fatal("Stop")
		}

		//cpu.DebugPrintNextOperation()
		//cpu.DebugPrintInternalState()

		if pc == 0x0000 {
			//fmt.Printf("Jump to 0000 from %04X\n", pc)
			os.Exit(1)
		}
		if pc == 0x0005 {
			putchar(cpu, ram)
			//log.Fatal("Stop")
		}

		cpu.Step()
	}
}

func putchar(cpu *i8085.I8085, ram []byte) {

	if cpu.GetRegs().C == 9 {
		var i uint16
		for i = cpu.GetRegDE(); ram[i] != '$'; i += 1 {
			fmt.Printf("%c", ram[i])
		}
	}
	if cpu.GetRegs().C == 2 {
		fmt.Printf("%c", cpu.GetRegs().E)
	}
}

func loadFile(name string, buffer *bytes.Buffer) {

	filerc, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer filerc.Close()

	buffer.ReadFrom(filerc)
}
