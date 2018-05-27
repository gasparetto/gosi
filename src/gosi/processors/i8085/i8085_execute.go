package i8085

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

func execute(fp func() int) int {

	return fp()
}

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
	r := fetch()
	dasmOpRegVal("MVI", "B", r)
	regs.B = r
	return 2
}

func opMVI_C() int {
	r := fetch()
	dasmOpRegVal("MVI", "C", r)
	regs.C = r
	return 2
}

func opMVI_D() int {
	r := fetch()
	dasmOpRegVal("MVI", "D", r)
	regs.D = r
	return 2
}

func opMVI_E() int {
	r := fetch()
	dasmOpRegVal("MVI", "E", r)
	regs.E = r
	return 2
}

func opMVI_H() int {
	r := fetch()
	dasmOpRegVal("MVI", "H", r)
	regs.H = r
	return 2
}

func opMVI_L() int {
	r := fetch()
	dasmOpRegVal("MVI", "L", r)
	regs.L = r
	return 2
}

func opMVI_A() int {
	r := fetch()
	dasmOpRegVal("MVI", "A", r)
	regs.A = r
	return 2
}

//endregion

//region MVI M, data (Move to memory immediate)

//   ((H) (L)) <- (byte 2)
//   The content of byte 2 of the instruction is moved to the memory location whose address is in registers H and L.
//   00110110
//   Cycles: 3  States: 10  Addressing: immed./reg. indirect  Flags: none

func opMVI_M() int {
	r := fetch()
	addr := getHL()
	dasmOpAddrVal("MVI", addr, r)
	memWrite(addr, r)
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
	rl := fetch()
	rh := fetch()
	dasmOpRegVal1Val2("LXI", "BC", rh, rl)
	regs.B = rh
	regs.C = rl
	return 3
}

func opLXI_DE() int {
	rl := fetch()
	rh := fetch()
	dasmOpRegVal1Val2("LXI", "DE", rh, rl)
	regs.D = rh
	regs.E = rl
	return 3
}

func opLXI_HL() int {
	rl := fetch()
	rh := fetch()
	dasmOpRegVal1Val2("LXI", "HL", rh, rl)
	regs.H = rh
	regs.L = rl
	return 3
}

func opLXI_SP() int {
	rp := fetch16()
	dasmOpRegAddr("LXI", "SP", rp)
	regs.SP = rp
	return 3
}

//endregion

//region LDA addr (Load Accumulator direct)

//   (A) <- ((byte 3)(byte 2))
//   The content of the memory location, whose address is specified in byte 2 and byte 3 of the instruction, is moved to register A.
//   00111010
//   Cycles: 4  States: 13  Addressing: direct  Flags: none

func opLDA() int {
	addr := fetch16()
	dasmOpAddr("LDA", addr)
	regs.A = memRead(addr)
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
	addr := fetch16()
	dasmOpAddr("STA", addr)
	memWrite(addr, regs.A)
	return 4
}

//endregion

//region LHLD addr (Load H and L direct)

//   (L) <- ((byte 3)(byte 2)) ; (H) <- ((byte 3)(byte 2) + 1)
//   The content of the memory location, whose address is specified in byte 2 and byte 3 of the instruction, is moved to register L.
//   The content of the memory location at the succeeding address is moved to register H.
//   00101010
//   Cycles: 5  States: 16  Addressing: direct  Flags: none

func opLHLD() int {
	addr := fetch16()
	dasmOpAddr("LHLD", addr)
	setHL(memRead16(addr))
	return 5
}

//endregion

//region SHLD rp (Store H and L direct)

//   ((byte 3)(byte 2)) <- (L) ; ((byte 3)(byte 2) + 1) <- (H)
//   The content of register L is moved to the memory location whose address is specified in byte 2 and byte 3.
//   The content of register H is moved to the succeeding memory location.
//   00100010
//   Cycles: 5  States: 16  Addressing: direct  Flags: none

