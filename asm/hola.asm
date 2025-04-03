section .data
    msg db "Hola Mundo!", 0

section .text
    global holaMundo
    extern printf

holaMundo:
    sub rsp, 32
    mov rcx, msg
    call printf
    add rsp, 32
    ret
