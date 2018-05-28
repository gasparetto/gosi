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
	//regs.B = regs.B
	return 1
}

func opMOV_B_C() int {
	regs.B = regs.C
	return 1
}

func opMOV_B_D() int {
	regs.B = regs.D
	return 1
}

func opMOV_B_E() int {
	regs.B = regs.E
	return 1
}

func opMOV_B_H() int {
	regs.B = regs.H
	return 1
}

func opMOV_B_L() int {
	regs.B = regs.L
	return 1
}

func opMOV_B_A() int {
	regs.B = regs.A
	return 1
}

func opMOV_C_B() int {
	regs.C = regs.B
	return 1
}

func opMOV_C_C() int {
	//regs.C = regs.C
	return 1
}

func opMOV_C_D() int {
	regs.C = regs.D
	return 1
}

func opMOV_C_E() int {
	regs.C = regs.E
	return 1
}

func opMOV_C_H() int {
	regs.C = regs.H
	return 1
}

func opMOV_C_L() int {
	regs.C = regs.L
	return 1
}

func opMOV_C_A() int {
	regs.C = regs.A
	return 1
}

func opMOV_D_B() int {
	regs.D = regs.B
	return 1
}

func opMOV_D_C() int {
	regs.D = regs.C
	return 1
}

func opMOV_D_D() int {
	//regs.D = regs.D
	return 1
}

func opMOV_D_E() int {
	regs.D = regs.E
	return 1
}

func opMOV_D_H() int {
	regs.D = regs.H
	return 1
}

func opMOV_D_L() int {
	regs.D = regs.L
	return 1
}

func opMOV_D_A() int {
	regs.D = regs.A
	return 1
}

func opMOV_E_B() int {
	regs.E = regs.B
	return 1
}

func opMOV_E_C() int {
	regs.E = regs.C
	return 1
}

func opMOV_E_D() int {
	regs.E = regs.D
	return 1
}

func opMOV_E_E() int {
	//regs.E = regs.E
	return 1
}

func opMOV_E_H() int {
	regs.E = regs.H
	return 1
}

func opMOV_E_L() int {
	regs.E = regs.L
	return 1
}

func opMOV_E_A() int {
	regs.E = regs.A
	return 1
}

func opMOV_H_B() int {
	regs.H = regs.B
	return 1
}

func opMOV_H_C() int {
	regs.H = regs.C
	return 1
}

func opMOV_H_D() int {
	regs.H = regs.D
	return 1
}

func opMOV_H_E() int {
	regs.H = regs.E
	return 1
}

func opMOV_H_H() int {
	//regs.H = regs.H
	return 1
}

func opMOV_H_L() int {
	regs.H = regs.L
	return 1
}

func opMOV_H_A() int {
	regs.H = regs.A
	return 1
}

func opMOV_L_B() int {
	regs.L = regs.B
	return 1
}

func opMOV_L_C() int {
	regs.L = regs.C
	return 1
}

func opMOV_L_D() int {
	regs.L = regs.D
	return 1
}

func opMOV_L_E() int {
	regs.L = regs.E
	return 1
}

func opMOV_L_H() int {
	regs.L = regs.H
	return 1
}

func opMOV_L_L() int {
	//regs.L = regs.L
	return 1
}

func opMOV_L_A() int {
	regs.L = regs.A
	return 1
}

func opMOV_A_B() int {
	regs.A = regs.B
	return 1
}

func opMOV_A_C() int {
	regs.A = regs.C
	return 1
}

func opMOV_A_D() int {
	regs.A = regs.D
	return 1
}

func opMOV_A_E() int {
	regs.A = regs.E
	return 1
}

func opMOV_A_H() int {
	regs.A = regs.H
	return 1
}

func opMOV_A_L() int {
	regs.A = regs.L
	return 1
}

func opMOV_A_A() int {
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
	regs.B = memRead(getHL())
	return 2
}

func opMOV_C_M() int {
	regs.C = memRead(getHL())
	return 2
}

func opMOV_D_M() int {
	regs.D = memRead(getHL())
	return 2
}

func opMOV_E_M() int {
	regs.E = memRead(getHL())
	return 2
}

func opMOV_H_M() int {
	regs.H = memRead(getHL())
	return 2
}