func opSHLD() int {
	addr := fetch16()
	dasmOpAddr("SHLD", addr)
	memWrite16(addr, getHL())
	return 5
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

//region ADD r (Add Register)

//   (A) <- (A) + r
//   The content of register r is added to the content of the accumulator. The result is placed in the accumulator.
//   10000SSS
//   Cycles: 1  States: 4  Addressing: register  Flags: Z,S,P,CY,AC

func opADD_B() int {
	dasmOpReg("ADD", "B")
	regs.A = add_CY(regs.A, regs.B)
	return 1
}

func opADD_C() int {
	dasmOpReg("ADD", "C")
	regs.A = add_CY(regs.A, regs.C)
	return 1
}

func opADD_D() int {
	dasmOpReg("ADD", "D")
	regs.A = add_CY(regs.A, regs.D)
	return 1
}

func opADD_E() int {
	dasmOpReg("ADD", "E")
	regs.A = add_CY(regs.A, regs.E)
	return 1
}

func opADD_H() int {
	dasmOpReg("ADD", "H")
	regs.A = add_CY(regs.A, regs.H)
	return 1
}

func opADD_L() int {
	dasmOpReg("ADD", "L")
	regs.A = add_CY(regs.A, regs.L)
	return 1
}

func opADD_A() int {
	dasmOpReg("ADD", "A")
	regs.A = add_CY(regs.A, regs.A)
	return 1
}

//endregion

//region ADD M (Add memory)

//   (A) <- (A) + ((H) (L))
//   The content of the memory location whose address is contained in the H and L registers is added to the content of the
//   accumulator. The result is placed in the accumulator.
//   10000110
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: Z,S,P,CY,AC

func opADD_M() int {
	addr := getHL()
	dasmOpAddr("ADD", addr)
	regs.A = add_CY(regs.A, memRead(addr))
	return 2
}

//endregion

//region ADI data (Add immediate)

//   (A) <- (A) + (byte 2)
//   The content of the second byte of the instruction is added to the content of the accumulator. The result is placed in the
//   accumulator.
//   11000110
//   Cycles: 2  States: 7  Addressing: immediate  Flags: Z,S,P,CY,AC

func opADI() int {
	r := fetch()
	dasmOpVal("ADI", r)
	regs.A = add_CY(regs.A, r)
	return 2
}

//endregion

//region ADC r (Add Register with carry)

//   (A) <- (A) + r + (CY)
//   The content of register r and the content of the carry bit are added to the content of the accumulator.
//   The result is placed in the accumulator.
//   10001SSS
//   Cycles: 1  States: 4  Addressing: register  Flags: Z,S,P,CY,AC

func opADC_B() int {
	dasmOpReg("ADC", "B")
	regs.A = add3_CY(regs.A, regs.B, btoi(flags.CY))
	return 1
}

func opADC_C() int {
	dasmOpReg("ADC", "C")
	regs.A = add3_CY(regs.A, regs.C, btoi(flags.CY))
	return 1
}

func opADC_D() int {
	dasmOpReg("ADC", "D")
	regs.A = add3_CY(regs.A, regs.D, btoi(flags.CY))
	return 1
}

func opADC_E() int {
	dasmOpReg("ADC", "E")
	regs.A = add3_CY(regs.A, regs.E, btoi(flags.CY))
	return 1
}

func opADC_H() int {
	dasmOpReg("ADC", "H")
	regs.A = add3_CY(regs.A, regs.H, btoi(flags.CY))
	return 1
}

func opADC_L() int {
	dasmOpReg("ADC", "L")
	regs.A = add3_CY(regs.A, regs.L, btoi(flags.CY))
	return 1
}

func opADC_A() int {
	dasmOpReg("ADC", "A")
	regs.A = add3_CY(regs.A, regs.A, btoi(flags.CY))
	return 1
}

//endregion

//region ADC M (Add memory with carry)

//   (A) <- (A) + ((H) (L)) + (CY)
//   The content of the memory location whose address is contained in the H and L registers and the content of the CY flag are
//   added to the accumulator. The result is placed in the accumulator.
//   10001110
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: Z,S,P,CY,AC

func opADC_M() int {
	addr := getHL()
	dasmOpAddr("ADC", addr)
	regs.A = add3_CY(regs.A, memRead(addr), btoi(flags.CY))
	return 2
}

//endregion

//region ACI data (Add immediate with carry)

//   (A) <- (A) + (byte 2) + (CY)
//   The content of the second byte of the instruction and the content of the CY flag are added to the contents of the accumulator.
//   The result is placed in the accumulator.
//   11001110
//   Cycles: 2  States: 7  Addressing: immediate  Flags: Z,S,P,CY,AC

func opACI() int {
	r := fetch()
	dasmOpVal("ACI", r)
	regs.A = add3_CY(regs.A, r, btoi(flags.CY))
	return 2
}

//endregion

//region SUB r (Subtract Register)

//   (A) <- (A) - r
//   The content of register r is subtracted from the content of the accumulator. The result is placed in the accumulator.
//   10010SSS
//   Cycles: 1  States: 4  Addressing: register  Flags: Z,S,P,CY,AC

func opSUB_B() int {
	dasmOpReg("SUB", "B")
	regs.A = sub_CY(regs.A, regs.B)
	return 1
}

func opSUB_C() int {
	dasmOpReg("SUB", "C")
	regs.A = sub_CY(regs.A, regs.C)
	return 1
}

func opSUB_D() int {
	dasmOpReg("SUB", "D")
	regs.A = sub_CY(regs.A, regs.D)
	return 1
}

func opSUB_E() int {
	dasmOpReg("SUB", "E")
	regs.A = sub_CY(regs.A, regs.E)
	return 1
}

func opSUB_H() int {
	dasmOpReg("SUB", "H")
	regs.A = sub_CY(regs.A, regs.H)
	return 1
}

func opSUB_L() int {
	dasmOpReg("SUB", "L")
	regs.A = sub_CY(regs.A, regs.L)
	return 1
}

func opSUB_A() int {
	dasmOpReg("SUB", "A")
	regs.A = sub_CY(regs.A, regs.A)
	return 1
}

//endregion

//region SUB M (Subtract memory)

//   (A) <- (A) - ((H) (L))
//   The content of the memory location whose address is contained in the H and L registers is subtracted from the content of the
//   accumulator. The result is placed in the accumulator.
//   10010110
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: Z,S,P,CY,AC

func opSUB_M() int {
	addr := getHL()
	dasmOpAddr("SUB", addr)
	regs.A = sub_CY(regs.A, memRead(addr))
	return 2
}

//endregion

//region SUI data (Subtract immediate)

//   (A) <- (A) - (byte 2)
//   The content of the second byte of the instruction is subtracted from the content of the accumulator. The result is placed in the
//   accumulator.
//   11010110
//   Cycles: 2  States: 7  Addressing: immediate  Flags: Z,S,P,CY,AC

func opSUI() int {
	r := fetch()
	dasmOpVal("SUI", r)
	regs.A = sub_CY(regs.A, r)
	return 2
}

//endregion

//region SBB r (Subtract Register with borrow)

//   (A) <- (A) - r - (CY)
//   The content of register r and the content of the CY flag are both subtracted from the the accumulator.
//   The result is placed in the accumulator.
//   10011SSS
//   Cycles: 1  States: 4  Addressing: register  Flags: Z,S,P,CY,AC

func opSBB_B() int {
	dasmOpReg("SBB", "B")
	regs.A = sub3_CY(regs.A, regs.B, btoi(flags.CY))
	return 1
}

func opSBB_C() int {
	dasmOpReg("SBB", "C")
	regs.A = sub3_CY(regs.A, regs.C, btoi(flags.CY))
	return 1
}

func opSBB_D() int {
	dasmOpReg("SBB", "D")
	regs.A = sub3_CY(regs.A, regs.D, btoi(flags.CY))
	return 1
}

func opSBB_E() int {
	dasmOpReg("SBB", "E")
	regs.A = sub3_CY(regs.A, regs.E, btoi(flags.CY))
	return 1
}

func opSBB_H() int {
	dasmOpReg("SBB", "H")
	regs.A = sub3_CY(regs.A, regs.H, btoi(flags.CY))
	return 1
}

func opSBB_L() int {
	dasmOpReg("SBB", "L")
	regs.A = sub3_CY(regs.A, regs.L, btoi(flags.CY))
	return 1
}

func opSBB_A() int {
	dasmOpReg("SBB", "A")
	regs.A = sub3_CY(regs.A, regs.A, btoi(flags.CY))
	return 1
}

//endregion

//region SBB M (Subtract memory with borrow)

//   (A) <- (A) - ((H) (L)) - (CY)
//   The content of the memory location whose address is contained in the H and L registers and the content of the CY flag are
//   both subtracted from the accumulator. The result is placed in the accumulator.
//   10011110
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: Z,S,P,CY,AC

func opSBB_M() int {
	addr := getHL()
	dasmOpAddr("SBB", addr)
	regs.A = sub3_CY(regs.A, memRead(addr), btoi(flags.CY))
	return 2
}

//endregion

//region SBI data (Subtract immediate with borrow)

//   (A) <- (A) - (byte 2) - (CY)
//   The content of the second byte of the instruction and the content of the CY flag are both subtracted from the contents of the accumulator.
//   The result is placed in the accumulator.
//   11011110
//   Cycles: 2  States: 7  Addressing: immediate  Flags: Z,S,P,CY,AC

func opSBI() int {
	r := fetch()
	dasmOpVal("SBI", r)
	regs.A = sub3_CY(regs.A, r, btoi(flags.CY))
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

//region INR M (Increment memory)

//   ((H) (L)) <- ((H) (L)) + 1
//   The content of the memory location whose address is contained in the H and L registers is incremented by one.
//   00110100
//   Cycles: 3  States: 10  Addressing: reg. indirect  Flags: Z,S,P,AC

func opINR_M() int {
	addr := getHL()
	dasmOpAddr("INR", addr)
	v := memRead(addr)
	v = add(v, 1)
	memWrite(addr, v)
	return 3
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

//region DCR M (Decrement memory)

//   ((H) (L)) <- ((H) (L)) - 1
//   The content of the memory location whose address is contained in the H and L registers is decremented by one.
//   00110101
//   Cycles: 3  States: 10  Addressing: reg. indirect  Flags: Z,S,P,AC

func opDCR_M() int {
	addr := getHL()
	dasmOpAddr("DCR", addr)
	v := memRead(addr)
	v = sub(v, 1)
	memWrite(addr, v)
	return 3
}

//endregion

//region INX rp (Increment register pair)

//   (rh) (rl) <- (rh) (rl) + 1
//   The content of the register pair rp is incremented by one. Note: No condition flags are affected.
//   00RP0011
//   Cycles: 1  States: 5  Addressing: register  Flags: none

func opINX_BC() int {
	dasmOpReg("INX", "BC")
	r := getBC()
	r++
	setBC(r)
	return 1
}

func opINX_DE() int {
	dasmOpReg("INX", "DE")
	r := getDE()
	r++
	setDE(r)
	return 1
}

func opINX_HL() int {
	dasmOpReg("INX", "HL")
	r := getHL()
	r++
	setHL(r)
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
	r := getBC()
	r--
	setBC(r)
	return 1
}

func opDCX_DE() int {
	dasmOpReg("DCX", "DE")
	r := getDE()
	r--
	setDE(r)
	return 1
}

func opDCX_HL() int {
	dasmOpReg("DCX", "HL")
	r := getHL()
	r--
	setHL(r)
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
	r := uint32(getHL()) + uint32(getBC())
	setHL(uint16(r))
	flags.CY = (r & 0xffff0000) > 0
	return 3
}

func opDAD_DE() int {
	dasmOpReg("DAD", "DE")
	r := uint32(getHL()) + uint32(getDE())
	setHL(uint16(r))
	flags.CY = (r & 0xffff0000) > 0
	return 3
}

func opDAD_HL() int {
	dasmOpReg("DAD", "HL")
	r := uint32(getHL()) + uint32(getHL())
	setHL(uint16(r))
	flags.CY = (r & 0xffff0000) > 0
	return 3
}

func opDAD_SP() int {
	dasmOpReg("DAD", "SP")
	r := uint32(getHL()) + uint32(regs.SP)
	setHL(uint16(r))
	flags.CY = (r & 0xffff0000) > 0
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
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 1
}

func opANA_C() int {
	dasmOpReg("ANA", "C")
	regs.A = regs.A & regs.C
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 1
}

func opANA_D() int {
	dasmOpReg("ANA", "D")
	regs.A = regs.A & regs.D
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 1
}

func opANA_E() int {
	dasmOpReg("ANA", "E")
	regs.A = regs.A & regs.E
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 1
}

func opANA_H() int {
	dasmOpReg("ANA", "H")
	regs.A = regs.A & regs.H
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 1
}

func opANA_L() int {
	dasmOpReg("ANA", "L")
	regs.A = regs.A & regs.L
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 1
}

func opANA_A() int {
	dasmOpReg("ANA", "A")
	regs.A = regs.A & regs.A
	flags_Z_S_P(regs.A)
	flags.AC = false
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
	flags_Z_S_P(regs.A)
	//fixme flags.AC
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
	r := fetch()
	dasmOpVal("ANI", r)
	regs.A = regs.A & r
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
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

//todo XRI data

//region ORA r (OR Register)

//   (A) <- (A) V (r)
//   The content of register r is inclusive-or'd with the content of the accumulator. The result is placed in the accumulator.
//   The CY and AC flags are cleared.
//   10110SSS
//   Cycles: 1  States: 4  Addressing: register  Flags: Z,S,P,CY,AC

func opORA_B() int {
	dasmOpReg("ORA", "B")
	regs.A = regs.A | regs.B
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opORA_C() int {
	dasmOpReg("ORA", "C")
	regs.A = regs.A | regs.C
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opORA_D() int {
	dasmOpReg("ORA", "D")
	regs.A = regs.A | regs.D
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opORA_E() int {
	dasmOpReg("ORA", "E")
	regs.A = regs.A | regs.E
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opORA_H() int {
	dasmOpReg("ORA", "H")
	regs.A = regs.A | regs.H
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opORA_L() int {
	dasmOpReg("ORA", "L")
	regs.A = regs.A | regs.L
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opORA_A() int {
	dasmOpReg("ORA", "A")
	regs.A = regs.A | regs.A
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

//endregion

//region ORA M (OR Memory)

//   (A) <- (A) V ((H) (L))
//   The content of the memory location whose address is contained in the H and L registers is inclusive-OR'd with the content of
//   the accumulator. The result is placed in the accumulator. The CY and AC flags are cleared.
//   10110110
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: Z,S,P,CY,AC

func opORA_M() int {
	addr := getHL()
	dasmOpAddr("ORA", addr)
	regs.A = regs.A | memRead(addr)
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

// endregion

//region ORI data (OR immediate)

//   (A) <- (A) V (byte 2)
//   The content of the second byte of the instruction is inclusive-OR'd with the contents of the accumulator.
//   The result is placed in the accumulator. The CY and AC flags are cleared.
//   11100110
//   Cycles: 2  States: 7  Addressing: immediate  Flags: Z,S,P,CY,AC

func opORI() int {
	r := fetch()
	dasmOpVal("ORI", r)
	regs.A = regs.A | r
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

//endregion

//region CMP r (Compare Register)

//   (A) - (r)
//   The content of register r is subtracted from the accumulator. The accumulator remains unchanged. The condition flags are set
//   as a result of the subtraction. The Z flag is set to 1 if (A) = (r). The CY flag is set to 1 if (A) < (r).
//   10111SSS
//   Cycles: 1  States: 4  Addressing: register  Flags: Z,S,P,CY,AC

func opCMP_B() int {
	dasmOpReg("CMP", "B")
	sub(regs.A, regs.B)
	flags.CY = regs.A < regs.B
	return 1
}

func opCMP_C() int {
	dasmOpReg("CMP", "C")
	sub(regs.A, regs.C)
	flags.CY = regs.A < regs.C
	return 1
}

func opCMP_D() int {
	dasmOpReg("CMP", "D")
	sub(regs.A, regs.D)
	flags.CY = regs.A < regs.D
	return 1
}

func opCMP_E() int {
	dasmOpReg("CMP", "E")
	sub(regs.A, regs.E)
	flags.CY = regs.A < regs.E
	return 1
}

func opCMP_H() int {
	dasmOpReg("CMP", "H")
	sub(regs.A, regs.H)
	flags.CY = regs.A < regs.H
	return 1
}

func opCMP_L() int {
	dasmOpReg("CMP", "L")
	sub(regs.A, regs.L)
	flags.CY = regs.A < regs.L
	return 1
}

func opCMP_A() int {
	dasmOpReg("CMP", "A")
	sub(regs.A, regs.A)
	flags.CY = regs.A < regs.A
	return 1
}

//endregion

//region CMP M (Compare memory)

//   (A) - ((H) (L))
//   The content of the memory location whose address is contained in the H and L registers is subtracted from the accumulator.
//   The accumulator remains unchanged. The condition flags are set as a result of the subtraction. The Z flag is set to 1 if
//   (A) = ((H) (L)). The CY flag is set to 1 if (A) < ((H) (L)).
//   10111110
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: Z,S,P,CY,AC

func opCMP_M() int {
	addr := getHL()
	dasmOpAddr("CMP", addr)
	v := memRead(addr)
	sub(regs.A, v)
	flags.CY = regs.A < v
	return 2
}

//endregion

//region CPI data (Compare immediate)

//   (A) - (byte 2)
//   The content of the second byte of the instruction is subtracted from the accumulator. The condition flags are set by the
//   result of the subtraction. The Z flag is set to 1 if (A) = (byte 2). The CY flag is set to 1 if (A) < (byte 2).
//   11111110
//   Cycles: 2  States: 7  Addressing: immediate  Flags: Z,S,P,CY,AC

func opCPI() int {
	r := fetch()
	dasmOpVal("CPI", r)
	sub_CY(regs.A, r)
	return 2
}

//endregion

//region RLC (Rotate left)

//   (An+1) <- (An) ; (A0) <- (A7) ; (CY) <- (A7)
//   The content of the accumulator is rotated left one position. The low order bit and the CY flag are both set to the value
//   shifted out of the high order bit position. Only the CY flag is affected.
//   00000111
//   Cycles: 1  States: 4  Flags: CY

func opRLC() int {
	dasmOp("RLC")
	a0 := regs.A & 0x80
	regs.A = (regs.A << 1 & 0xfe) | a0>>7
	flags.CY = a0 != 0
	return 1
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
	regs.A = (regs.A >> 1 & 0x7f) | a0<<7
	flags.CY = a0 != 0
	return 1
}

//endregion

//region RAL (Rotate left through carry)

//   (An+1) <- (An) ; (CY) <- (A7) ; (A0) <- (CY)
//   The content of the accumulator is rotated left one position through the CY flag. The low order bit is set equal to the CY
//   flag and the CY flag is set to the value shifted out of the high order bit. Only the CY flag is affected.
//   00010111
//   Cycles: 1  States: 4  Flags: CY

func opRAL() int {
	dasmOp("RAL")
	a0 := regs.A & 0x80
	regs.A = (regs.A << 1 & 0xfe) | btoi(flags.CY)
	flags.CY = a0 != 0
	return 1
}

//endregion

//region RAR (Rotate right through carry)

//   (An) <- (An-1) ; (CY) <- (A0) ; (A7) <- (CY)
//   The content of the accumulator is rotated right one position through the CY flag. The high order bit is set to the CY flag
//   and the CY flag is set to the value shifted out of the low order bit. Only the CY flag is affected.
//   00011111
//   Cycles: 1  States: 4  Flags: CY

func opRAR() int {
	dasmOp("RAR")
	a0 := regs.A & 0x01
	regs.A = (regs.A >> 1 & 0x7f) | btoi(flags.CY)<<7
	flags.CY = a0 != 0
	return 1
}

//endregion


//region CMA (Complement accumulator)

//   (A) <- (!A)
//   The contents of the accumulator are complemented (zero bits become 1, one bits become 0). No flags are affected.
//   00101111
//   Cycles: 1  States: 4  Flags: none

func opCMA() int {
	dasmOp("CMA")
	regs.A = ^regs.A
	return 1
}

//endregion

//todo CMC

//region STC (Set carry)

//   (CY) <- 1
//   The CY flag is set to 1.
//   00110111
//   Cycles: 1  States: 4  Flags: CY

func opSTC() int {
	dasmOp("STC")
	flags.CY = true
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
	addr := fetch16()
	dasmOpAddr("JMP", addr)
	regs.PC = addr
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
	addr := fetch16()
	dasmOpAddr("JNZ", addr)
	if !flags.Z {
		regs.PC = addr
	}
	return 3
}

func opJZ() int {
	addr := fetch16()
	dasmOpAddr("JZ", addr)
	if flags.Z {
		regs.PC = addr
	}
	return 3
}

func opJNC() int {
	addr := fetch16()
	dasmOpAddr("JNC", addr)
	if !flags.CY {
		regs.PC = addr
	}
	return 3
}

func opJC() int {
	addr := fetch16()
	dasmOpAddr("JC", addr)
	if flags.CY {
		regs.PC = addr
	}
	return 3
}

func opJPO() int {
	addr := fetch16()
	dasmOpAddr("JPO", addr)
	if !flags.P {
		regs.PC = addr
	}
	return 3
}

func opJPE() int {
	addr := fetch16()
	dasmOpAddr("JPE", addr)
	if flags.P {
		regs.PC = addr
	}
	return 3
}

func opJP() int {
	addr := fetch16()
	dasmOpAddr("JP", addr)
	if !flags.S {
		regs.PC = addr
	}
	return 3
}

func opJM() int {
	addr := fetch16()
	dasmOpAddr("JM", addr)
	if flags.S {
		regs.PC = addr
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
	addr := fetch16()
	dasmOpAddr("CALL", addr)
	memWrite16(regs.SP-2, regs.PC)
	regs.SP -= 2
	regs.PC = addr
	return 5
}

//endregion

//region Ccondition (Conditional call)

//   If (CCC): ((SP) - 1) <- (PCH) ; ((SP) - 2) <- (PCL) ; (SP) <- (SP) - 2 ; (PC) <- (byte 3) (byte2)
//   If the specified condition is true, the actions specified in the CALL instruction (see above) are performed; otherwise,
//   control continues sequentially.
//   11CCC100
//   Cycles: 3/5  States: 11/17  Addressing: immediate/reg. indirect  Flags: none

func opCNZ() int {
	addr := fetch16()
	dasmOpAddr("CNZ", addr)
	if !flags.Z {
		memWrite16(regs.SP-2, regs.PC)
		regs.SP -= 2
		regs.PC = addr
		return 5
	}
	return 3
}

func opCZ() int {
	addr := fetch16()
	dasmOpAddr("CZ", addr)
	if flags.Z {
		memWrite16(regs.SP-2, regs.PC)
		regs.SP -= 2
		regs.PC = addr
		return 5
	}
	return 3
}

func opCNC() int {
	addr := fetch16()
	dasmOpAddr("CNC", addr)
	if !flags.CY {
		memWrite16(regs.SP-2, regs.PC)
		regs.SP -= 2
		regs.PC = addr
		return 5
	}
	return 3
}

func opCC() int {
	addr := fetch16()
	dasmOpAddr("CC", addr)
	if flags.CY {
		memWrite16(regs.SP-2, regs.PC)
		regs.SP -= 2
		regs.PC = addr
		return 5
	}
	return 3
}

func opCPO() int {
	addr := fetch16()
	dasmOpAddr("CPO", addr)
	if !flags.P {
		memWrite16(regs.SP-2, regs.PC)
		regs.SP -= 2
		regs.PC = addr
		return 5
	}
	return 3
}

func opCPE() int {
	addr := fetch16()
	dasmOpAddr("CPE", addr)
	if flags.P {
		memWrite16(regs.SP-2, regs.PC)
		regs.SP -= 2
		regs.PC = addr
		return 5
	}
	return 3
}

func opCP() int {
	addr := fetch16()
	dasmOpAddr("CP", addr)
	if !flags.S {
		memWrite16(regs.SP-2, regs.PC)
		regs.SP -= 2
		regs.PC = addr
		return 5
	}
	return 3
}

func opCM() int {
	addr := fetch16()
	dasmOpAddr("CM", addr)
	if flags.S {
		memWrite16(regs.SP-2, regs.PC)
		regs.SP -= 2
		regs.PC = addr
		return 5
	}
	return 3
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
	regs.PC = memRead16(regs.SP)
	regs.SP += 2
	return 5
}

//endregion

//region Rcondition (Conditional return)

//   If (CCC): (PCL) <- ((SP)) ; (PCH) <- ((SP) + 1) ; (SP) <- (SP) + 2
//   If the specified condition is true, the actions specified in the RET instruction (see above) are performed; otherwise,
//   control continues sequentially.
//   11CCC000
//   Cycles: 1/3  States: 5/11  Addressing: reg. indirect  Flags: none

func opRNZ() int {
	dasmOp("RNZ")
	if !flags.Z {
		regs.PC = memRead16(regs.SP)
		regs.SP += 2
		return 3
	}
	return 1
}

func opRZ() int {
	dasmOp("RZ")
	if flags.Z {
		regs.PC = memRead16(regs.SP)
		regs.SP += 2
		return 3
	}
	return 1
}

func opRNC() int {
	dasmOp("RNC")
	if !flags.CY {
		regs.PC = memRead16(regs.SP)
		regs.SP += 2
		return 3
	}
	return 1
}

func opRC() int {
	dasmOp("RC")
	if flags.CY {
		regs.PC = memRead16(regs.SP)
		regs.SP += 2
		return 3
	}
	return 1
}

func opRPO() int {
	dasmOp("RPO")
	if !flags.P {
		regs.PC = memRead16(regs.SP)
		regs.SP += 2
		return 3
	}
	return 1
}

func opRPE() int {
	dasmOp("RPE")
	if flags.P {
		regs.PC = memRead16(regs.SP)
		regs.SP += 2
		return 3
	}
	return 1
}

func opRP() int {
	dasmOp("RP")
	if !flags.S {
		regs.PC = memRead16(regs.SP)
		regs.SP += 2
		return 3
	}
	return 1
}

func opRM() int {
	dasmOp("RM")
	if flags.S {
		regs.PC = memRead16(regs.SP)
		regs.SP += 2
		return 3
	}
	return 1
}

//endregion

//region RST n (Restart)

//   ((SP) - 1) <- (PCH) ; ((SP) - 2) <- (PCL) ; (SP) <- (SP) - 2 ; (PC) <- 8 * (NNN)
//   The high-order eight bits of the next instruction address are moved to the memory location whose address is one less than the
//   content of register SP. The low-order eight bits of the next instruction address are moved to the memory location whose
//   address is two less than the content of register SP. The content of register SP is decremented by two. Control is transferred
//   to the instruction whose address is eight times the content of NNN.
//   11NNN111
//   Cycles: 3  States: 11  Addressing: reg. indirect  Flags: none

func opRST_0() int {
	dasmOpReg("RST", "0")
	memWrite16(regs.SP-2, regs.PC)
	regs.SP -= 2
	regs.PC = 0x0000
	return 3
}

func opRST_1() int {
	dasmOpReg("RST", "1")
	memWrite16(regs.SP-2, regs.PC)
	regs.SP -= 2
	regs.PC = 0x0008
	return 3
}

func opRST_2() int {
	dasmOpReg("RST", "2")
	memWrite16(regs.SP-2, regs.PC)
	regs.SP -= 2
	regs.PC = 0x0010
	return 3
}

func opRST_3() int {
	dasmOpReg("RST", "3")
	memWrite16(regs.SP-2, regs.PC)
	regs.SP -= 2
	regs.PC = 0x0018
	return 3
}

func opRST_4() int {
	dasmOpReg("RST", "4")
	memWrite16(regs.SP-2, regs.PC)
	regs.SP -= 2
	regs.PC = 0x0020
	return 3
}

func opRST_5() int {
	dasmOpReg("RST", "5")
	memWrite16(regs.SP-2, regs.PC)
	regs.SP -= 2
	regs.PC = 0x0028
	return 3
}

func opRST_6() int {
	dasmOpReg("RST", "6")
	memWrite16(regs.SP-2, regs.PC)
	regs.SP -= 2
	regs.PC = 0x0030
	return 3
}

func opRST_7() int {
	dasmOpReg("RST", "7")
	memWrite16(regs.SP-2, regs.PC)
	regs.SP -= 2
	regs.PC = 0x0038
	return 3
}

//endregion

//region PCHL (Jump H and L indirect - move H and L to PC)

//   (PCH) <- (H) ; (PCL) <- (L)
//   The content of register H is moved to the high-order eight bits of register PC. The content of register l is moved to the
//   low-order eight bits of register PC.
//   11101001
//   Cycles: 1  States: 5  Addressing: register  Flags: none

func opPCHL() int {
	dasmOp("PCHL")
	regs.PC = getHL()
	return 1
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
	regs.B = memRead(regs.SP + 1)
	regs.SP += 2
	return 3
}

func opPOP_DE() int {
	dasmOpReg("POP", "DE")
	regs.E = memRead(regs.SP)
	regs.D = memRead(regs.SP + 1)
	regs.SP += 2
	return 3
}

func opPOP_HL() int {
	dasmOpReg("POP", "HL")
	regs.L = memRead(regs.SP)
	regs.H = memRead(regs.SP + 1)
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
	regs.A = memRead(regs.SP + 1)
	regs.SP += 2
	return 3
}

//endregion

//region XTHL (Exchange stack top with H and L)

//   (L) <-> ((SP)) ; (H) <-> ((SP + 1))
//   The content of the L register is exchanged with the content of the memory location whose address is specified by the content
//   of register SP. The content of the H register is exchanged with the content of the memory location whose address is one more
//   than the content of register SP.
//   11100011
//   Cycles: 5  States: 18  Addressing: reg. indirect  Flags: none

func opXTHL() int {
	dasmOp("XTHL")
	addr := regs.SP
	r1 := regs.H
	regs.H = memRead(addr+1)
	memWrite(addr+1, r1)
	r2 := regs.L
	regs.L = memRead(addr)
	memWrite(addr, r2)
	return 5
}

//endregion

//todo SPHL

//region IN port (Input)

//   (A) <- (data)
//   The data placed on the eight bit bi-directional data bus by the specified port is moved to register A.
//   11011011
//   Cycles: 3  States: 10  Addressing: direct  Flags: none

func opIN() int {
	port := fetch()
	dasmOpVal("IN", port)
	fp, ok := ports_in[port]
	if ok {
		regs.A = fp()
	}
	return 3
}

//endregion

//region OUT port (Output)

//   (data) <- (A)
//   The content of register A is placed on the eight bit bi-directional data bus for transmission to the specified port.
//   11010011
//   Cycles: 3  States: 10  Addressing: direct  Flags: none

func opOUT() int {
	port := fetch()
	dasmOpVal("OUT", port)
	fp, ok := ports_out[port]
	if ok {
		fp(regs.A)
	}
	return 3
}

//endregion

//region EI (Enable interrupts)

//   The interrupt system is enabled following the execution of the next instruction.
//   11111011
//   Cycles: 1  States: 4  Flags: none

func opEI() int {
	dasmOp("EI")
	interrupt = true
	return 1
}

//endregion

//region DI (Disable interrupts)

//   The interrupt system is disabled immediately following the execution of the DI instruction.
//   1110011
//   Cycles: 1  States: 4  Flags: none

func opDI() int {
	dasmOp("DI")
	interrupt = false
	return 1
}

//endregion

//todo HLT

//region NOP (No op)

//   No operation is performed. The registers and flags are unaffected.
//   00000000
//   Cycles: 1  States: 4  Flags: none

func opNOP() int {
	dasmOp("NOP")
	return 1
}

//endregion

// ********** Utils **********

func add(a uint8, b uint8) uint8 {
	res16 := uint16(a) + uint16(b)
	res8 := uint8(res16)
	flags_Z_S_P(res8)
	flags.AC = (res16 & 0xff00) != 0
	return res8
}

func add_CY(a uint8, b uint8) uint8 {
	res16 := uint16(a) + uint16(b)
	res8 := uint8(res16)
	flags_Z_S_P(res8)
	flags.AC = (res16 & 0xff00) != 0
	//todo flags.CY
	return res8
}

func add3_CY(a uint8, b uint8, c uint8) uint8 {
	res16 := uint16(a) + uint16(b) + uint16(c)
	res8 := uint8(res16)
	flags_Z_S_P(res8)
	flags.AC = (res16 & 0xff00) != 0
	//todo flags.CY
	return res8
}

func sub(a uint8, b uint8) uint8 {
	res16 := uint16(a) - uint16(b)
	res8 := uint8(res16)
	flags_Z_S_P(res8)
	flags.AC = (res16 & 0xff00) != 0
	return res8
}

func sub_CY(a uint8, b uint8) uint8 {
	res16 := uint16(a) - uint16(b)
	res8 := uint8(res16)
	flags_Z_S_P(res8)
	flags.AC = (res16 & 0xff00) != 0
	//todo flags.CY
	return res8
}

func sub3_CY(a uint8, b uint8, c uint8) uint8 {
	res16 := uint16(a) - uint16(b) - uint16(c)
	res8 := uint8(res16)
	flags_Z_S_P(res8)
	flags.AC = (res16 & 0xff00) != 0
	//todo flags.CY
	return res8
}

func flags_Z_S_P(v uint8) {
	flags.Z = v == 0
	flags.S = (v & 0x80) != 0
	flags.P = evenParity(v)
}

func evenParity(v uint8) bool {
	b0 := btoi((v & 0x01) != 0)
	b1 := btoi((v & 0x02) != 0)
	b2 := btoi((v & 0x04) != 0)
	b3 := btoi((v & 0x08) != 0)
	b4 := btoi((v & 0x10) != 0)
	b5 := btoi((v & 0x20) != 0)
	b6 := btoi((v & 0x40) != 0)
	b7 := btoi((v & 0x80) != 0)
	r := b7 + b6 + b5 + b4 + b3 + b2 + b1 + b0
	return (r & 0x01) == 0
}
