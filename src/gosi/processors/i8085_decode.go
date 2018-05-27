package processors

import "log"

func decode(op uint8) func() int {

	switch op {
	case 0x00:
		return opNOP
	case 0x01:
		return opLXI_BC
	case 0x02:
		return opSTAX_BC
	case 0x03:
		return opINX_BC
	case 0x04:
		return opINR_B
	case 0x05:
		return opDCR_B
	case 0x06:
		return opMVI_B
	case 0x07:
		return opRLC
	case 0x08:
		//undocumented NOP
	case 0x09:
		return opDAD_BC
	case 0x0a:
		return opLDAX_BC
	case 0x0b:
		return opDCX_BC
	case 0x0c:
		return opINR_C
	case 0x0d:
		return opDCR_C
	case 0x0e:
		return opMVI_C
	case 0x0f:
		return opRRC
	case 0x10:
		//undocumented NOP
	case 0x11:
		return opLXI_DE
	case 0x12:
		return opSTAX_DE
	case 0x13:
		return opINX_DE
	case 0x14:
		return opINR_D
	case 0x15:
		return opDCR_D
	case 0x16:
		return opMVI_D
	case 0x17:
		//fixme RAL
	case 0x18:
		//undocumented NOP
	case 0x19:
		return opDAD_DE
	case 0x1a:
		return opLDAX_DE
	case 0x1b:
		return opDCX_DE
	case 0x1c:
		return opINR_E
	case 0x1d:
		return opDCR_E
	case 0x1e:
		return opMVI_E
	case 0x20:
		//undocumented NOP
	case 0x21:
		return opLXI_HL
	case 0x22:
		return opSHLD
	case 0x23:
		return opINX_HL
	case 0x24:
		return opINR_H
	case 0x25:
		return opDCR_H
	case 0x26:
		return opMVI_H
	case 0x27:
		//fixme DAA
	case 0x28:
		//undocumented NOP
	case 0x29:
		return opDAD_HL
	case 0x2a:
		return opLHLD
	case 0x2b:
		return opDCX_HL
	case 0x2c:
		return opINR_L
	case 0x2d:
		return opDCR_L
	case 0x2e:
		return opMVI_L
	case 0x2f:
		//fixme CMA
	case 0x30:
		//undocumented NOP
	case 0x31:
		return opLXI_SP
	case 0x32:
		return opSTA
	case 0x33:
		return opINX_SP
	case 0x34:
		return opINR_M
	case 0x35:
		return opDCR_M
	case 0x36:
		return opMVI_M
	case 0x37:
		return opSTC
	case 0x38:
		//undocumented NOP
	case 0x39:
		return opDAD_SP
	case 0x3a:
		return opLDA
	case 0x3b:
		return opDCX_SP
	case 0x3c:
		return opINR_A
	case 0x3d:
		return opDCR_A
	case 0x3e:
		return opMVI_A
	case 0x3f:
		//fixme CMC
	case 0x40:
		return opMOV_B_B
	case 0x41:
		return opMOV_C_B
	case 0x42:
		return opMOV_D_B
	case 0x43:
		return opMOV_E_B
	case 0x44:
		return opMOV_H_B
	case 0x45:
		return opMOV_L_B
	case 0x46:
		return opMOV_B_M
	case 0x47:
		return opMOV_B_A
	case 0x48:
		return opMOV_B_C
	case 0x49:
		return opMOV_C_C
	case 0x4a:
		return opMOV_D_C
	case 0x4b:
		return opMOV_E_C
	case 0x4c:
		return opMOV_H_C
	case 0x4d:
		return opMOV_L_C
	case 0x4e:
		return opMOV_C_M
	case 0x4f:
		return opMOV_C_A
	case 0x50:
		return opMOV_B_D
	case 0x51:
		return opMOV_C_D
	case 0x52:
		return opMOV_D_D
	case 0x53:
		return opMOV_E_D
	case 0x54:
		return opMOV_H_D
	case 0x55:
		return opMOV_L_D
	case 0x56:
		return opMOV_D_M
	case 0x57:
		return opMOV_D_A
	case 0x58:
		return opMOV_B_E
	case 0x59:
		return opMOV_C_E
	case 0x5a:
		return opMOV_D_E
	case 0x5b:
		return opMOV_E_E
	case 0x5c:
		return opMOV_H_E
	case 0x5d:
		return opMOV_L_E
	case 0x5e:
		return opMOV_E_M
	case 0x5f:
		return opMOV_E_A
	case 0x60:
		return opMOV_B_H
	case 0x61:
		return opMOV_C_H
	case 0x62:
		return opMOV_D_H
	case 0x63:
		return opMOV_E_H
	case 0x64:
		return opMOV_H_H
	case 0x65:
		return opMOV_L_H
	case 0x66:
		return opMOV_H_M
	case 0x67:
		return opMOV_H_A
	case 0x68:
		return opMOV_B_L
	case 0x69:
		return opMOV_C_L
	case 0x6a:
		return opMOV_D_L
	case 0x6b:
		return opMOV_E_L
	case 0x6c:
		return opMOV_H_L
	case 0x6d:
		return opMOV_L_L
	case 0x6e:
		return opMOV_L_M
	case 0x6f:
		return opMOV_L_A
	case 0x70:
		return opMOV_M_B
	case 0x71:
		return opMOV_M_C
	case 0x72:
		return opMOV_M_D
	case 0x73:
		return opMOV_M_E
	case 0x74:
		return opMOV_M_H
	case 0x75:
		return opMOV_M_L
	case 0x76:
		//fixme HLT
	case 0x77:
		return opMOV_M_A
	case 0x78:
		return opMOV_A_B
	case 0x79:
		return opMOV_A_C
	case 0x7a:
		return opMOV_A_D
	case 0x7b:
		return opMOV_A_E
	case 0x7c:
		return opMOV_A_H
	case 0x7d:
		return opMOV_A_L
	case 0x7e:
		return opMOV_A_M
	case 0x7f:
		return opMOV_A_A
	case 0x80:
		return opADD_B
	case 0x81:
		return opADD_C
	case 0x82:
		return opADD_D
	case 0x83:
		return opADD_E
	case 0x84:
		return opADD_H
	case 0x85:
		return opADD_L
	case 0x86:
		return opADD_M
	case 0x87:
		return opADD_A
	case 0x88:
		return opADC_B
	case 0x89:
		return opADC_C
	case 0x8a:
		return opADC_D
	case 0x8b:
		return opADC_E
	case 0x8c:
		return opADC_H
	case 0x8d:
		return opADC_L
	case 0x8e:
		return opADC_M
	case 0x8f:
		return opADC_A
	case 0x90:
		return opSUB_B
	case 0x91:
		return opSUB_C
	case 0x92:
		return opSUB_D
	case 0x93:
		return opSUB_E
	case 0x94:
		return opSUB_H
	case 0x95:
		return opSUB_L
	case 0x96:
		return opSUB_M
	case 0x97:
		return opSUB_A
	case 0x98:
		return opSBB_B
	case 0x99:
		return opSBB_C
	case 0x9a:
		return opSBB_D
	case 0x9b:
		return opSBB_E
	case 0x9c:
		return opSBB_H
	case 0x9d:
		return opSBB_L
	case 0x9e:
		return opSBB_M
	case 0x9f:
		return opSBB_A
	case 0xa0:
		return opANA_B
	case 0xa1:
		return opANA_C
	case 0xa2:
		return opANA_D
	case 0xa3:
		return opANA_E
	case 0xa4:
		return opANA_H
	case 0xa5:
		return opANA_L
	case 0xa6:
		return opANA_M
	case 0xa7:
		return opANA_A
	case 0xa8:
		return opXRA_B
	case 0xa9:
		return opXRA_C
	case 0xaa:
		return opXRA_D
	case 0xab:
		return opXRA_E
	case 0xac:
		return opXRA_H
	case 0xad:
		return opXRA_L
	case 0xae:
		return opXRA_M
	case 0xaf:
		return opXRA_A
	case 0xb0:
		return opORA_B
	case 0xb1:
		return opORA_C
	case 0xb2:
		return opORA_D
	case 0xb3:
		return opORA_E
	case 0xb4:
		return opORA_H
	case 0xb5:
		return opORA_L
	case 0xb6:
		return opORA_M
	case 0xb7:
		return opORA_A
	case 0xc0:
		return opRNZ
	case 0xc1:
		return opPOP_BC
	case 0xc2:
		return opJNZ
	case 0xc3:
		return opJMP
	case 0xc4:
		return opCNZ
	case 0xc5:
		return opPUSH_BC
	case 0xc6:
		return opADI
	case 0xc7:
		return opRST_0
	case 0xc8:
		return opRZ
	case 0xc9:
		return opRET
	case 0xca:
		return opJZ
	case 0xcb:
		//undocumented JMP nnnn
	case 0xcc:
		return opCZ
	case 0xcd:
		return opCALL
	case 0xce:
		return opACI
	case 0xcf:
		return opRST_1
	case 0xd0:
		return opRNC
	case 0xd1:
		return opPOP_DE
	case 0xd2:
		return opJNC
	case 0xd3:
		return opOUT
	case 0xd4:
		return opCNC
	case 0xd5:
		return opPUSH_DE
	case 0xd6:
		return opSUI
	case 0xd7:
		return opRST_2
	case 0xd8:
		return opRC
	case 0xd9:
		//undocumented RET
	case 0xda:
		return opJC
	case 0xdb:
		return opIN
	case 0xdc:
		return opCC
	case 0xde:
		return opSBI
	case 0xdf:
		return opRST_3
	case 0xe0:
		return opRPO
	case 0xe1:
		return opPOP_HL
	case 0xe2:
		return opJPO
	case 0xe3:
		return opXTHL
	case 0xe4:
		return opCPO
	case 0xe5:
		return opPUSH_HL
	case 0xe6:
		return opANI
	case 0xe7:
		return opRST_4
	case 0xe8:
		return opRPE
	case 0xe9:
		return opPCHL
	case 0xea:
		return opJPE
	case 0xeb:
		return opXCHG
	case 0xec:
		return opCPE
	case 0xed:
		//undocumented CALL nnnn
	case 0xee:
		//fixme XRI nn
	case 0xef:
		return opRST_5
	case 0xf0:
		return opRP
	case 0xf1:
		return opPOP_PSW
	case 0xf2:
		return opJP
	case 0xf3:
		return opDI
	case 0xf4:
		return opCP
	case 0xf5:
		return opPUSH_PSW
	case 0xf6:
		//fixme ORI nn
	case 0xf7:
		return opRST_6
	case 0xf8:
		return opRM
	case 0xf9:
		//fixme SPHL
	case 0xfa:
		return opJM
	case 0xfb:
		return opEI
	case 0xfc:
		return opCM
	case 0xfd:
		//fixme undocumented CALL nnnn
	case 0xfe:
		return opCPI
	case 0xff:
		return opRST_7
	}

	trace()
	log.Fatalf("ERROR: Unknown opcode 0x%02x\n", op)
	return nil
}
