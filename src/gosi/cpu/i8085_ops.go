package cpu

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

func opMOV_B_B() {
	dasmOpReg1Reg2("MOV", "B", "B")
	regs.B = regs.B
}

func opMOV_B_C() {
	dasmOpReg1Reg2("MOV", "B", "C")
	regs.B = regs.C
}

func opMOV_B_D() {
	dasmOpReg1Reg2("MOV", "B", "D")
	regs.B = regs.D
}

func opMOV_B_E() {
	dasmOpReg1Reg2("MOV", "B", "E")
	regs.B = regs.E
}

func opMOV_B_H() {
	dasmOpReg1Reg2("MOV", "B", "H")
	regs.B = regs.H
}

func opMOV_B_L() {
	dasmOpReg1Reg2("MOV", "B", "L")
	regs.B = regs.L
}

func opMOV_C_B() {
	dasmOpReg1Reg2("MOV", "C", "B")
	regs.C = regs.B
}

func opMOV_C_C() {
	dasmOpReg1Reg2("MOV", "C", "C")
	regs.C = regs.C
}

func opMOV_C_D() {
	dasmOpReg1Reg2("MOV", "C", "D")
	regs.C = regs.D
}

func opMOV_C_E() {
	dasmOpReg1Reg2("MOV", "C", "E")
	regs.C = regs.E
}

func opMOV_C_H() {
	dasmOpReg1Reg2("MOV", "C", "H")
	regs.C = regs.H
}

func opMOV_C_L() {
	dasmOpReg1Reg2("MOV", "C", "L")
	regs.C = regs.L
}

func opMOV_D_B() {
	dasmOpReg1Reg2("MOV", "D", "B")
	regs.D = regs.B
}

func opMOV_D_C() {
	dasmOpReg1Reg2("MOV", "D", "C")
	regs.D = regs.C
}

func opMOV_D_D() {
	dasmOpReg1Reg2("MOV", "D", "D")
	regs.D = regs.D
}

func opMOV_D_E() {
	dasmOpReg1Reg2("MOV", "D", "E")
	regs.D = regs.E
}

func opMOV_D_H() {
	dasmOpReg1Reg2("MOV", "D", "H")
	regs.D = regs.H
}

func opMOV_D_L() {
	dasmOpReg1Reg2("MOV", "D", "L")
	regs.D = regs.L
}

func opMOV_E_B() {
	dasmOpReg1Reg2("MOV", "E", "B")
	regs.E = regs.B
}

func opMOV_E_C() {
	dasmOpReg1Reg2("MOV", "E", "C")
	regs.E = regs.C
}

func opMOV_E_D() {
	dasmOpReg1Reg2("MOV", "E", "D")
	regs.E = regs.D
}

func opMOV_E_E() {
	dasmOpReg1Reg2("MOV", "E", "E")
	regs.E = regs.E
}

func opMOV_E_H() {
	dasmOpReg1Reg2("MOV", "E", "H")
	regs.E = regs.H
}

func opMOV_E_L() {
	dasmOpReg1Reg2("MOV", "E", "L")
	regs.E = regs.L
}

func opMOV_H_B() {
	dasmOpReg1Reg2("MOV", "H", "B")
	regs.H = regs.B
}

func opMOV_H_C() {
	dasmOpReg1Reg2("MOV", "H", "C")
	regs.H = regs.C
}

func opMOV_H_D() {
	dasmOpReg1Reg2("MOV", "H", "D")
	regs.H = regs.D
}

func opMOV_H_E() {
	dasmOpReg1Reg2("MOV", "H", "E")
	regs.H = regs.E
}

func opMOV_H_H() {
	dasmOpReg1Reg2("MOV", "H", "H")
	regs.H = regs.H
}

func opMOV_H_L() {
	dasmOpReg1Reg2("MOV", "H", "L")
	regs.H = regs.L
}

func opMOV_L_B() {
	dasmOpReg1Reg2("MOV", "L", "B")
	regs.L = regs.B
}

func opMOV_L_C() {
	dasmOpReg1Reg2("MOV", "L", "C")
	regs.L = regs.C
}

func opMOV_L_D() {
	dasmOpReg1Reg2("MOV", "L", "D")
	regs.L = regs.D
}

func opMOV_L_E() {
	dasmOpReg1Reg2("MOV", "L", "E")
	regs.L = regs.E
}

func opMOV_L_H() {
	dasmOpReg1Reg2("MOV", "L", "H")
	regs.L = regs.H
}

func opMOV_L_L() {
	dasmOpReg1Reg2("MOV", "L", "L")
	regs.L = regs.L
}

// endregion

//region MOV r, M (Move from memory)

//   (r) <- ((H) (L))
//   The content of the memory location, whose address is in registers Hand L, is moved to register r.
//   01DDD110
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: none

func opMOV_B_M() {
	addr := hlGet()
	dasmOpRegAddr("MOV", "B", addr)
	regs.B = ramGet(addr)
}

