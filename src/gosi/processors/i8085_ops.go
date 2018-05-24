package processors

import (
	"encoding/binary"
)

// ********** Symbols and Abbreviations **********
//
// DDD,SSS
//   The bit pattern designating one of the registers A,B,C,D,E,H,L (DDD=destination, SSS=source):
//   DDD or SSS   REGISTER NAME
//   111          A
//   000          B
//   001          C
//   010          D
//   011          E
//   100          H
//   101          L
//
// RP
//   The bit pattern designating one of the register pairs B,D,H,SP:
//   RP   REGISTER PAIR
//   00   B-C
//   01   D-E
//   10   H-L
//   11   SP

// ********** Data Transfer Group **********
// This group of instructions transfers data to and from registers and memory.
// Condition flags are not affected by any instruction in this group.

//region MOV r1, r2 (Move Register)

//   (r1) <- (r2)
//   The content of register r2 is moved to register r1.
//   01DDDSSS
//   Cycles: 1  States: 5  Addressing: register  Flags: none

func opMOV_B_B() int {
	dasmOpReg1Reg2("MOV", "B", "B")
	//regs.B = regs.B
	return 1
}

func opMOV_B_C() int {
	dasmOpReg1Reg2("MOV", "B", "C")
	regs.B = regs.C
	return 1
}

func opMOV_B_D() int {
	dasmOpReg1Reg2("MOV", "B", "D")
	regs.B = regs.D
	return 1
}

func opMOV_B_E() int {
	dasmOpReg1Reg2("MOV", "B", "E")
	regs.B = regs.E
	return 1
}

func opMOV_B_H() int {
	dasmOpReg1Reg2("MOV", "B", "H")
	regs.B = regs.H
	return 1
}

func opMOV_B_L() int {
	dasmOpReg1Reg2("MOV", "B", "L")
	regs.B = regs.L
	return 1
}

func opMOV_C_B() int {
	dasmOpReg1Reg2("MOV", "C", "B")
	regs.C = regs.B
	return 1
}

func opMOV_C_C() int {
	dasmOpReg1Reg2("MOV", "C", "C")
	//regs.C = regs.C
	return 1
}

func opMOV_C_D() int {
	dasmOpReg1Reg2("MOV", "C", "D")
	regs.C = regs.D
	return 1
}

func opMOV_C_E() int {
	dasmOpReg1Reg2("MOV", "C", "E")
	regs.C = regs.E
	return 1
}

func opMOV_C_H() int {
	dasmOpReg1Reg2("MOV", "C", "H")
	regs.C = regs.H
	return 1
}

func opMOV_C_L() int {
	dasmOpReg1Reg2("MOV", "C", "L")
	regs.C = regs.L
	return 1
}

func opMOV_D_B() int {
	dasmOpReg1Reg2("MOV", "D", "B")
	regs.D = regs.B
	return 1
}

func opMOV_D_C() int {
	dasmOpReg1Reg2("MOV", "D", "C")
	regs.D = regs.C
	return 1
}

func opMOV_D_D() int {
	dasmOpReg1Reg2("MOV", "D", "D")
	//regs.D = regs.D
	return 1
}

func opMOV_D_E() int {
	dasmOpReg1Reg2("MOV", "D", "E")
	regs.D = regs.E
	return 1
}

func opMOV_D_H() int {
	dasmOpReg1Reg2("MOV", "D", "H")
	regs.D = regs.H
	return 1
}

func opMOV_D_L() int {
	dasmOpReg1Reg2("MOV", "D", "L")
	regs.D = regs.L
	return 1
}

func opMOV_E_B() int {
	dasmOpReg1Reg2("MOV", "E", "B")
	regs.E = regs.B
	return 1
}

func opMOV_E_C() int {
	dasmOpReg1Reg2("MOV", "E", "C")
	regs.E = regs.C
	return 1
}

func opMOV_E_D() int {
	dasmOpReg1Reg2("MOV", "E", "D")
	regs.E = regs.D
	return 1
}

func opMOV_E_E() int {
	dasmOpReg1Reg2("MOV", "E", "E")
	//regs.E = regs.E
	return 1
}

func opMOV_E_H() int {
	dasmOpReg1Reg2("MOV", "E", "H")
	regs.E = regs.H
	return 1
}

func opMOV_E_L() int {
	dasmOpReg1Reg2("MOV", "E", "L")
	regs.E = regs.L
	return 1
}