func opMOV_L_M() int {
	regs.L = memRead(getHL())
	return 2
}

func opMOV_A_M() int {
	regs.A = memRead(getHL())
	return 2
}

//endregion

//region MOV M, r (Move to memory)

//   ((H) (L)) <- (r)
//   The content of register r is moved to the memory location whose address is in registers H and L.
//   01110SSS
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: none

func opMOV_M_B() int {
	memWrite(getHL(), regs.B)
	return 2
}

func opMOV_M_C() int {
	memWrite(getHL(), regs.C)
	return 2
}

func opMOV_M_D() int {
	memWrite(getHL(), regs.D)
	return 2
}

func opMOV_M_E() int {
	memWrite(getHL(), regs.E)
	return 2
}

func opMOV_M_H() int {
	memWrite(getHL(), regs.H)
	return 2
}

func opMOV_M_L() int {
	memWrite(getHL(), regs.L)
	return 2
}

func opMOV_M_A() int {
	memWrite(getHL(), regs.A)
	return 2
}

//endregion

//region MVI r, data (Move Immediate)

//   (r) <- (byte 2)
//   The content of byte 2 of the instruction is moved to register r.
//   00DDD110
//   Cycles: 2  States: 7  Addressing: immediate  Flags: none

func opMVI_B() int {
	regs.B = fetch()
	return 2
}

func opMVI_C() int {
	regs.C = fetch()
	return 2
}

func opMVI_D() int {
	regs.D = fetch()
	return 2
}

func opMVI_E() int {
	regs.E = fetch()
	return 2
}

func opMVI_H() int {
	regs.H = fetch()
	return 2
}

func opMVI_L() int {
	regs.L = fetch()
	return 2
}

func opMVI_A() int {
	regs.A = fetch()
	return 2
}

//endregion

//region MVI M, data (Move to memory immediate)

//   ((H) (L)) <- (byte 2)
//   The content of byte 2 of the instruction is moved to the memory location whose address is in registers H and L.
//   00110110
//   Cycles: 3  States: 10  Addressing: immed./reg. indirect  Flags: none

func opMVI_M() int {
	memWrite(getHL(), fetch())
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
	regs.C = fetch()
	regs.B = fetch()
	return 3
}

func opLXI_DE() int {
	regs.E = fetch()
	regs.D = fetch()
	return 3
}

func opLXI_HL() int {
	regs.L = fetch()
	regs.H = fetch()
	return 3
}

func opLXI_SP() int {
	regs.SP = fetch16()
	return 3
}

//endregion

//region LDA addr (Load Accumulator direct)

//   (A) <- ((byte 3)(byte 2))
//   The content of the memory location, whose address is specified in byte 2 and byte 3 of the instruction, is moved to register A.
//   00111010
//   Cycles: 4  States: 13  Addressing: direct  Flags: none

func opLDA() int {
	regs.A = memRead(fetch16())
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
	memWrite(fetch16(), regs.A)
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
	setHL(memRead16(fetch16()))
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
	memWrite16(fetch16(), getHL())
	return 5
}

//endregion

//region LDAX rp (Load accumulator indirect)

//   (A) <- ((rp))
//   The content of the memory location, whose address is in the register pair rp, is moved to register A.
//   00RP1010
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: none

func opLDAX_BC() int {
	regs.A = memRead(getBC())
	return 2
}

func opLDAX_DE() int {
	regs.A = memRead(getDE())
	return 2
}

//endregion

//region STAX rp (Store accumulator indirect)

//   ((rp)) <- (A)
//   The content of register A is moved to the memory location whose address is in the register pair rp.
//   00RP0010
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: none

func opSTAX_BC() int {
	memWrite(getBC(), regs.A)
	return 2
}

func opSTAX_DE() int {
	memWrite(getDE(), regs.A)
	return 2
}

//endregion

//region XCHG (Exchange H and L with D and E)

//   (H) <-> (D) ; (L) <-> (E)
//   The contents of registers H and L are exchanged with the contents of registers D and E.
//   11101011
//   Cycles: 1  States: 4  Addressing: register  Flags: none

func opXCHG() int {
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
	regs.A = add_CY(regs.A, regs.B)
	return 1
}

func opADD_C() int {
	regs.A = add_CY(regs.A, regs.C)
	return 1
}

