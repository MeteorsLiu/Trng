
// func Uint64(ptr *uint64)
TEXT ·Uint64(SB),NOSPLIT,$0-8
    LEAQ ptr+0(FP), DX
    // RDRAND %RAX
    BYTE $0x0f; BYTE $0xc7; BYTE $0xf0
    MOVQ AX, DX
    RET