func opMOV_C_M() {
	addr := hlGet()
	dasmOpRegAddr("MOV", "C", addr)
	regs.C = ramGet(addr)
}

func opMOV_D_M() {
	addr := hlGet()
	dasmOpRegAddr("MOV", "D", addr)
	regs.D = ramGet(addr)
}

func opMOV_E_M() {
	addr := hlGet()
	dasmOpRegAddr("MOV", "E", addr)
	regs.E = ramGet(addr)
}

func opMOV_H_M() {
	addr := hlGet()
	dasmOpRegAddr("MOV", "H", addr)
	regs.H = ramGet(addr)
}

func opMOV_L_M() {
	addr := hlGet()
	dasmOpRegAddr("MOV", "L", addr)
	regs.L = ramGet(addr)
}

func opMOV_A_M() {
	addr := hlGet()
	dasmOpRegAddr("MOV", "A", addr)
	regs.A = ramGet(addr)
}

//endregion

//region MOV M, r (Move to memory)

//   ((H) (L)) <- (r)
//   The content of register r is moved to the memory location whose address is in registers H and L.
//   01110SSS
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: none

func opMOV_M_B() {
	addr := hlGet()
	dasmOpAddrReg("MOV", addr, "B")
	ramSet(addr, regs.B)
}

func opMOV_M_C() {
	addr := hlGet()
	dasmOpAddrReg("MOV", addr, "C")
	ramSet(addr, regs.C)
}

func opMOV_M_D() {
	addr := hlGet()
	dasmOpAddrReg("MOV", addr, "D")
	ramSet(addr, regs.D)
}

func opMOV_M_E() {
	addr := hlGet()
	dasmOpAddrReg("MOV", addr, "E")
	ramSet(addr, regs.E)
}

func opMOV_M_H() {
	addr := hlGet()
	dasmOpAddrReg("MOV", addr, "H")
	ramSet(addr, regs.H)
}

func opMOV_M_L() {
	addr := hlGet()
	dasmOpAddrReg("MOV", addr, "L")
	ramSet(addr, regs.L)
}

func opMOV_M_A() {
	addr := hlGet()
	dasmOpAddrReg("MOV", addr, "A")
	ramSet(addr, regs.A)
}

//endregion

//region MVI r, data (Move Immediate)

//   (r) <- (byte 2)
//   The content of byte 2 of the instruction is moved to register r.
//   00DDD110
//   Cycles: 2  States: 7  Addressing: immediate  Flags: none

func opMVI_B_data() {
	r := rom[regs.PC+1]
	dasmOpRegVal("MVI", "B", r)
	regs.B = r
	regs.PC++
}

func opMVI_C_data() {
	r := rom[regs.PC+1]
	dasmOpRegVal("MVI", "C", r)
	regs.C = r
	regs.PC++
}

func opMVI_D_data() {
	r := rom[regs.PC+1]
	dasmOpRegVal("MVI", "D", r)
	regs.D = r
	regs.PC++
}

func opMVI_E_data() {
	r := rom[regs.PC+1]
	dasmOpRegVal("MVI", "E", r)
	regs.E = r
	regs.PC++
}

func opMVI_H_data() {
	r := rom[regs.PC+1]
	dasmOpRegVal("MVI", "H", r)
	regs.H = r
	regs.PC++
}

func opMVI_L_data() {
	r := rom[regs.PC+1]
	dasmOpRegVal("MVI", "L", r)
	regs.L = r
	regs.PC++
}

//endregion

//region MVI M, data (Move to memory immediate)

//   ((H) (L)) <- (byte 2)
//   The content of byte 2 of the instruction is moved to the memory location whose address is in registers H and L.
//   00110110
//   Cycles: 3  States: 10  Addressing: immed./reg. indirect  Flags: none

func opMVI_M_data() {
	r := rom[regs.PC+1]
	addr := hlGet()
	dasmOpAddrVal("MVI", addr, r)
	ramSet(addr, r)
	regs.PC++
}

//endregion

//region LXI rp, data 16 (Load register pair immediate)

//   (rh) <- (byte 3)
//   (rl) <- (byte 2)
//   Byte 3 of the instruction is moved into the high-order register (rh) of the register pair rp.
//   Byte 2 of the instruction is moved into the low-order register (rl) of the register pair rp.
//   00RP0001
//   Cycles: 3  States: 10  Addressing: immediate  Flags: none

func opLXI_BC_data16() {
	rh := rom[regs.PC+2]
	rl := rom[regs.PC+1]
	dasmOpRegVal1Val2("LXI", "BC", rh, rl)
	regs.B = rh
	regs.C = rl
	regs.PC += 2
}

func opLXI_DE_data16() {
	rh := rom[regs.PC+2]
	rl := rom[regs.PC+1]
	dasmOpRegVal1Val2("LXI", "DE", rh, rl)
	regs.D = rh
	regs.E = rl
	regs.PC += 2
}