func opADD_D() int {
	regs.A = add_CY(regs.A, regs.D)
	return 1
}

func opADD_E() int {
	regs.A = add_CY(regs.A, regs.E)
	return 1
}

func opADD_H() int {
	regs.A = add_CY(regs.A, regs.H)
	return 1
}

func opADD_L() int {
	regs.A = add_CY(regs.A, regs.L)
	return 1
}

func opADD_A() int {
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
	regs.A = add_CY(regs.A, memRead(getHL()))
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
	regs.A = add_CY(regs.A, fetch())
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
	regs.A = add3_CY(regs.A, regs.B, btoi(flags.CY))
	return 1
}

func opADC_C() int {
	regs.A = add3_CY(regs.A, regs.C, btoi(flags.CY))
	return 1
}

func opADC_D() int {
	regs.A = add3_CY(regs.A, regs.D, btoi(flags.CY))
	return 1
}

func opADC_E() int {
	regs.A = add3_CY(regs.A, regs.E, btoi(flags.CY))
	return 1
}

func opADC_H() int {
	regs.A = add3_CY(regs.A, regs.H, btoi(flags.CY))
	return 1
}

func opADC_L() int {
	regs.A = add3_CY(regs.A, regs.L, btoi(flags.CY))
	return 1
}

func opADC_A() int {
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
	regs.A = add3_CY(regs.A, memRead(getHL()), btoi(flags.CY))
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
	regs.A = add3_CY(regs.A, fetch(), btoi(flags.CY))
	return 2
}

//endregion

//region SUB r (Subtract Register)

//   (A) <- (A) - r
//   The content of register r is subtracted from the content of the accumulator. The result is placed in the accumulator.
//   10010SSS
//   Cycles: 1  States: 4  Addressing: register  Flags: Z,S,P,CY,AC

func opSUB_B() int {
	regs.A = sub_CY(regs.A, regs.B)
	return 1
}

func opSUB_C() int {
	regs.A = sub_CY(regs.A, regs.C)
	return 1
}

func opSUB_D() int {
	regs.A = sub_CY(regs.A, regs.D)
	return 1
}

func opSUB_E() int {
	regs.A = sub_CY(regs.A, regs.E)
	return 1
}

func opSUB_H() int {
	regs.A = sub_CY(regs.A, regs.H)
	return 1
}

func opSUB_L() int {
	regs.A = sub_CY(regs.A, regs.L)
	return 1
}

func opSUB_A() int {
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
	regs.A = sub_CY(regs.A, memRead(getHL()))
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
	regs.A = sub_CY(regs.A, fetch())
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
	regs.A = sub3_CY(regs.A, regs.B, btoi(flags.CY))
	return 1
}

func opSBB_C() int {
	regs.A = sub3_CY(regs.A, regs.C, btoi(flags.CY))
	return 1
}

func opSBB_D() int {
	regs.A = sub3_CY(regs.A, regs.D, btoi(flags.CY))
	return 1
}

func opSBB_E() int {
	regs.A = sub3_CY(regs.A, regs.E, btoi(flags.CY))
	return 1
}

func opSBB_H() int {
	regs.A = sub3_CY(regs.A, regs.H, btoi(flags.CY))
	return 1
}

func opSBB_L() int {
	regs.A = sub3_CY(regs.A, regs.L, btoi(flags.CY))
	return 1
}

func opSBB_A() int {
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
	regs.A = sub3_CY(regs.A, memRead(getHL()), btoi(flags.CY))
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
	regs.A = sub3_CY(regs.A, fetch(), btoi(flags.CY))
	return 2
}

//endregion

//region INR r (Increment Register)

//   (r) <- (r) + 1
//   The content of register r is incremented by one.
//   00DDD100
//   Cycles: 1  States: 5  Addressing: register  Flags: Z,S,P,AC

func opINR_B() int {
	regs.B = add(regs.B, 1)
	return 1
}

func opINR_C() int {
	regs.C = add(regs.C, 1)
	return 1
}

func opINR_D() int {
	regs.D = add(regs.D, 1)
	return 1
}

func opINR_E() int {
	regs.E = add(regs.E, 1)
	return 1
}

func opINR_H() int {
	regs.H = add(regs.H, 1)
	return 1
}

func opINR_L() int {
	regs.L = add(regs.L, 1)
	return 1
}

