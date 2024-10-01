// +build amd64
// Code generated by asm2asm, DO NOT EDIT.

package sse

var _text_lookup_small_key = []byte{
	// .p2align 4, 0x00
	// LCPI0_0
	0xbf, 0xbf, 0xbf, 0xbf, 0xbf, 0xbf, 0xbf, 0xbf, 0xbf, 0xbf, 0xbf, 0xbf, 0xbf, 0xbf, 0xbf, 0xbf, // QUAD $0xbfbfbfbfbfbfbfbf; QUAD $0xbfbfbfbfbfbfbfbf  // .space 16, '\xbf\xbf\xbf\xbf\xbf\xbf\xbf\xbf\xbf\xbf\xbf\xbf\xbf\xbf\xbf\xbf'
	//0x00000010 LCPI0_1
	0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, //0x00000010 QUAD $0x1919191919191919; QUAD $0x1919191919191919  // .space 16, '\x19\x19\x19\x19\x19\x19\x19\x19\x19\x19\x19\x19\x19\x19\x19\x19'
	//0x00000020 LCPI0_2
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, //0x00000020 QUAD $0x2020202020202020; QUAD $0x2020202020202020  // .space 16, '                '
	//0x00000030 .p2align 4, 0x90
	//0x00000030 _lookup_small_key
	0x55, //0x00000030 pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000031 movq         %rsp, %rbp
	0x41, 0x57, //0x00000034 pushq        %r15
	0x41, 0x56, //0x00000036 pushq        %r14
	0x41, 0x55, //0x00000038 pushq        %r13
	0x41, 0x54, //0x0000003a pushq        %r12
	0x53, //0x0000003c pushq        %rbx
	0x48, 0x83, 0xec, 0x28, //0x0000003d subq         $40, %rsp
	0x4c, 0x8b, 0x57, 0x08, //0x00000041 movq         $8(%rdi), %r10
	0x4c, 0x8b, 0x1e, //0x00000045 movq         (%rsi), %r11
	0x45, 0x0f, 0xb6, 0xc2, //0x00000048 movzbl       %r10b, %r8d
	0x4b, 0x8d, 0x0c, 0x80, //0x0000004c leaq         (%r8,%r8,4), %rcx
	0x45, 0x0f, 0xb6, 0x0c, 0x0b, //0x00000050 movzbl       (%r11,%rcx), %r9d
	0x48, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x00000055 movq         $-1, %rax
	0x45, 0x85, 0xc9, //0x0000005c testl        %r9d, %r9d
	0x0f, 0x84, 0x28, 0x03, 0x00, 0x00, //0x0000005f je           LBB0_39
	0x48, 0x89, 0x55, 0xb0, //0x00000065 movq         %rdx, $-80(%rbp)
	0x4c, 0x8b, 0x3f, //0x00000069 movq         (%rdi), %r15
	0x41, 0x8b, 0x44, 0x0b, 0x01, //0x0000006c movl         $1(%r11,%rcx), %eax
	0x48, 0x89, 0x45, 0xb8, //0x00000071 movq         %rax, $-72(%rbp)
	0x8d, 0xb0, 0xa5, 0x00, 0x00, 0x00, //0x00000075 leal         $165(%rax), %esi
	0x4c, 0x01, 0xde, //0x0000007b addq         %r11, %rsi
	0x41, 0x0f, 0xb6, 0xca, //0x0000007e movzbl       %r10b, %ecx
	0x41, 0x83, 0xf8, 0x09, //0x00000082 cmpl         $9, %r8d
	0x0f, 0x83, 0xd0, 0x00, 0x00, 0x00, //0x00000086 jae          LBB0_2
	0x45, 0x8a, 0x27, //0x0000008c movb         (%r15), %r12b
	0x45, 0x8d, 0x68, 0x01, //0x0000008f leal         $1(%r8), %r13d
	0x44, 0x89, 0xcb, //0x00000093 movl         %r9d, %ebx
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000096 .p2align 4, 0x90
	//0x000000a0 LBB0_7
	0x44, 0x38, 0x26, //0x000000a0 cmpb         %r12b, (%rsi)
	0x0f, 0x85, 0x97, 0x00, 0x00, 0x00, //0x000000a3 jne          LBB0_8
	0x44, 0x0f, 0xb6, 0x76, 0x01, //0x000000a9 movzbl       $1(%rsi), %r14d
	0xbf, 0x01, 0x00, 0x00, 0x00, //0x000000ae movl         $1, %edi
	0x45, 0x3a, 0x77, 0x01, //0x000000b3 cmpb         $1(%r15), %r14b
	0x0f, 0x85, 0x85, 0x00, 0x00, 0x00, //0x000000b7 jne          LBB0_16
	0x0f, 0xb6, 0x56, 0x02, //0x000000bd movzbl       $2(%rsi), %edx
	0xbf, 0x02, 0x00, 0x00, 0x00, //0x000000c1 movl         $2, %edi
	0x41, 0x3a, 0x57, 0x02, //0x000000c6 cmpb         $2(%r15), %dl
	0x0f, 0x85, 0x72, 0x00, 0x00, 0x00, //0x000000ca jne          LBB0_16
	0x0f, 0xb6, 0x56, 0x03, //0x000000d0 movzbl       $3(%rsi), %edx
	0xbf, 0x03, 0x00, 0x00, 0x00, //0x000000d4 movl         $3, %edi
	0x41, 0x3a, 0x57, 0x03, //0x000000d9 cmpb         $3(%r15), %dl
	0x0f, 0x85, 0x5f, 0x00, 0x00, 0x00, //0x000000dd jne          LBB0_16
	0x0f, 0xb6, 0x56, 0x04, //0x000000e3 movzbl       $4(%rsi), %edx
	0xbf, 0x04, 0x00, 0x00, 0x00, //0x000000e7 movl         $4, %edi
	0x41, 0x3a, 0x57, 0x04, //0x000000ec cmpb         $4(%r15), %dl
	0x0f, 0x85, 0x4c, 0x00, 0x00, 0x00, //0x000000f0 jne          LBB0_16
	0x0f, 0xb6, 0x56, 0x05, //0x000000f6 movzbl       $5(%rsi), %edx
	0xbf, 0x05, 0x00, 0x00, 0x00, //0x000000fa movl         $5, %edi
	0x41, 0x3a, 0x57, 0x05, //0x000000ff cmpb         $5(%r15), %dl
	0x0f, 0x85, 0x39, 0x00, 0x00, 0x00, //0x00000103 jne          LBB0_16
	0x0f, 0xb6, 0x56, 0x06, //0x00000109 movzbl       $6(%rsi), %edx
	0xbf, 0x06, 0x00, 0x00, 0x00, //0x0000010d movl         $6, %edi
	0x41, 0x3a, 0x57, 0x06, //0x00000112 cmpb         $6(%r15), %dl
	0x0f, 0x85, 0x26, 0x00, 0x00, 0x00, //0x00000116 jne          LBB0_16
	0x0f, 0xb6, 0x56, 0x07, //0x0000011c movzbl       $7(%rsi), %edx
	0x31, 0xff, //0x00000120 xorl         %edi, %edi
	0x41, 0x3a, 0x57, 0x07, //0x00000122 cmpb         $7(%r15), %dl
	0x40, 0x0f, 0x94, 0xc7, //0x00000126 sete         %dil
	0x48, 0x83, 0xc7, 0x07, //0x0000012a addq         $7, %rdi
	0xe9, 0x0f, 0x00, 0x00, 0x00, //0x0000012e jmp          LBB0_16
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000133 .p2align 4, 0x90
	//0x00000140 LBB0_8
	0x31, 0xff, //0x00000140 xorl         %edi, %edi
	//0x00000142 LBB0_16
	0x48, 0x39, 0xcf, //0x00000142 cmpq         %rcx, %rdi
	0x0f, 0x83, 0x91, 0x01, 0x00, 0x00, //0x00000145 jae          LBB0_17
	0x4c, 0x01, 0xee, //0x0000014b addq         %r13, %rsi
	0x83, 0xc3, 0xff, //0x0000014e addl         $-1, %ebx
	0x0f, 0x85, 0x49, 0xff, 0xff, 0xff, //0x00000151 jne          LBB0_7
	0xe9, 0x59, 0x00, 0x00, 0x00, //0x00000157 jmp          LBB0_20
	//0x0000015c LBB0_2
	0xf3, 0x41, 0x0f, 0x6f, 0x07, //0x0000015c movdqu       (%r15), %xmm0
	0xf3, 0x41, 0x0f, 0x6f, 0x4f, 0x10, //0x00000161 movdqu       $16(%r15), %xmm1
	0x48, 0xc7, 0xc7, 0xff, 0xff, 0xff, 0xff, //0x00000167 movq         $-1, %rdi
	0x48, 0xd3, 0xe7, //0x0000016e shlq         %cl, %rdi
	0x45, 0x8d, 0x60, 0x01, //0x00000171 leal         $1(%r8), %r12d
	0x44, 0x89, 0xcb, //0x00000175 movl         %r9d, %ebx
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000178 .p2align 4, 0x90
	//0x00000180 LBB0_3
	0xf3, 0x0f, 0x6f, 0x16, //0x00000180 movdqu       (%rsi), %xmm2
	0x66, 0x0f, 0x74, 0xd0, //0x00000184 pcmpeqb      %xmm0, %xmm2
	0xf3, 0x0f, 0x6f, 0x5e, 0x10, //0x00000188 movdqu       $16(%rsi), %xmm3
	0x66, 0x0f, 0x74, 0xd9, //0x0000018d pcmpeqb      %xmm1, %xmm3
	0x66, 0x0f, 0xd7, 0xd2, //0x00000191 pmovmskb     %xmm2, %edx
	0x66, 0x0f, 0xd7, 0xc3, //0x00000195 pmovmskb     %xmm3, %eax
	0xc1, 0xe0, 0x10, //0x00000199 shll         $16, %eax
	0x09, 0xd0, //0x0000019c orl          %edx, %eax
	0x09, 0xf8, //0x0000019e orl          %edi, %eax
	0x83, 0xf8, 0xff, //0x000001a0 cmpl         $-1, %eax
	0x0f, 0x84, 0x42, 0x01, 0x00, 0x00, //0x000001a3 je           LBB0_4
	0x4c, 0x01, 0xe6, //0x000001a9 addq         %r12, %rsi
	0x83, 0xc3, 0xff, //0x000001ac addl         $-1, %ebx
	0x0f, 0x85, 0xcb, 0xff, 0xff, 0xff, //0x000001af jne          LBB0_3
	//0x000001b5 LBB0_20
	0x48, 0x8b, 0x45, 0xb8, //0x000001b5 movq         $-72(%rbp), %rax
	0x48, 0x03, 0x45, 0xb0, //0x000001b9 addq         $-80(%rbp), %rax
	0x49, 0x01, 0xc3, //0x000001bd addq         %rax, %r11
	0xf3, 0x0f, 0x6f, 0x15, 0x38, 0xfe, 0xff, 0xff, //0x000001c0 movdqu       $-456(%rip), %xmm2  /* LCPI0_0+0(%rip) */
	0xf3, 0x41, 0x0f, 0x6f, 0x27, //0x000001c8 movdqu       (%r15), %xmm4
	0x66, 0x0f, 0x6f, 0xdc, //0x000001cd movdqa       %xmm4, %xmm3
	0x66, 0x0f, 0xfc, 0xda, //0x000001d1 paddb        %xmm2, %xmm3
	0xf3, 0x0f, 0x6f, 0x0d, 0x33, 0xfe, 0xff, 0xff, //0x000001d5 movdqu       $-461(%rip), %xmm1  /* LCPI0_1+0(%rip) */
	0x66, 0x0f, 0x6f, 0xc3, //0x000001dd movdqa       %xmm3, %xmm0
	0x66, 0x0f, 0xda, 0xc1, //0x000001e1 pminub       %xmm1, %xmm0
	0x66, 0x0f, 0x74, 0xc3, //0x000001e5 pcmpeqb      %xmm3, %xmm0
	0xf3, 0x0f, 0x6f, 0x1d, 0x2f, 0xfe, 0xff, 0xff, //0x000001e9 movdqu       $-465(%rip), %xmm3  /* LCPI0_2+0(%rip) */
	0x66, 0x0f, 0xdb, 0xc3, //0x000001f1 pand         %xmm3, %xmm0
	0x66, 0x0f, 0xfc, 0xc4, //0x000001f5 paddb        %xmm4, %xmm0
	0x41, 0x0f, 0xb6, 0xca, //0x000001f9 movzbl       %r10b, %ecx
	0x41, 0x83, 0xf8, 0x09, //0x000001fd cmpl         $9, %r8d
	0x0f, 0x83, 0xef, 0x00, 0x00, 0x00, //0x00000201 jae          LBB0_21
	0xf3, 0x0f, 0x7f, 0x45, 0xc0, //0x00000207 movdqu       %xmm0, $-64(%rbp)
	0x8a, 0x55, 0xc0, //0x0000020c movb         $-64(%rbp), %dl
	0x8a, 0x5d, 0xc1, //0x0000020f movb         $-63(%rbp), %bl
	0x44, 0x8a, 0x65, 0xc2, //0x00000212 movb         $-62(%rbp), %r12b
	0x44, 0x8a, 0x6d, 0xc3, //0x00000216 movb         $-61(%rbp), %r13b
	0x44, 0x8a, 0x55, 0xc4, //0x0000021a movb         $-60(%rbp), %r10b
	0x44, 0x8a, 0x75, 0xc5, //0x0000021e movb         $-59(%rbp), %r14b
	0x44, 0x8a, 0x7d, 0xc6, //0x00000222 movb         $-58(%rbp), %r15b
	0x8a, 0x45, 0xc7, //0x00000226 movb         $-57(%rbp), %al
	0x41, 0x83, 0xc0, 0x01, //0x00000229 addl         $1, %r8d
	0x41, 0x83, 0xf9, 0x02, //0x0000022d cmpl         $2, %r9d
	0xbf, 0x01, 0x00, 0x00, 0x00, //0x00000231 movl         $1, %edi
	0x41, 0x0f, 0x43, 0xf9, //0x00000236 cmovael      %r9d, %edi
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x0000023a .p2align 4, 0x90
	//0x00000240 LBB0_25
	0x41, 0x38, 0x13, //0x00000240 cmpb         %dl, (%r11)
	0x0f, 0x85, 0x77, 0x00, 0x00, 0x00, //0x00000243 jne          LBB0_26
	0xbe, 0x01, 0x00, 0x00, 0x00, //0x00000249 movl         $1, %esi
	0x41, 0x38, 0x5b, 0x01, //0x0000024e cmpb         %bl, $1(%r11)
	0x0f, 0x85, 0x6a, 0x00, 0x00, 0x00, //0x00000252 jne          LBB0_34
	0xbe, 0x02, 0x00, 0x00, 0x00, //0x00000258 movl         $2, %esi
	0x45, 0x38, 0x63, 0x02, //0x0000025d cmpb         %r12b, $2(%r11)
	0x0f, 0x85, 0x5b, 0x00, 0x00, 0x00, //0x00000261 jne          LBB0_34
	0xbe, 0x03, 0x00, 0x00, 0x00, //0x00000267 movl         $3, %esi
	0x45, 0x38, 0x6b, 0x03, //0x0000026c cmpb         %r13b, $3(%r11)
	0x0f, 0x85, 0x4c, 0x00, 0x00, 0x00, //0x00000270 jne          LBB0_34
	0xbe, 0x04, 0x00, 0x00, 0x00, //0x00000276 movl         $4, %esi
	0x45, 0x38, 0x53, 0x04, //0x0000027b cmpb         %r10b, $4(%r11)
	0x0f, 0x85, 0x3d, 0x00, 0x00, 0x00, //0x0000027f jne          LBB0_34
	0xbe, 0x05, 0x00, 0x00, 0x00, //0x00000285 movl         $5, %esi
	0x45, 0x38, 0x73, 0x05, //0x0000028a cmpb         %r14b, $5(%r11)
	0x0f, 0x85, 0x2e, 0x00, 0x00, 0x00, //0x0000028e jne          LBB0_34
	0xbe, 0x06, 0x00, 0x00, 0x00, //0x00000294 movl         $6, %esi
	0x45, 0x38, 0x7b, 0x06, //0x00000299 cmpb         %r15b, $6(%r11)
	0x0f, 0x85, 0x1f, 0x00, 0x00, 0x00, //0x0000029d jne          LBB0_34
	0x31, 0xf6, //0x000002a3 xorl         %esi, %esi
	0x41, 0x38, 0x43, 0x07, //0x000002a5 cmpb         %al, $7(%r11)
	0x40, 0x0f, 0x94, 0xc6, //0x000002a9 sete         %sil
	0x48, 0x83, 0xc6, 0x07, //0x000002ad addq         $7, %rsi
	0xe9, 0x0c, 0x00, 0x00, 0x00, //0x000002b1 jmp          LBB0_34
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x000002b6 .p2align 4, 0x90
	//0x000002c0 LBB0_26
	0x31, 0xf6, //0x000002c0 xorl         %esi, %esi
	//0x000002c2 LBB0_34
	0x48, 0x39, 0xce, //0x000002c2 cmpq         %rcx, %rsi
	0x0f, 0x83, 0xa8, 0x00, 0x00, 0x00, //0x000002c5 jae          LBB0_35
	0x4d, 0x01, 0xc3, //0x000002cb addq         %r8, %r11
	0x83, 0xc7, 0xff, //0x000002ce addl         $-1, %edi
	0x0f, 0x85, 0x69, 0xff, 0xff, 0xff, //0x000002d1 jne          LBB0_25
	0xe9, 0x8b, 0x00, 0x00, 0x00, //0x000002d7 jmp          LBB0_38
	//0x000002dc LBB0_17
	0x4c, 0x01, 0xee, //0x000002dc addq         %r13, %rsi
	0x48, 0x83, 0xc6, 0xff, //0x000002df addq         $-1, %rsi
	0x0f, 0xb6, 0x06, //0x000002e3 movzbl       (%rsi), %eax
	0xe9, 0xa2, 0x00, 0x00, 0x00, //0x000002e6 jmp          LBB0_39
	//0x000002eb LBB0_4
	0x48, 0x01, 0xce, //0x000002eb addq         %rcx, %rsi
	0x0f, 0xb6, 0x06, //0x000002ee movzbl       (%rsi), %eax
	0xe9, 0x97, 0x00, 0x00, 0x00, //0x000002f1 jmp          LBB0_39
	//0x000002f6 LBB0_21
	0xf3, 0x41, 0x0f, 0x6f, 0x67, 0x10, //0x000002f6 movdqu       $16(%r15), %xmm4
	0x66, 0x0f, 0xfc, 0xd4, //0x000002fc paddb        %xmm4, %xmm2
	0x66, 0x0f, 0xda, 0xca, //0x00000300 pminub       %xmm2, %xmm1
	0x66, 0x0f, 0x74, 0xca, //0x00000304 pcmpeqb      %xmm2, %xmm1
	0x66, 0x0f, 0xdb, 0xcb, //0x00000308 pand         %xmm3, %xmm1
	0x66, 0x0f, 0xfc, 0xcc, //0x0000030c paddb        %xmm4, %xmm1
	0x48, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x00000310 movq         $-1, %rax
	0x48, 0xd3, 0xe0, //0x00000317 shlq         %cl, %rax
	0x41, 0x83, 0xc0, 0x01, //0x0000031a addl         $1, %r8d
	0x41, 0x83, 0xf9, 0x02, //0x0000031e cmpl         $2, %r9d
	0xba, 0x01, 0x00, 0x00, 0x00, //0x00000322 movl         $1, %edx
	0x41, 0x0f, 0x43, 0xd1, //0x00000327 cmovael      %r9d, %edx
	0x90, 0x90, 0x90, 0x90, 0x90, //0x0000032b .p2align 4, 0x90
	//0x00000330 LBB0_22
	0xf3, 0x41, 0x0f, 0x6f, 0x13, //0x00000330 movdqu       (%r11), %xmm2
	0xf3, 0x41, 0x0f, 0x6f, 0x5b, 0x10, //0x00000335 movdqu       $16(%r11), %xmm3
	0x66, 0x0f, 0x74, 0xd9, //0x0000033b pcmpeqb      %xmm1, %xmm3
	0x66, 0x0f, 0x74, 0xd0, //0x0000033f pcmpeqb      %xmm0, %xmm2
	0x66, 0x0f, 0xd7, 0xf2, //0x00000343 pmovmskb     %xmm2, %esi
	0x66, 0x0f, 0xd7, 0xfb, //0x00000347 pmovmskb     %xmm3, %edi
	0xc1, 0xe7, 0x10, //0x0000034b shll         $16, %edi
	0x09, 0xf7, //0x0000034e orl          %esi, %edi
	0x09, 0xc7, //0x00000350 orl          %eax, %edi
	0x83, 0xff, 0xff, //0x00000352 cmpl         $-1, %edi
	0x0f, 0x84, 0x28, 0x00, 0x00, 0x00, //0x00000355 je           LBB0_23
	0x4d, 0x01, 0xc3, //0x0000035b addq         %r8, %r11
	0x83, 0xc2, 0xff, //0x0000035e addl         $-1, %edx
	0x0f, 0x85, 0xc9, 0xff, 0xff, 0xff, //0x00000361 jne          LBB0_22
	//0x00000367 LBB0_38
	0x48, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x00000367 movq         $-1, %rax
	0xe9, 0x1a, 0x00, 0x00, 0x00, //0x0000036e jmp          LBB0_39
	//0x00000373 LBB0_35
	0x4b, 0x8d, 0x34, 0x18, //0x00000373 leaq         (%r8,%r11), %rsi
	0x48, 0x83, 0xc6, 0xff, //0x00000377 addq         $-1, %rsi
	0x0f, 0xb6, 0x06, //0x0000037b movzbl       (%rsi), %eax
	0xe9, 0x0a, 0x00, 0x00, 0x00, //0x0000037e jmp          LBB0_39
	//0x00000383 LBB0_23
	0x49, 0x01, 0xcb, //0x00000383 addq         %rcx, %r11
	0x4c, 0x89, 0xde, //0x00000386 movq         %r11, %rsi
	0x41, 0x0f, 0xb6, 0x03, //0x00000389 movzbl       (%r11), %eax
	//0x0000038d LBB0_39
	0x48, 0x83, 0xc4, 0x28, //0x0000038d addq         $40, %rsp
	0x5b, //0x00000391 popq         %rbx
	0x41, 0x5c, //0x00000392 popq         %r12
	0x41, 0x5d, //0x00000394 popq         %r13
	0x41, 0x5e, //0x00000396 popq         %r14
	0x41, 0x5f, //0x00000398 popq         %r15
	0x5d, //0x0000039a popq         %rbp
	0xc3, //0x0000039b retq         
}
 
