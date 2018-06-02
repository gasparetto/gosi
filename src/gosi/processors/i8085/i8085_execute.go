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
	add(regs.B)
	return 1
}

func opADD_C() int {
	add(regs.C)
	return 1
}

func opADD_D() int {
	add(regs.D)
	return 1
}

func opADD_E() int {
	add(regs.E)
	return 1
}

func opADD_H() int {
	add(regs.H)
	return 1
}

func opADD_L() int {
	add(regs.L)
	return 1
}

func opADD_A() int {
	add(regs.A)
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
	add(memRead(getHL()))
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
	add(fetch())
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
	adc(regs.B)
	return 1
}

func opADC_C() int {
	adc(regs.C)
	return 1
}

func opADC_D() int {
	adc(regs.D)
	return 1
}

func opADC_E() int {
	adc(regs.E)
	return 1
}

func opADC_H() int {
	adc(regs.H)
	return 1
}

func opADC_L() int {
	adc(regs.L)
	return 1
}

func opADC_A() int {
	adc(regs.A)
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
	adc(memRead(getHL()))
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
	adc(fetch())
	return 2
}

//endregion

//region SUB r (Subtract Register)

//   (A) <- (A) - r
//   The content of register r is subtracted from the content of the accumulator. The result is placed in the accumulator.
//   10010SSS
//   Cycles: 1  States: 4  Addressing: register  Flags: Z,S,P,CY,AC

func opSUB_B() int {
	sub(regs.B)
	return 1
}

func opSUB_C() int {
	sub(regs.C)
	return 1
}

func opSUB_D() int {
	sub(regs.D)
	return 1
}

func opSUB_E() int {
	sub(regs.E)
	return 1
}

func opSUB_H() int {
	sub(regs.H)
	return 1
}

func opSUB_L() int {
	sub(regs.L)
	return 1
}

func opSUB_A() int {
	sub(regs.A)
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
	sub(memRead(getHL()))
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
	sub(fetch())
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
	sbb(regs.B)
	return 1
}

func opSBB_C() int {
	sbb(regs.C)
	return 1
}

func opSBB_D() int {
	sbb(regs.D)
	return 1
}

func opSBB_E() int {
	sbb(regs.E)
	return 1
}

func opSBB_H() int {
	sbb(regs.H)
	return 1
}

func opSBB_L() int {
	sbb(regs.L)
	return 1
}

func opSBB_A() int {
	sbb(regs.A)
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
	sbb(memRead(getHL()))
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
	sbb(fetch())
	return 2
}

//endregion

//region INR r (Increment Register)

//   (r) <- (r) + 1
//   The content of register r is incremented by one. Note: All condition flags except CY are affected.
//   00DDD100
//   Cycles: 1  States: 5  Addressing: register  Flags: Z,S,P,AC

func opINR_B() int {
	regs.B = inr(regs.B)
	return 1
}

func opINR_C() int {
	regs.C = inr(regs.C)
	return 1
}

func opINR_D() int {
	regs.D = inr(regs.D)
	return 1
}

func opINR_E() int {
	regs.E = inr(regs.E)
	return 1
}

func opINR_H() int {
	regs.H = inr(regs.H)
	return 1
}

func opINR_L() int {
	regs.L = inr(regs.L)
	return 1
}

func opINR_A() int {
	regs.A = inr(regs.A)
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
	memWrite(addr, inr(memRead(addr)))
	return 3
}

//endregion

//region DCR r (Decrement Register)

//   (r) <- (r) - 1
//   The content of register r is decremented by one. Note: All condition except CY are affected.
//   00DDD101
//   Cycles: 1  States: 5  Addressing: register  Flags: Z,S,P,AC

func opDCR_B() int {
	regs.B = dcr(regs.B)
	return 1
}

func opDCR_C() int {
	regs.C = dcr(regs.C)
	return 1
}

func opDCR_D() int {
	regs.D = dcr(regs.D)
	return 1
}

func opDCR_E() int {
	regs.E = dcr(regs.E)
	return 1
}

func opDCR_H() int {
	regs.H = dcr(regs.H)
	return 1
}

func opDCR_L() int {
	regs.L = dcr(regs.L)
	return 1
}

func opDCR_A() int {
	regs.A = dcr(regs.A)
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
	memWrite(addr, dcr(memRead(addr)))
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
	dad(getBC())
	return 3
}

func opDAD_DE() int {
	dad(getDE())
	return 3
}

func opDAD_HL() int {
	dad(getHL())
	return 3
}

func opDAD_SP() int {
	dad(regs.SP)
	return 3
}

//endregion

//region DAA (Decimal Adjust Accumulator)