func opINR_A() int {
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
	regs.B = sub(regs.B, 1)
	return 1
}

func opDCR_C() int {
	regs.C = sub(regs.C, 1)
	return 1
}

func opDCR_D() int {
	regs.D = sub(regs.D, 1)
	return 1
}

func opDCR_E() int {
	regs.E = sub(regs.E, 1)
	return 1
}

func opDCR_H() int {
	regs.H = sub(regs.H, 1)
	return 1
}

func opDCR_L() int {
	regs.L = sub(regs.L, 1)
	return 1
}

func opDCR_A() int {
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
	r := getBC()
	r++
	setBC(r)
	return 1
}

func opINX_DE() int {
	r := getDE()
	r++
	setDE(r)
	return 1
}

func opINX_HL() int {
	r := getHL()
	r++
	setHL(r)
	return 1
}

func opINX_SP() int {
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
	r := getBC()
	r--
	setBC(r)
	return 1
}

func opDCX_DE() int {
	r := getDE()
	r--
	setDE(r)
	return 1
}

func opDCX_HL() int {
	r := getHL()
	r--
	setHL(r)
	return 1
}

func opDCX_SP() int {
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
	r := uint32(getHL()) + uint32(getBC())
	setHL(uint16(r))
	flags.CY = (r & 0xffff0000) > 0
	return 3
}

func opDAD_DE() int {
	r := uint32(getHL()) + uint32(getDE())
	setHL(uint16(r))
	flags.CY = (r & 0xffff0000) > 0
	return 3
}

func opDAD_HL() int {
	r := uint32(getHL()) + uint32(getHL())
	setHL(uint16(r))
	flags.CY = (r & 0xffff0000) > 0
	return 3
}

func opDAD_SP() int {
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
	regs.A = regs.A & regs.B
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 1
}

func opANA_C() int {
	regs.A = regs.A & regs.C
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 1
}

func opANA_D() int {
	regs.A = regs.A & regs.D
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 1
}

func opANA_E() int {
	regs.A = regs.A & regs.E
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 1
}

func opANA_H() int {
	regs.A = regs.A & regs.H
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 1
}

func opANA_L() int {
	regs.A = regs.A & regs.L
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 1
}

