package processors

import "log"

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
//
// Flags
//   D7 D6 D5 D4 D3 D2 D1 D0
//   S  Z  0  AC 0  P  1  CY

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

func opMOV_B_A() int {
	dasmOpReg1Reg2("MOV", "B", "A")
	regs.B = regs.A
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

func opMOV_C_A() int {
	dasmOpReg1Reg2("MOV", "C", "A")
	regs.C = regs.A
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

func opMOV_D_A() int {
	dasmOpReg1Reg2("MOV", "D", "A")
	regs.D = regs.A
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

func opMOV_E_A() int {
	dasmOpReg1Reg2("MOV", "E", "A")
	regs.E = regs.A
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

func opMOV_H_A() int {
	dasmOpReg1Reg2("MOV", "H", "A")
	regs.H = regs.A
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

func opMOV_L_A() int {
	dasmOpReg1Reg2("MOV", "L", "A")
	regs.L = regs.A
	return 1
}

func opMOV_A_B() int {
	dasmOpReg1Reg2("MOV", "A", "B")
	regs.A = regs.B
	return 1
}

func opMOV_A_C() int {
	dasmOpReg1Reg2("MOV", "A", "C")
	regs.A = regs.C
	return 1
}

func opMOV_A_D() int {
	dasmOpReg1Reg2("MOV", "A", "D")
	regs.A = regs.D
	return 1
}

func opMOV_A_E() int {
	dasmOpReg1Reg2("MOV", "A", "E")
	regs.A = regs.E
	return 1
}

func opMOV_A_H() int {
	dasmOpReg1Reg2("MOV", "A", "H")
	regs.A = regs.H
	return 1
}

func opMOV_A_L() int {
	dasmOpReg1Reg2("MOV", "A", "L")
	regs.A = regs.L
	return 1
}

func opMOV_A_A() int {
	dasmOpReg1Reg2("MOV", "A", "A")
	//regs.A = regs.A
	return 1
}

// endregion

//region MOV r, M (Move from memory)

//   (r) <- ((H) (L))
//   The content of the memory location, whose address is in registers H and L, is moved to register r.
//   01DDD110
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: none

func opMOV_B_M() int {
	addr := getHL()
	dasmOpRegAddr("MOV", "B", addr)
	regs.B = memRead(addr)
	return 2
}

func opMOV_C_M() int {
	addr := getHL()
	dasmOpRegAddr("MOV", "C", addr)
	regs.C = memRead(addr)
	return 2
}

func opMOV_D_M() int {
	addr := getHL()
	dasmOpRegAddr("MOV", "D", addr)
	regs.D = memRead(addr)
	return 2
}

func opMOV_E_M() int {
	addr := getHL()
	dasmOpRegAddr("MOV", "E", addr)
	regs.E = memRead(addr)
	return 2
}

func opMOV_H_M() int {
	addr := getHL()
	dasmOpRegAddr("MOV", "H", addr)
	regs.H = memRead(addr)
	return 2
}

func opMOV_L_M() int {
	addr := getHL()
	dasmOpRegAddr("MOV", "L", addr)
	regs.L = memRead(addr)
	return 2
}

func opMOV_A_M() int {
	addr := getHL()
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
	addr := getHL()
	dasmOpAddrReg("MOV", addr, "B")
	memWrite(addr, regs.B)
	return 2
}

func opMOV_M_C() int {
	addr := getHL()
	dasmOpAddrReg("MOV", addr, "C")
	memWrite(addr, regs.C)
	return 2
}

func opMOV_M_D() int {
	addr := getHL()
	dasmOpAddrReg("MOV", addr, "D")
	memWrite(addr, regs.D)
	return 2
}

func opMOV_M_E() int {
	addr := getHL()
	dasmOpAddrReg("MOV", addr, "E")
	memWrite(addr, regs.E)
	return 2
}

func opMOV_M_H() int {
	addr := getHL()
	dasmOpAddrReg("MOV", addr, "H")
	memWrite(addr, regs.H)
	return 2
}

func opMOV_M_L() int {
	addr := getHL()
	dasmOpAddrReg("MOV", addr, "L")
	memWrite(addr, regs.L)
	return 2
}

func opMOV_M_A() int {
	addr := getHL()
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

func opMVI_B() int {
	r := memRead(regs.PC+1)
	dasmOpRegVal("MVI", "B", r)
	regs.B = r
	regs.PC++
	return 2
}

func opMVI_C() int {
	r := memRead(regs.PC+1)
	dasmOpRegVal("MVI", "C", r)
	regs.C = r
	regs.PC++
	return 2
}

func opMVI_D() int {
	r := memRead(regs.PC+1)
	dasmOpRegVal("MVI", "D", r)
	regs.D = r
	regs.PC++
	return 2
}

func opMVI_E() int {
	r := memRead(regs.PC+1)
	dasmOpRegVal("MVI", "E", r)
	regs.E = r
	regs.PC++
	return 2
}

func opMVI_H() int {
	r := memRead(regs.PC+1)
	dasmOpRegVal("MVI", "H", r)
	regs.H = r
	regs.PC++
	return 2
}

func opMVI_L() int {
	r := memRead(regs.PC+1)
	dasmOpRegVal("MVI", "L", r)
	regs.L = r
	regs.PC++
	return 2
}

func opMVI_A() int {
	r := memRead(regs.PC+1)
	dasmOpRegVal("MVI", "A", r)
	regs.A = r
	regs.PC++
	return 2
}

//endregion

//region MVI M, data (Move to memory immediate)

//   ((H) (L)) <- (byte 2)
//   The content of byte 2 of the instruction is moved to the memory location whose address is in registers H and L.
//   00110110
//   Cycles: 3  States: 10  Addressing: immed./reg. indirect  Flags: none

func opMVI_M() int {
	r := memRead(regs.PC+1)
	addr := getHL()
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

func opLXI_BC() int {
	rh := memRead(regs.PC+2)
	rl := memRead(regs.PC+1)
	dasmOpRegVal1Val2("LXI", "BC", rh, rl)
	regs.B = rh
	regs.C = rl
	regs.PC += 2
	return 3
}

func opLXI_DE() int {
	rh := memRead(regs.PC+2)
	rl := memRead(regs.PC+1)
	dasmOpRegVal1Val2("LXI", "DE", rh, rl)
	regs.D = rh
	regs.E = rl
	regs.PC += 2
	return 3
}

func opLXI_HL() int {
	rh := memRead(regs.PC+2)
	rl := memRead(regs.PC+1)
	dasmOpRegVal1Val2("LXI", "HL", rh, rl)
	regs.H = rh
	regs.L = rl
	regs.PC += 2
	return 3
}

func opLXI_SP() int {
	rp := memRead16(regs.PC+1)
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
	addr := memRead16(regs.PC+1)
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
	addr := memRead16(regs.PC+1)
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
	addr := getBC()
	regs.A = memRead(addr)
	return 2
}

func opLDAX_DE() int {
	dasmOpReg("LDAX", "DE")
	addr := getDE()
	regs.A = memRead(addr)
	return 2
}

//endregion

//region STAX rp (Store accumulator indirect)

//   ((rp)) <- (A)
//   The content of register A is moved to the memory location whose address is in the register pair rp.
//   00RP0010
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: none

func opSTAX_BC() int {
	dasmOpReg("STAX", "BC")
	addr := getBC()
	memWrite(addr, regs.A)
	return 2
}

func opSTAX_DE() int {
	dasmOpReg("STAX", "DE")
	addr := getDE()
	memWrite(addr, regs.A)
	return 2
}

//endregion

//region XCHG (Exchange H and L with D and E)

//   (H) <-> (D) ; (L) <-> (E)
//   The contents of registers H and L are exchanged with the contents of registers D and E.
//   11101011
//   Cycles: 1  States: 4  Addressing: register  Flags: none

func opXCHG() int {
	dasmOp("XCHG")
	r1 := regs.H
	regs.H = regs.D
	regs.D = r1
	r2 := regs.L
	regs.L = regs.E
	regs.E = r2
	return 1
}

//endregion

// ********** Arithmetic Group **********
// This group of instructions performs arithmetic operations on data in registers and memory.
// Unless indicated otherwise, all instructions in this group affect the flags according to the standard rules.

//region ADI data (Add immediate)

//   (A) <- (A) + (byte 2)
//   The content of the second byte of the instruction is added to the content of the accumulator. The result is placed in the
//   accumulator.
//   11000110
//   Cycles: 2  States: 7  Addressing: immediate  Flags: Z,S,P,CY,AC

func opADI() int {
	r := memRead(regs.PC+1)
	dasmOpVal("ADI", r)
	regs.A = add_CY(regs.A, r)
	regs.PC++
	return 2
}

//endregion

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

func opINR_A() int {
	dasmOpReg("INR", "A")
	regs.A = add(regs.A, 1)
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

func opDCR_A() int {
	dasmOpReg("DCR", "A")
	regs.A = sub(regs.A, 1)
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
	addr := getBC()
	addr++
	setBC(addr)
	return 1
}

func opINX_DE() int {
	dasmOpReg("INX", "DE")
	addr := getDE()
	addr++
	setDE(addr)
	return 1
}

func opINX_HL() int {
	dasmOpReg("INX", "HL")
	addr := getHL()
	addr++
	setHL(addr)
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
	addr := getBC()
	addr--
	setBC(addr)
	return 1
}

func opDCX_DE() int {
	dasmOpReg("DCX", "DE")
	addr := getDE()
	addr--
	setDE(addr)
	return 1
}

func opDCX_HL() int {
	dasmOpReg("DCX", "HL")
	addr := getHL()
	addr--
	setHL(addr)
	return 1
}

func opDCX_SP() int {
	dasmOpReg("DCX", "SP")
	regs.SP--
	return 1
}

//endregion

//region DAD rp (Add register pair to H and L)

//   (H) (L) <- (H) (L) + (rh) (rl)
//   The content of the register pair rp is added to the content of the register pair H and L. The result is placed in the register
//   pair H and L. Note: Only the CY flag is affected. It is set if there is a carry out of the double precision add; otherwise it
//   is reset.
//   00RP1001
//   Cycles: 3  States: 10  Addressing: register  Flags: CY

func opDAD_BC() int {
	dasmOpReg("DAD", "BC")
	addr := getHL()
	addr += getBC()
	// todo flag CY
	setHL(addr)
	return 3
}

func opDAD_DE() int {
	dasmOpReg("DAD", "DE")
	addr := getHL()
	addr += getDE()
	// todo flag CY
	setHL(addr)
	return 3
}

func opDAD_HL() int {
	dasmOpReg("DAD", "HL")
	addr := getHL()
	addr += addr
	// todo flag CY
	setHL(addr)
	return 3
}

func opDAD_SP() int {
	dasmOpReg("DAD", "SP")
	addr := getHL()
	addr += regs.SP
	// todo flag CY
	setHL(addr)
	return 3
}

//endregion

// ********** Logical Group **********
// This group of instructions performs logical (Boolean) operations on data in registers and memory and on condition flags.
// Unless indicated otherwise, all instructions in this group affect the flags according to the standard rules.

//region ANA r (AND Register)

//   (A) <- (A) ^ (r)
//   The content of the second byte of the instruction is logically anded with the contents of the accumulator.
//   The result is placed in the accumulator. The CY and AC flags are cleared.
//   11100SSS
//   Cycles: 1  States: 4  Addressing: register  Flags: Z,S,P,CY,AC

func opANA_B() int {
	dasmOpReg("ANA", "B")
	regs.A = regs.A & regs.B
	sub(regs.A, regs.B) // flags Z,S,P & AC
	flags.CY = false
	return 1
}

func opANA_C() int {
	dasmOpReg("ANA", "C")
	regs.A = regs.A & regs.C
	sub(regs.A, regs.C) // flags Z,S,P & AC
	flags.CY = false
	return 1
}

func opANA_D() int {
	dasmOpReg("ANA", "D")
	regs.A = regs.A & regs.D
	sub(regs.A, regs.D) // flags Z,S,P & AC
	flags.CY = false
	return 1
}

func opANA_E() int {
	dasmOpReg("ANA", "E")
	regs.A = regs.A & regs.E
	sub(regs.A, regs.E) // flags Z,S,P & AC
	flags.CY = false
	return 1
}

func opANA_H() int {
	dasmOpReg("ANA", "H")
	regs.A = regs.A & regs.H
	sub(regs.A, regs.H) // flags Z,S,P & AC
	flags.CY = false
	return 1
}

func opANA_L() int {
	dasmOpReg("ANA", "L")
	regs.A = regs.A & regs.L
	sub(regs.A, regs.L) // flags Z,S,P & AC
	flags.CY = false
	return 1
}

func opANA_A() int {
	dasmOpReg("ANA", "A")
	regs.A = regs.A & regs.A
	sub(regs.A, regs.A) // flags Z,S,P & AC
	flags.CY = false
	return 1
}

//endregion

//region ANA M (AND memory)

//   (A) <- (A) ^ ((H) (L))
//   The contents of the memory location whose address is contained in the H and L registers is logically anded with the content of
//   the accumulator. The result is placed in the accumulator. The CY flag is cleared.
//   11100110
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: Z,S,P,CY,AC

func opANA_M() int {
	addr := getHL()
	dasmOpAddr("ANA", addr)
	v := memRead(addr)
	regs.A = regs.A & v
	sub(regs.A, v) // flags Z,S,P & AC
	flags.CY = false
	return 2
}

//endregion

//region ANI data (AND immediate)

//   (A) <- (A) ^ (byte 2)
//   The content of the second byte of the instruction is logically anded with the contents of the accumulator.
//   The result is placed in the accumulator. The CY and AC flags are cleared.
//   11100110
//   Cycles: 2  States: 7  Addressing: immediate  Flags: Z,S,P,CY,AC

func opANI() int {
	r := memRead(regs.PC+1)
	dasmOpVal("ANI", r)
	regs.A = regs.A & r
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	regs.PC++
	return 2
}

//endregion

//region XRA r (Exclusive OR register)

//   (A) <- (A) Y (r)
//   The content of register r is exclusive-or'd with the content of the accumulator. The result is placed in the accumulator.
//   The CY and AC flags are cleared.
//   10101SSS
//   Cycles: 1  States: 4  Addressing: register  Flags: Z,S,P,CY,AC

func opXRA_B() int {
	dasmOpReg("XRA", "B")
	regs.A = regs.A ^ regs.B
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opXRA_C() int {
	dasmOpReg("XRA", "C")
	regs.A = regs.A ^ regs.C
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opXRA_D() int {
	dasmOpReg("XRA", "D")
	regs.A = regs.A ^ regs.D
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opXRA_E() int {
	dasmOpReg("XRA", "E")
	regs.A = regs.A ^ regs.E
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opXRA_H() int {
	dasmOpReg("XRA", "H")
	regs.A = regs.A ^ regs.H
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opXRA_L() int {
	dasmOpReg("XRA", "L")
	regs.A = regs.A ^ regs.L
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opXRA_A() int {
	dasmOpReg("XRA", "A")
	regs.A = regs.A ^ regs.A
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

//endregion

//region XRA M (Exclusive OR Memory)

//   (A) <- (A) Y ((H) (L))
//   The content of the memory location whose address is contained in the H and L registers is exclusive-OR'd with the content of
//   the accumulator. The result is placed in the accumulator. The CY and AC flags are cleared.
//   10101110
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: Z,S,P,CY,AC

func opXRA_M() int {
	addr := getHL()
	dasmOpAddr("XRA", addr)
	regs.A = regs.A ^ memRead(addr)
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

// endregion

//region CPI data (Compare immediate)

//   (A) - (byte 2)
//   The content of the second byte of the instruction is subtracted from the accumulator. The condition flags are set by the
//   result of the subtraction. The Z flag is set to 1 if (A) = (byte 2). The CY flag is set to 1 if (A) < (byte 2).
//   11111110
//   Cycles: 2  States: 7  Addressing: immediate  Flags: Z,S,P,CY,AC

func opCPI() int {
	r := memRead(regs.PC+1)
	dasmOpVal("CPI", r)
	sub_CY(regs.A, r)
	regs.PC++
	return 2
}

//endregion

//region RRC (Rotate right)

//   (An) <- (An-1) ; (A7) <- (A0) ; (CY) <- (A0)
//   The content of the accumulator is rotated right one position. The high order bit and the CY flag are both set to the value
//   shifted out of the low order bit position. Only the CY flag is affected.
//   00001111
//   Cycles: 1  States: 4  Flags: CY

func opRRC() int {
	dasmOp("RRC")
	a0 := regs.A & 0x01
	regs.A = (regs.A>>1 & 0x7f) | a0<<7
	flags.CY = a0 != 0
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
	addr := memRead16(regs.PC+1)
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
	addr := memRead16(regs.PC+1)
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
	addr := memRead16(regs.PC+1)
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
	addr := memRead16(regs.PC+1)
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
	addr := memRead16(regs.PC+1)
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
	addr := memRead16(regs.PC+1)
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
	addr := memRead16(regs.PC+1)
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
	addr := memRead16(regs.PC+1)
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
	addr := memRead16(regs.PC+1)
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
	addr := memRead16(regs.PC+1)
	dasmOpAddr("CALL", addr)
	nextOp := regs.PC+3
	memWrite16(regs.SP-2, nextOp)
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
	//trace()
	dasmOp("RET")
	regs.PC = memRead16(regs.SP)
	regs.SP += 2
	regs.PC--
	//breakpoint = true
	return 5
}

//endregion

// ********** Stack, I/O, and Machine Control Group **********
// This group of instructions performs I/O, manipulates the Stack, and alters internal control flags.
// Unless otherwise specified, condition flags are not affected by any instructions in this group.

//region PUSH rp (Push)

//   ((SP) - 1) <- (rh) ; ((SP) - 2) <- (rl) ; (SP) <- (SP) - 2
//   The content of the high-order register of register pair rp is moved to the memory location whose address is one less than the
//   content of register SP. The content of the low-order register of register pair rp is moved to the memory location whose
//   address is two less than the content of register SP. The of register SP is decremented by 2.
//   11RP0101
//   Cycles: 3  States: 11  Addressing: reg. indirect  Flags: none

func opPUSH_BC() int {
	dasmOpReg("PUSH", "BC")
	memWrite(regs.SP-1, regs.B)
	memWrite(regs.SP-2, regs.C)
	regs.SP -= 2
	return 3
}

func opPUSH_DE() int {
	dasmOpReg("PUSH", "DE")
	memWrite(regs.SP-1, regs.D)
	memWrite(regs.SP-2, regs.E)
	regs.SP -= 2
	return 3
}

func opPUSH_HL() int {
	dasmOpReg("PUSH", "HL")
	memWrite(regs.SP-1, regs.H)
	memWrite(regs.SP-2, regs.L)
	regs.SP -= 2
	return 3
}

//endregion

//region PUSH PSW (Push processor status word)

//   ((SP) - 1) <- (A) ; ((SP) - 2) <- (flags) ; (SP) <- (SP) - 2
//   The content of register A is moved to the memory location whose address is one less than register SP. The contents of the
//   condition flags are assembled into a processor status word and the word is moved to the memory location whose address is two
//   less than the content of register SP. The content of register SP is decremented by two.
//   11110101
//   Cycles: 3  States: 11  Addressing: reg. indirect  Flags: none

func opPUSH_PSW() int {
	dasmOpReg("PUSH", "PSW")
	memWrite(regs.SP-1, regs.A)
	memWrite(regs.SP-2, getFlags())
	regs.SP -= 2
	return 3
}

//endregion

//region POP rp (Pop)

//   (rl) <- ((SP)) ; (rh) <- ((SP) + 1) ; (SP) <- (SP) + 2
//   The content of the memory location, whose address is specified by the content of register SP, is moved to the low-order
//   register of register pair rp. The content of the memory location, whose address is one more than the content of register SP,
//   is moved to the high-order register of register pair rp. The content of register SP is incremented by 2.
//   11RP0001
//   Cycles: 3  States: 11  Addressing: reg. indirect  Flags: none

func opPOP_BC() int {
	dasmOpReg("POP", "BC")
	regs.C = memRead(regs.SP)
	regs.B = memRead(regs.SP+1)
	regs.SP += 2
	return 3
}

func opPOP_DE() int {
	dasmOpReg("POP", "DE")
	regs.E = memRead(regs.SP)
	regs.D = memRead(regs.SP+1)
	regs.SP += 2
	return 3
}

func opPOP_HL() int {
	dasmOpReg("POP", "HL")
	regs.L = memRead(regs.SP)
	regs.H = memRead(regs.SP+1)
	regs.SP += 2
	return 3
}

//endregion

//region POP PSW (Pop processor status word)

//   (flags) <- ((SP)) ; (A) <- ((SP) + 1) ; (SP) <- (SP) + 2
//   The content of the memory location whose address is specified by the content of register SP is used to restore the condition
//   flags. The content of the memory location whose address is one more than the content of register SP is moved to register A.
//   The content of register SP is incremented by 2.
//   11110001
//   Cycles: 3  States: 10  Addressing: reg. indirect  Flags: Z,S,P,CY,AC

func opPOP_PSW() int {
	dasmOpReg("POP", "PSW")
	setFlags(memRead(regs.SP))
	regs.A = memRead(regs.SP+1)
	regs.SP += 2
	return 3
}

//endregion

//region OUT port (Output)

//   (data) <- (A)
//   The content of register A is placed on the eight bit bi-directional data bus for transmission to the specified port.
//   11010011
//   Cycles: 3  States: 10  Addressing: direct  Flags: none

func opOUT() int {
	port := memRead(regs.PC+1)
	dasmOpVal("OUT", port)
	if port == 0x02 || port == 0x04 {
		log.Fatal("16 bit shift register")
	}
	//fixme
	regs.PC++
	return 3
}

//endregion

//region EI (Enable interrupts)

//   The interrupt system is enabled following the execution of the next instruction.
//   11111011
//   Cycles: 1  States: 4  Flags: none

func opEI() int {
	dasmOp("EI")
	//fixme
	return 1
}

//endregion

//region NOP (No op)

//   No operation is performed. The registers and flags are unaffected.
//   00000000
//   Cycles: 1  States: 4  Flags: none

func opNOP() int {
	dasmOp("NOP")
	return 1
}

//endregion