func opMOV_H_B() int {
	dasmOpReg1Reg2("MOV", "H", "B")
	regs.H = regs.B
	return 1
}

func opMOV_H_C() int {
	dasmOpReg1Reg2("MOV", "H", "C")
	regs.H = regs.C
	return 1
}

func opMOV_H_D() int {
	dasmOpReg1Reg2("MOV", "H", "D")
	regs.H = regs.D
	return 1
}

func opMOV_H_E() int {
	dasmOpReg1Reg2("MOV", "H", "E")
	regs.H = regs.E
	return 1
}

func opMOV_H_H() int {
	dasmOpReg1Reg2("MOV", "H", "H")
	//regs.H = regs.H
	return 1
}

func opMOV_H_L() int {
	dasmOpReg1Reg2("MOV", "H", "L")
	regs.H = regs.L
	return 1
}

func opMOV_L_B() int {
	dasmOpReg1Reg2("MOV", "L", "B")
	regs.L = regs.B
	return 1
}

func opMOV_L_C() int {
	dasmOpReg1Reg2("MOV", "L", "C")
	regs.L = regs.C
	return 1
}

func opMOV_L_D() int {
	dasmOpReg1Reg2("MOV", "L", "D")
	regs.L = regs.D
	return 1
}

func opMOV_L_E() int {
	dasmOpReg1Reg2("MOV", "L", "E")
	regs.L = regs.E
	return 1
}

func opMOV_L_H() int {
	dasmOpReg1Reg2("MOV", "L", "H")
	regs.L = regs.H
	return 1
}

func opMOV_L_L() int {
	dasmOpReg1Reg2("MOV", "L", "L")
	//regs.L = regs.L
	return 1
}

// endregion

//region MOV r, M (Move from memory)

//   (r) <- ((H) (L))
//   The content of the memory location, whose address is in registers Hand L, is moved to register r.
//   01DDD110
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: none

func opMOV_B_M() int {
	addr := hlRead()
	dasmOpRegAddr("MOV", "B", addr)
	regs.B = memRead(addr)
	return 2
}

func opMOV_C_M() int {
	addr := hlRead()
	dasmOpRegAddr("MOV", "C", addr)
	regs.C = memRead(addr)
	return 2
}

func opMOV_D_M() int {
	addr := hlRead()
	dasmOpRegAddr("MOV", "D", addr)
	regs.D = memRead(addr)
	return 2
}

func opMOV_E_M() int {
	addr := hlRead()
	dasmOpRegAddr("MOV", "E", addr)
	regs.E = memRead(addr)
	return 2
}

func opMOV_H_M() int {
	addr := hlRead()
	dasmOpRegAddr("MOV", "H", addr)
	regs.H = memRead(addr)
	return 2
}

func opMOV_L_M() int {
	addr := hlRead()
	dasmOpRegAddr("MOV", "L", addr)
	regs.L = memRead(addr)
	return 2
}

func opMOV_A_M() int {
	addr := hlRead()
	dasmOpRegAddr("MOV", "A", addr)
	regs.A = memRead(addr)
	return 2
}

//endregion

//region MOV M, r (Move to memory)

//   ((H) (L)) <- (r)
//   The content of register r is moved to the memory location whose address is in registers H and L.
//   01110SSS
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: none

func opMOV_M_B() int {
	addr := hlRead()
	dasmOpAddrReg("MOV", addr, "B")
	memWrite(addr, regs.B)
	return 2
}

func opMOV_M_C() int {
	addr := hlRead()
	dasmOpAddrReg("MOV", addr, "C")
	memWrite(addr, regs.C)
	return 2
}

func opMOV_M_D() int {
	addr := hlRead()
	dasmOpAddrReg("MOV", addr, "D")
	memWrite(addr, regs.D)
	return 2
}

func opMOV_M_E() int {
	addr := hlRead()
	dasmOpAddrReg("MOV", addr, "E")
	memWrite(addr, regs.E)
	return 2
}

func opMOV_M_H() int {
	addr := hlRead()
	dasmOpAddrReg("MOV", addr, "H")
	memWrite(addr, regs.H)
	return 2
}

func opMOV_M_L() int {
	addr := hlRead()
	dasmOpAddrReg("MOV", addr, "L")
	memWrite(addr, regs.L)
	return 2
}

func opMOV_M_A() int {
	addr := hlRead()
	dasmOpAddrReg("MOV", addr, "A")
	memWrite(addr, regs.A)
	return 2
}

//endregion

//region MVI r, data (Move Immediate)