//   The eight-bit number in the accumulator is adjusted to form two four-bit Binary-Coded-Decimal digits by the following
//   process:
//   1. If the value of the least significant 4 bits of the accumulator is greater than 9 or if the AC flag is set, 6 is added
//      to the accumulator.
//   2. If the value of the most significant 4 bits of the accumulator is now greater than 9, or if the CY flag is set, 6 is added
//      to the most significant 4 bits of the accumulator.
//   NOTE: All flags are affected.
//   00100111
//   Cycles: 1  States: 4  Flags: Z,S,P,CY,AC

func opDAA() int {
	res16 := uint16(regs.A)
	if (res16 & 0x0f) > 0x09 || flags.AC {
		res16 += 0x06
	}
	if (res16 & 0xf0) > 0x90 || flags.CY {
		res16 += 0x60
	}
	regs.A = uint8(res16)
	flags_Z_S_P(regs.A)
	flags.CY = res16 > 0x00ff
	flags.AC = res16 > 0x000f
	return 1
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
	ana(regs.B)
	return 1
}

func opANA_C() int {
	ana(regs.C)
	return 1
}

func opANA_D() int {
	ana(regs.D)
	return 1
}

func opANA_E() int {
	ana(regs.E)
	return 1
}

func opANA_H() int {
	ana(regs.H)
	return 1
}

func opANA_L() int {
	ana(regs.L)
	return 1
}

func opANA_A() int {
	ana(regs.A)
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
	ana(memRead(getHL()))
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
	ana(fetch())
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
	xra(regs.B)
	return 1
}

func opXRA_C() int {
	xra(regs.C)
	return 1
}

func opXRA_D() int {
	xra(regs.D)
	return 1
}

func opXRA_E() int {
	xra(regs.E)
	return 1
}

func opXRA_H() int {
	xra(regs.H)
	return 1
}

func opXRA_L() int {
	xra(regs.L)
	return 1
}

func opXRA_A() int {
	xra(regs.A)
	return 1
}

//endregion

//region XRA M (Exclusive OR Memory)

//   (A) <- (A) Y ((H) (L))
//   The content of the memory location whose address is contained in the H and L registers is exclusive-OR'd with the content of
//   the accumulator. The result is placed in the accumulator. The CY and AC flags are cleared.
//   10101110
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: Z,S,P,CY,AC

func opXRA_M() int {
	xra(memRead(getHL()))
	return 2
}

// endregion

//region XRI data (Exclusive OR immediate)

//   (A) <- (A) Y (byte 2)
//   The content of the second byte of the instruction is exclusive-OR'd with the content of the accumulator. The result is placed
//   in the accumulator. The CY and AC flags are cleared.
//   11101110
//   Cycles: 2  States: 7  Addressing: immediate  Flags: Z,S,P,CY,AC

func opXRI() int {
	xra(fetch())
	return 2
}

// endregion

//region ORA r (OR Register)

//   (A) <- (A) V (r)
//   The content of register r is inclusive-or'd with the content of the accumulator. The result is placed in the accumulator.
//   The CY and AC flags are cleared.
//   10110SSS
//   Cycles: 1  States: 4  Addressing: register  Flags: Z,S,P,CY,AC

func opORA_B() int {
	ora(regs.B)
	return 1
}

func opORA_C() int {
	ora(regs.C)
	return 1
}

func opORA_D() int {
	ora(regs.D)
	return 1
}

func opORA_E() int {
	ora(regs.E)
	return 1
}

func opORA_H() int {
	ora(regs.H)
	return 1
}

func opORA_L() int {
	ora(regs.L)
	return 1
}

func opORA_A() int {
	ora(regs.A)
	return 1
}

//endregion

//region ORA M (OR Memory)

//   (A) <- (A) V ((H) (L))
//   The content of the memory location whose address is contained in the H and L registers is inclusive-OR'd with the content of
//   the accumulator. The result is placed in the accumulator. The CY and AC flags are cleared.
//   10110110
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: Z,S,P,CY,AC

func opORA_M() int {
	ora(memRead(getHL()))
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
	ora(fetch())
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
	cmp(regs.B)
	return 1
}

func opCMP_C() int {
	cmp(regs.C)
	return 1
}

func opCMP_D() int {
	cmp(regs.D)
	return 1
}

func opCMP_E() int {
	cmp(regs.E)
	return 1
}

func opCMP_H() int {
	cmp(regs.H)
	return 1
}

func opCMP_L() int {
	cmp(regs.L)
	return 1
}