func opLXI_HL_data16() {
	rh := rom[regs.PC+2]
	rl := rom[regs.PC+1]
	dasmOpRegVal1Val2("LXI", "HL", rh, rl)
	regs.H = rh
	regs.L = rl
	regs.PC += 2
}

func opLXI_SP_data16() {
	rp := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpRegAddr("LXI", "SP", rp)
	regs.SP = rp
	regs.PC += 2
}

//endregion

//region LDAX rp (Load accumulator indirect)

//   (A) <- ((rp))
//   The content of the memory location, whose address is in the register pair rp, is moved to register A.
//   00RP1010
//   Cycles: 2  States: 7  Addressing: reg. indirect  Flags: none

func opLDAX_BC() {
	dasmOpReg("LDAX", "BC")
	addr := bcGet()
	regs.A = ramGet(addr)
}

func opLDAX_DE() {
	dasmOpReg("LDAX", "DE")
	addr := deGet()
	regs.A = ramGet(addr)
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

func opINR_B() {
	dasmOpReg("INR", "B")
	regs.B = add(regs.B, 1)
}

func opINR_C() {
	dasmOpReg("INR", "C")
	regs.C = add(regs.C, 1)
}

func opINR_D() {
	dasmOpReg("INR", "D")
	regs.D = add(regs.D, 1)
}

func opINR_E() {
	dasmOpReg("INR", "E")
	regs.E = add(regs.E, 1)
}

func opINR_H() {
	dasmOpReg("INR", "H")
	regs.H = add(regs.H, 1)
}

func opINR_L() {
	dasmOpReg("INR", "L")
	regs.L = add(regs.L, 1)
}

//endregion

//region DCR r (Decrement Register)

//   (r) <- (r) - 1
//   The content of register r is decremented by one.
//   00DDD101
//   Cycles: 1  States: 5  Addressing: register  Flags: Z,S,P,AC

func opDCR_B() {
	dasmOpReg("DCR", "B")
	regs.B = sub(regs.B, 1)
}

func opDCR_C() {
	dasmOpReg("DCR", "C")
	regs.C = sub(regs.C, 1)
}

func opDCR_D() {
	dasmOpReg("DCR", "D")
	regs.D = sub(regs.D, 1)
}

func opDCR_E() {
	dasmOpReg("DCR", "E")
	regs.E = sub(regs.E, 1)
}

func opDCR_H() {
	dasmOpReg("DCR", "H")
	regs.H = sub(regs.H, 1)
}

func opDCR_L() {
	dasmOpReg("DCR", "L")
	regs.L = sub(regs.L, 1)
}

//endregion

//region INX rp (Increment register pair)

//   (rh) (rl) <- (rh) (rl) + 1
//   The content of the register pair rp is incremented by one. Note: No condition flags are affected.
//   00RP0011
//   Cycles: 1  States: 5  Addressing: register  Flags: none

func opINX_BC() {
	dasmOpReg("INX", "BC")
	addr := bcGet()
	addr++
	bcSet(addr)
}

func opINX_DE() {
	dasmOpReg("INX", "DE")
	addr := deGet()
	addr++
	deSet(addr)
}

func opINX_HL() {
	dasmOpReg("INX", "HL")
	addr := hlGet()
	addr++
	hlSet(addr)
}

func opINX_SP() {
	dasmOpReg("INX", "SP")
	regs.SP++
}

//endregion

//region DCX rp (Decrement register pair)

//   (rh) (rl) <- (rh) (rl) - 1
//   The content of the register pair rp is decremented by one. Note: No condition flags are affected.
//   00RP1011
//   Cycles: 1  States: 5  Addressing: register  Flags: none

func opDCX_BC() {
	dasmOpReg("DCX", "BC")
	addr := bcGet()
	addr--
	bcSet(addr)
}

func opDCX_DE() {
	dasmOpReg("DCX", "DE")
	addr := deGet()
	addr--
	deSet(addr)
}

func opDCX_HL() {
	dasmOpReg("DCX", "HL")
	addr := hlGet()
	addr--
	hlSet(addr)
}

func opDCX_SP() {
	dasmOpReg("DCX", "SP")
	regs.SP--
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

func opJMP() {
	addr := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpAddr("JMP", addr)
	regs.PC = addr
	regs.PC-- // fixme
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

func opCALL() {
	addr := binary.LittleEndian.Uint16(rom[regs.PC+1:])
	dasmOpAddr("CALL", addr)
	regs.SP = regs.PC
	regs.SP -= 2
	regs.PC = addr
	regs.PC-- // fixme
}

//endregion

// ********** Stack, I/O, and Machine Control Group **********
// This group of instructions performs I/O, manipulates the Stack, and alters internal control flags.
// Unless otherwise specified, condition flags are not affected by any instructions in this group.

//region NOP (No op)

//   No operation is performed. The registers and flags are unaffected.
//   00000000
//   Cycles: 1  States: 4  Flags: none

func opNOP() {
	dasmOp("NOP")
}

//endregion