//   (r) <- (byte 2)
//   The content of byte 2 of the instruction is moved to register r.
//   00DDD110
//   Cycles: 2  States: 7  Addressing: immediate  Flags: none

func opMVI_B_data() int {
	r := rom[regs.PC+1]
	dasmOpRegVal("MVI", "B", r)
	regs.B = r
	regs.PC++
	return 2
}

func opMVI_C_data() int {
	r := rom[regs.PC+1]
	dasmOpRegVal("MVI", "C", r)
	regs.C = r
	regs.PC++
	return 2
}

func opMVI_D_data() int {
	r := rom[regs.PC+1]
	dasmOpRegVal("MVI", "D", r)
	regs.D = r
	regs.PC++
	return 2
}

func opMVI_E_data() int {
	r := rom[regs.PC+1]
	dasmOpRegVal("MVI", "E", r)
	regs.E = r
	regs.PC++
	return 2
}

func opMVI_H_data() int {
	r := rom[regs.PC+1]
	dasmOpRegVal("MVI", "H", r)
	regs.H = r
	regs.PC++
	return 2
}

func opMVI_L_data() int {
	r := rom[regs.PC+1]
	dasmOpRegVal("MVI", "L", r)
	regs.L = r
	regs.PC++
	return 2
}

//endregion

//region MVI M, data (Move to memory immediate)

//   ((H) (L)) <- (byte 2)
//   The content of byte 2 of the instruction is moved to the memory location whose address is in registers H and L.
//   00110110
//   Cycles: 3  States: 10  Addressing: immed./reg. indirect  Flags: none

func opMVI_M_data() int {
	r := rom[regs.PC+1]
	addr := hlRead()
	dasmOpAddrVal("MVI", addr, r)
	memWrite(addr, r)
	regs.PC++
	return 3
}

//endregion

//region LXI rp, data 16 (Load register pair immediate)

//   (rh) <- (byte 3)
//   (rl) <- (byte 2)
//   Byte 3 of the instruction is moved into the high-order register (rh) of the register pair rp.
//   Byte 2 of the instruction is moved into the low-order register (rl) of the register pair rp.
//   00RP0001
//   Cycles: 3  States: 10  Addressing: immediate  Flags: none

func opLXI_BC_data16() int {
	rh := rom[regs.PC+2]
	rl := rom[regs.PC+1]
	dasmOpRegVal1Val2("LXI", "BC", rh, rl)
	regs.B = rh
	regs.C = rl
	regs.PC += 2
	return 3
}

func opLXI_DE_data16() int {
	rh := rom[regs.PC+2]
	rl := rom[regs.PC+1]
	dasmOpRegVal1Val2("LXI", "DE", rh, rl)
	regs.D = rh
	regs.E = rl
	regs.PC += 2
	return 3
}

func opLXI_HL_data16() int {
	rh := rom[regs.PC+2]
	rl := rom[regs.PC+1]
	dasmOpRegVal1Val2("LXI", "HL", rh, rl)
	regs.H = rh
	regs.L = rl
	regs.PC += 2
	return 3
}

func opLXI_SP_data16() int {
	rp := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpRegAddr("LXI", "SP", rp)
	regs.SP = rp
	regs.PC += 2
	return 3
}

//endregion

//region LDA addr (Load Accumulator direct)

//   (A) <- ((byte 3)(byte 2))
//   The content of the memory location, whose address is specified in byte 2 and byte 3 of the instruction, is moved to register A.
//   00111010
//   Cycles: 4  States: 13  Addressing: direct  Flags: none

func opLDA() int {
	addr := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpAddr("LDA", addr)
	regs.A = memRead(addr)
	regs.PC += 2
	return 4
}

//endregion

//region STA addr (Store Accumulator direct)

//   ((byte 3)(byte 2)) <- (A)
//   The content of the accumulator is moved to the memory location whose address is specified in byte 2 and byte 3 of the
//   instruction.
//   00110010
//   Cycles: 4  States: 13  Addressing: direct  Flags: none

func opSTA() int {
	addr := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpAddr("STA", addr)
	memWrite(addr, regs.A)
	regs.PC += 2
	return 4
}

//endregion

//region LDAX rp (Load accumulator indirect)

//   (A) <- ((rp))
//   The content of the memory location, whose address is in the register pair rp, is moved to register A.
//   00RP1010
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: none

func opLDAX_BC() int {
	dasmOpReg("LDAX", "BC")
	addr := bcRead()
	regs.A = memRead(addr)
	return 2
}

