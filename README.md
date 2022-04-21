# Disasm

Proof of concept linear disassembly built in Go. Support to both 32-bit and
64-bit Windows binaries.

Example output:

```
C:>.\disam.exe C:\Windows\System32\calc.exe
0x140001010: push rbx
0x140001012: sub rsp, 0x40
0x140001016: mov r10, qword ptr [rsp+0x80]
0x14000101e: mov rbx, rcx
0x140001021: test r10, r10
0x140001024: jz 0x9e
0x140001026: mov r11, qword ptr [rsp+0x70]
0x14000102b: mov eax, edx
0x14000102d: test edx, edx
0x14000102f: jz 0x65
[...]
```

Compared to the same disassembly generated by [Binary Ninja](https://binary.ninja/):

![Disassembly](binary_ninja_disass.png)
