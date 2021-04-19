package halting

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Execution int

const (
	Incomplete Execution = iota
	Complete
)

type Instruction struct {
	op  string
	arg int
}

type Bootloader struct {
	code    []*Instruction
	repairs []int
}

func (i *Instruction) TryRepair() *Instruction {
	if i.op == "nop" {
		return &Instruction{arg: i.arg, op: "jmp"}
	}

	return &Instruction{arg: i.arg, op: "nop"}
}

func (i *Instruction) Execute(ip *int, acc *int) {
	switch {
	case i.op == "jmp":
		*ip += i.arg
	case i.op == "acc":
		*acc += i.arg
		fallthrough
	default:
		*ip += 1
	}
}

func (b *Bootloader) Repair() (int, error) {
	for _, i := range b.repairs {
		ins := b.code[i]
		b.code[i] = ins.TryRepair()

		if result, acc := b.Run(); result == Complete {
			return acc, nil
		}

		b.code[i] = ins
	}

	return -1, errors.New("program corrupt, no fix found")
}

func (b *Bootloader) Run() (Execution, int) {
	done := map[int]struct{}{}
	ip, acc := 0, 0
	for {
		if _, run := done[ip]; run {
			return Incomplete, acc
		}

		if ip == len(b.code) {
			return Complete, acc
		}

		done[ip] = struct{}{}
		b.code[ip].Execute(&ip, &acc)
	}
}

func Load(path string) (program *Bootloader, lines int, err error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer fp.Close()

	boot := &Bootloader{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		arg, err := strconv.Atoi(line[1])
		if err != nil {
			return boot, len(boot.code), err
		}

		ins := &Instruction{op: line[0], arg: arg}
		boot.code = append(boot.code, ins)
		if ins.op == "nop" || ins.op == "jmp" {
			boot.repairs = append(boot.repairs, len(boot.code)-1)
		}
	}

	return boot, len(boot.code), nil
}