func opLDAX_DE() int {
	dasmOpReg("LDAX", "DE")
	addr := deRead()
	regs.A = memRead(addr)
	return 2
}

//endregion

// ********** Arithmetic Group **********
// This group of instructions performs arithmetic operations on data in registers and memory.
// Unless indicated otherwise, all instructions in this group affect the flags accord ing to the standard rules.

//region INR r (Increment Register)

//   (r) <- (r) + 1
//   The content of register r is incremented by one.
//   00DDD100
//   Cycles: 1  States: 5  Addressing: register  Flags: Z,S,P,AC

func opINR_B() int {
	dasmOpReg("INR", "B")
	regs.B = add(regs.B, 1)
	return 1
}

func opINR_C() int {
	dasmOpReg("INR", "C")
	regs.C = add(regs.C, 1)
	return 1
}

func opINR_D() int {
	dasmOpReg("INR", "D")
	regs.D = add(regs.D, 1)
	return 1
}

func opINR_E() int {
	dasmOpReg("INR", "E")
	regs.E = add(regs.E, 1)
	return 1
}

func opINR_H() int {
	dasmOpReg("INR", "H")
	regs.H = add(regs.H, 1)
	return 1
}

func opINR_L() int {
	dasmOpReg("INR", "L")
	regs.L = add(regs.L, 1)
	return 1
}

//endregion

//region DCR r (Decrement Register)

//   (r) <- (r) - 1
//   The content of register r is decremented by one.
//   00DDD101
//   Cycles: 1  States: 5  Addressing: register  Flags: Z,S,P,AC

func opDCR_B() int {
	dasmOpReg("DCR", "B")
	regs.B = sub(regs.B, 1)
	return 1
}

func opDCR_C() int {
	dasmOpReg("DCR", "C")
	regs.C = sub(regs.C, 1)
	return 1
}

func opDCR_D() int {
	dasmOpReg("DCR", "D")
	regs.D = sub(regs.D, 1)
	return 1
}

func opDCR_E() int {
	dasmOpReg("DCR", "E")
	regs.E = sub(regs.E, 1)
	return 1
}

func opDCR_H() int {
	dasmOpReg("DCR", "H")
	regs.H = sub(regs.H, 1)
	return 1
}

func opDCR_L() int {
	dasmOpReg("DCR", "L")
	regs.L = sub(regs.L, 1)
	return 1
}

//endregion

//region INX rp (Increment register pair)

//   (rh) (rl) <- (rh) (rl) + 1
//   The content of the register pair rp is incremented by one. Note: No condition flags are affected.
//   00RP0011
//   Cycles: 1  States: 5  Addressing: register  Flags: none

func opINX_BC() int {
	dasmOpReg("INX", "BC")
	addr := bcRead()
	addr++
	bcWrite(addr)
	return 1
}

func opINX_DE() int {
	dasmOpReg("INX", "DE")
	addr := deRead()
	addr++
	deWrite(addr)
	return 1
}

func opINX_HL() int {
	dasmOpReg("INX", "HL")
	addr := hlRead()
	addr++
	hlWrite(addr)
	return 1
}

func opINX_SP() int {
	dasmOpReg("INX", "SP")
	regs.SP++
	return 1
}

//endregion

//region DCX rp (Decrement register pair)

//   (rh) (rl) <- (rh) (rl) - 1
//   The content of the register pair rp is decremented by one. Note: No condition flags are affected.
//   00RP1011
//   Cycles: 1  States: 5  Addressing: register  Flags: none

func opDCX_BC() int {
	dasmOpReg("DCX", "BC")
	addr := bcRead()
	addr--
	bcWrite(addr)
	return 1
}

func opDCX_DE() int {
	dasmOpReg("DCX", "DE")
	addr := deRead()
	addr--
	deWrite(addr)
	return 1
}

func opDCX_HL() int {
	dasmOpReg("DCX", "HL")
	addr := hlRead()
	addr--
	hlWrite(addr)
	return 1
}

func opDCX_SP() int {
	dasmOpReg("DCX", "SP")
	regs.SP--
	return 1
}

//endregion

// ********** Branch Group **********
// This group of instructions alter normal sequential program flow.
// Condition flags are not affected by any instruction in this group.

//region JMP addr (Jump)

//   (PC) <- (byte 3) (byte 2)
//   Control is transferred to the instruction whose address is specified in byte 3 and byte 2 of the current instruction.
//   11000011
//   Cycles: 3  States: 10  Addressing: immediate  Flags: none

