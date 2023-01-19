#include "textflag.h"
// func rdrand(ptr *uint64)
TEXT ·rdrand(SB),NOSPLIT,$0-8
    LEAQ ptr+0(FP), DX
    RDRAND AX
    MOVQ AX, DX
    RET

