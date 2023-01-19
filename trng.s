#include "textflag.h"
// func rdrand(ptr *uint64)
TEXT ·rdrand(SB),NOSPLIT,$0-8
    LEAQ ptr+0(FP), DX
    // RDRAND %RAX
    BYTE $0x0f; BYTE $0xc7; BYTE $0x0f
    MOVQ AX, DX
    RET

