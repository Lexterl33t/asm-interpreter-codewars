package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type CPU_t struct {
	Registers map[string]int
	IP        int
}

func NewCPU() *CPU_t {
	return &CPU_t{
		Registers: make(map[string]int),
		IP:        0,
	}
}

func (cpu *CPU_t) mov(a, b string) {
	if unicode.IsDigit(rune([]byte(a)[0])) {
		panic("L'operande de destination doit être un registre")
	}

	if unicode.IsDigit(rune([]byte(b)[0])) {
		b_int, _ := strconv.Atoi(b)
		cpu.Registers[a] = b_int
	} else {
		cpu.Registers[a] = cpu.Registers[b]
	}

	cpu.IP++
}

func (cpu *CPU_t) inc(a string) {
	cpu.Registers[a]++
	cpu.IP++
}

func (cpu *CPU_t) dec(a string) {
	cpu.Registers[a]--
	cpu.IP++
}

func (cpu *CPU_t) jnz(a, offset string) {
	if unicode.IsDigit(rune([]byte(a)[0])) {
		panic("L'operande de destination doit être un registre")
	}

	if cpu.Registers[a] == 0 {
		cpu.IP++
		return
	}

	offset_int, _ := strconv.Atoi(offset)

	if offset_int >= 0 {
		cpu.IP = offset_int
	} else {
		cpu.IP += offset_int
	}

}

func asm_interpreter(program []string) map[string]int {
	cpu := NewCPU()

	size := len(program)

	for cpu.IP < size {
		opcode := strings.Split(program[cpu.IP], " ")

		switch opcode[0] {
		case "mov":
			cpu.mov(opcode[1], opcode[2])
			break
		case "inc":
			cpu.inc(opcode[1])
			break
		case "dec":
			cpu.dec(opcode[1])
			break
		case "jnz":
			cpu.jnz(opcode[1], opcode[2])
			break
		}

		fmt.Println(cpu.IP)
	}

	return cpu.Registers

}

func main() {

	code := []string{
		"mov a 5",
		"inc a",
		"dec a",
		"dec a",
		"jnz a -1",
		"inc a",
		"mov b a",
	}

	registers := asm_interpreter(code)

	fmt.Println(registers)

}