func opANA_A() int {
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
	v := memRead(getHL())
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
	regs.A = regs.A & fetch()
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
	regs.A = regs.A ^ regs.B
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opXRA_C() int {
	regs.A = regs.A ^ regs.C
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opXRA_D() int {
	regs.A = regs.A ^ regs.D
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opXRA_E() int {
	regs.A = regs.A ^ regs.E
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opXRA_H() int {
	regs.A = regs.A ^ regs.H
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opXRA_L() int {
	regs.A = regs.A ^ regs.L
	flags_Z_S_P(regs.A)
	flags.AC = false
	flags.CY = false
	return 2
}

func opXRA_A() int {
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
	regs.A = regs.A ^ memRead(getHL())
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
	return _or(regs.B)
}

func opORA_C() int {
	return _or(regs.C)
}

func opORA_D() int {
	return _or(regs.D)
}

func opORA_E() int {
	return _or(regs.E)
}

func opORA_H() int {
	return _or(regs.H)
}

func opORA_L() int {
	return _or(regs.L)
}

func opORA_A() int {
	return _or(regs.A)
}

func _or(v uint8) int {
	regs.A = regs.A | v
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
	return _or(memRead(getHL()))
}

// endregion

//region ORI data (OR immediate)

//   (A) <- (A) V (byte 2)
//   The content of the second byte of the instruction is inclusive-OR'd with the contents of the accumulator.
//   The result is placed in the accumulator. The CY and AC flags are cleared.
//   11100110
//   Cycles: 2  States: 7  Addressing: immediate  Flags: Z,S,P,CY,AC

func opORI() int {
	return _or(fetch())
}

//endregion

//region CMP r (Compare Register)

//   (A) - (r)
//   The content of register r is subtracted from the accumulator. The accumulator remains unchanged. The condition flags are set
//   as a result of the subtraction. The Z flag is set to 1 if (A) = (r). The CY flag is set to 1 if (A) < (r).
//   10111SSS
//   Cycles: 1  States: 4  Addressing: register  Flags: Z,S,P,CY,AC

func opCMP_B() int {
	sub(regs.A, regs.B)
	flags.CY = regs.A < regs.B
	return 1
}

func opCMP_C() int {
	sub(regs.A, regs.C)
	flags.CY = regs.A < regs.C
	return 1
}

func opCMP_D() int {
	sub(regs.A, regs.D)
	flags.CY = regs.A < regs.D
	return 1
}

func opCMP_E() int {
	sub(regs.A, regs.E)
	flags.CY = regs.A < regs.E
	return 1
}

func opCMP_H() int {
	sub(regs.A, regs.H)
	flags.CY = regs.A < regs.H
	return 1
}

func opCMP_L() int {
	sub(regs.A, regs.L)
	flags.CY = regs.A < regs.L
	return 1
}

func opCMP_A() int {
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
	v := memRead(getHL())
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
	sub_CY(regs.A, fetch())
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
	regs.PC = fetch16()
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
	return _j(!flags.Z)
}

func opJZ() int {
	return _j(flags.Z)
}

func opJNC() int {
	return _j(!flags.CY)
}

func opJC() int {
	return _j(flags.CY)
}

func opJPO() int {
	return _j(!flags.P)
}

func opJPE() int {
	return _j(flags.P)
}

func opJP() int {
	return _j(!flags.S)
}

func opJM() int {
	return _j(flags.S)
}

func _j(condition bool) int {
	if condition {
		regs.PC = fetch16()
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
	memWrite16(regs.SP-2, regs.PC+2)
	regs.SP -= 2
	regs.PC = fetch16()
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
	return _cp(!flags.Z)
}

func opCZ() int {
	return _cp(flags.Z)
}

func opCNC() int {
	return _cp(!flags.CY)
}

func opCC() int {
	return _cp(flags.CY)
}

func opCPO() int {
	return _cp(!flags.P)
}

func opCPE() int {
	return _cp(flags.P)
}

func opCP() int {
	return _cp(!flags.S)
}

func opCM() int {
	return _cp(flags.S)
}

func _cp(condition bool) int {
	if condition {
		memWrite16(regs.SP-2, regs.PC+2)
		regs.SP -= 2
		regs.PC = fetch16()
		return 5
	} else {
		regs.PC += 2
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
	return _r(!flags.Z)
}

func opRZ() int {
	return _r(flags.Z)
}

func opRNC() int {
	return _r(!flags.CY)
}

func opRC() int {
	return _r(flags.CY)
}

func opRPO() int {
	return _r(!flags.P)
}

func opRPE() int {
	return _r(flags.P)
}

func opRP() int {
	return _r(!flags.S)
}

func opRM() int {
	return _r(flags.S)
}

func _r(condition bool) int {
	if condition {
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
	return _rst(0x0000)
}

func opRST_1() int {
	return _rst(0x0008)
}

func opRST_2() int {
	return _rst(0x0010)
}

func opRST_3() int {
	return _rst(0x0018)
}

func opRST_4() int {
	return _rst(0x0020)
}

func opRST_5() int {
	return _rst(0x0028)
}

func opRST_6() int {
	return _rst(0x0030)
}

func opRST_7() int {
	return _rst(0x0038)
}

func _rst(addr uint16) int {
	memWrite16(regs.SP-2, regs.PC)
	regs.SP -= 2
	regs.PC = addr
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
	memWrite(regs.SP-1, regs.B)
	memWrite(regs.SP-2, regs.C)
	regs.SP -= 2
	return 3
}

func opPUSH_DE() int {
	memWrite(regs.SP-1, regs.D)
	memWrite(regs.SP-2, regs.E)
	regs.SP -= 2
	return 3
}

func opPUSH_HL() int {
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
	regs.C = memRead(regs.SP)
	regs.B = memRead(regs.SP + 1)
	regs.SP += 2
	return 3
}

func opPOP_DE() int {
	regs.E = memRead(regs.SP)
	regs.D = memRead(regs.SP + 1)
	regs.SP += 2
	return 3
}

func opPOP_HL() int {
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
	interrupt = true
	return 1
}

//endregion

//region DI (Disable interrupts)

//   The interrupt system is disabled immediately following the execution of the DI instruction.
//   1110011
//   Cycles: 1  States: 4  Flags: none

func opDI() int {
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
