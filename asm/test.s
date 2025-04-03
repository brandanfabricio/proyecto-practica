// Código ensamblador para Go en Windows (AMD64)
TEXT ·HolaMundo(SB), NOSPLIT, $0
    MOVQ $mensaje, CX     // Poner la dirección del mensaje en RCX (primer argumento de printf)
    CALL ·printf(SB)      // Llamar a printf
    RET                   // Retornar

DATA mensaje+0(SB)/13, $"Hola Mundo!\n"
GLOBL mensaje(SB), RODATA, $13