func opJMP() int {
	addr := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpAddr("JMP", addr)
	regs.PC = addr
	regs.PC--
	return 3
}

//endregion

//region Jconditional addr (Conditional jump)

//   If (CCC),
//     (PC) <- (byte 3) (byte 2)
//   If the specified condition is true, control is transferred to the instruction whose address is specified in byte 3 and byte 2
//   of the current instruction; otherwise, control continues sequentially.
//   11CCC010
//   Cycles: 3  States: 10  Addressing: immediate  Flags: none

func opJNZ() int {
	addr := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpAddr("JNZ", addr)
	if !flags.Z {
		regs.PC = addr
		regs.PC--
	} else {
		regs.PC += 2
	}
	return 3
}

func opJZ() int {
	addr := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpAddr("JZ", addr)
	if flags.Z {
		regs.PC = addr
		regs.PC--
	} else {
		regs.PC += 2
	}
	return 3
}

func opJNC() int {
	addr := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpAddr("JNC", addr)
	if !flags.CY {
		regs.PC = addr
		regs.PC--
	} else {
		regs.PC += 2
	}
	return 3
}

func opJC() int {
	addr := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpAddr("JC", addr)
	if flags.CY {
		regs.PC = addr
		regs.PC--
	} else {
		regs.PC += 2
	}
	return 3
}

func opJPO() int {
	addr := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpAddr("JPO", addr)
	if !flags.P {
		regs.PC = addr
		regs.PC--
	} else {
		regs.PC += 2
	}
	return 3
}

func opJPE() int {
	addr := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpAddr("JPE", addr)
	if flags.P {
		regs.PC = addr
		regs.PC--
	} else {
		regs.PC += 2
	}
	return 3
}

func opJP() int {
	addr := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpAddr("JP", addr)
	if !flags.S {
		regs.PC = addr
		regs.PC--
	} else {
		regs.PC += 2
	}
	return 3
}

func opJM() int {
	addr := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpAddr("JM", addr)
	if flags.S {
		regs.PC = addr
		regs.PC--
	} else {
		regs.PC += 2
	}
	return 3
}

//endregion

//region CALL addr (Call)

//   ((SP) - 1) <- (PCH) ; ((SP) - 2) <- (PCL) ; (SP) <- (SP) - 2 ; (PC) <- (byte 3) (byte2)
//   The high-order eight bits of the next instruction address are moved to the memory location whose address is one less than the
//   content of register SP. The low-order eight bits of the next instruction address are moved to the memory location whose
//   address is two less than the content of register SP. The content of register SP is decremented by 2. Control is transferred to
//   the instruction whose address is specified in byte 3 and byte 2 of the current instruction.
//   11001101
//   Cycles: 5  States: 17  Addressing: immediate/reg. indirect  Flags: none

func opCALL() int {
	addr := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpAddr("CALL", addr)
	nextAddr := regs.PC+3
	memWrite(regs.SP-1, uint8(nextAddr>>8))
	memWrite(regs.SP-2, uint8(nextAddr))
	regs.SP -= 2
	regs.PC = addr
	regs.PC--
	return 5
}

//endregion

//region RET (Return)

//   (PCL) <- ((SP)) ; (PCH) <- ((SP) + 1) ; (SP) <- (SP) + 2
//   The content of the memory location whose address is specified in register SP is moved to the low-order eight bits of register
//   PC. The content of the memory location whose address is one more than the content of register SP is moved to the high-order
//   eight bits of register PC. The content of register SP is incremented by 2.
//   11001001
//   Cycles: 3  States: 10  Addressing: reg. indirect  Flags: none

func opRET() int {
	dasmOp("RET")
	pcl := memRead(regs.SP)
	pch := memRead(regs.SP+1)
	regs.PC = uint16(pcl) | uint16(pch)<<8
	regs.SP += 2
	regs.PC--
	return 5
}

//endregion

// ********** Stack, I/O, and Machine Control Group **********
// This group of instructions performs I/O, manipulates the Stack, and alters internal control flags.
// Unless otherwise specified, condition flags are not affected by any instructions in this group.

//region NOP (No op)

//   No operation is performed. The registers and flags are unaffected.
//   00000000
//   Cycles: 1  States: 4  Flags: none

func opNOP() int {
	dasmOp("NOP")
	return 1
}

//endregion