func opCMP_A() int {
	cmp(regs.A)
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
	cmp(memRead(getHL()))
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
	cmp(fetch())
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
	regs.A = regs.A<<1 | a0>>7
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
	regs.A = regs.A>>1 | a0<<7
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
	regs.A = regs.A<<1 | btoi(flags.CY)
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
	regs.A = regs.A>>1 | btoi(flags.CY)<<7
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

//region CMC (Complement carry)

//   (CY) <- (!CY)
//   The CY flag is complemented. No other flags are affected.
//   00111111
//   Cycles: 1  States: 4  Flags: CY

func opCMC() int {
	flags.CY = !flags.CY
	return 1
}

//endregion

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

//region SPHL (Move HL to SP)

//   (SP) <- (H) (L)
//   The contents of registers H and L (16 bits) are moved to register SP.
//   11111001
//   Cycles: 1  States: 5  Addressing: register  Flags: none

func opSPHL() int {
	regs.SP = getHL()
	return 1
}

//endregion

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
//   11110011
//   Cycles: 1  States: 4  Flags: none

func opDI() int {
	interrupt = false
	return 1
}

//endregion

//region HLT (Halt)

//   The processor is stopped. The registers and flags are unaffected.
//   01110110
//   Cycles: 1  States: 7  Flags: none

func opHLT() int {
	_logFatal("Halt")
	return 1
}

//endregion

//region NOP (No op)

//   No operation is performed. The registers and flags are unaffected.
//   00000000
//   Cycles: 1  States: 4  Flags: none

func opNOP() int {
	return 1
}

//endregion

// ********** Utils **********

func inr(v uint8) uint8 {
	v++
	flags_Z_S_P(v)
	flags.AC = (v & 0x0f) == 0
	return v
}

func add(v uint8) {
	res16 := uint16(regs.A) + uint16(v)
	index := uint16((regs.A & 0x88) >> 1) | uint16((v & 0x88) >> 2) | ((res16 & 0x88) >> 3)
	regs.A = uint8(res16)
	flags_Z_S_P(regs.A)
	flags.AC = half_carry_table[index & 0x7] == 1
	flags.CY = (res16 & 0x0100) != 0
}

func adc(v uint8) {
	res16 := uint16(regs.A) + uint16(v) + uint16(btoi(flags.CY))
	index := uint16((regs.A & 0x88) >> 1) | uint16((v & 0x88) >> 2) | ((res16 & 0x88) >> 3)
	regs.A = uint8(res16)
	flags_Z_S_P(regs.A)
	flags.AC = half_carry_table[index & 0x7] == 1
	flags.CY = (res16 & 0x0100) != 0
}

func dcr(v uint8) uint8 {
	v--
	flags_Z_S_P(v)
	flags.AC = (v & 0x0f) != 0x0f
	return v
}

func sub(v uint8) {
	res16 := uint16(regs.A) - uint16(v)
	index := uint16((regs.A & 0x88) >> 1) | uint16((v & 0x88) >> 2) | ((res16 & 0x88) >> 3)
	regs.A = uint8(res16)
	flags_Z_S_P(regs.A)
	flags.AC = sub_half_carry_table[index & 0x7] == 0
	flags.CY = (res16 & 0x0100) != 0
}

func sbb(v uint8) {
	res16 := uint16(regs.A) - uint16(v) - uint16(btoi(flags.CY))
	index := uint16((regs.A & 0x88) >> 1) | uint16((v & 0x88) >> 2) | ((res16 & 0x88) >> 3)
	regs.A = uint8(res16)
	flags_Z_S_P(regs.A)
	flags.AC = sub_half_carry_table[index & 0x7] == 0
	flags.CY = (res16 & 0x0100) != 0
}

func cmp(v uint8) {
	res16 := uint16(regs.A) - uint16(v)
	index := uint16((regs.A & 0x88) >> 1) | uint16((v & 0x88) >> 2) | ((res16 & 0x88) >> 3)
	res8 := uint8(res16)
	flags_Z_S_P(res8)
	flags.AC = sub_half_carry_table[index & 0x7] == 0
	flags.CY = (res16 & 0x0100) != 0
}

func ana(v uint8) {
	flags.AC = ((regs.A | v) & 0x08) != 0
	regs.A &= v
	flags_Z_S_P(regs.A)
	flags.CY = false
}

func xra(v uint8) {
	regs.A ^= v
	flags_Z_S_P(regs.A)
	flags.CY = false
	flags.AC = false
}

func ora(v uint8) {
	regs.A |= v
	flags_Z_S_P(regs.A)
	flags.CY = false
	flags.AC = false
}

func dad(v uint16) {
	r := uint32(getHL()) + uint32(v)
	setHL(uint16(r))
	flags.CY = (r & 0x10000) != 0
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

var half_carry_table = []int { 0, 0, 1, 0, 1, 0, 1, 1 }
var sub_half_carry_table = []int { 0, 1, 1, 1, 0, 0, 0, 1 }
