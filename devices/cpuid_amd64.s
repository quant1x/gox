//+build amd64,!gccgo,!noasm,!appengine

#include "textflag.h"

// func asmCpuid(op uint32) (eax, ebx, ecx, edx uint32)
TEXT ·asmCpuid(SB), NOSPLIT, $0-24
	XORQ CX, CX
	MOVL op+0(FP), AX
	CPUID
	MOVL AX, eax+8(FP)
	MOVL BX, ebx+12(FP)
	MOVL CX, ecx+16(FP)
	MOVL DX, edx+20(FP)
	RET

// func asmCpuidex(op, op2 uint32) (eax, ebx, ecx, edx uint32)
TEXT ·asmCpuidex(SB), NOSPLIT, $0-24
	MOVL op+0(FP), AX
	MOVL op2+4(FP), CX
	CPUID
	MOVL AX, eax+8(FP)
	MOVL BX, ebx+12(FP)
	MOVL CX, ecx+16(FP)
	MOVL DX, edx+20(FP)
	RET
