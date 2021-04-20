package halting

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

var ErrPartialResult = errors.New("partial result")
var ErrRepairFailed = errors.New("repair faild")

type Instruction struct {
	op  string
	arg int
}

type Bootloader struct {
	code    []*Instruction
	repairs map[int]string
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
	for line, op := range b.repairs {
		ins := b.code[line]
		b.code[line] = &Instruction{op: op, arg: ins.arg}

		if acc, err := b.Run(); err == nil {
			return acc, nil
		}

		b.code[line] = ins
	}

	return -1, ErrRepairFailed
}

func (b *Bootloader) Run() (int, error) {
	done := map[int]struct{}{}
	ip, acc := 0, 0
	for {
		if _, run := done[ip]; run {
			return acc, ErrPartialResult
		}

		if ip == len(b.code) {
			return acc, nil
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

	boot := &Bootloader{repairs: make(map[int]string)}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		arg, err := strconv.Atoi(line[1])
		if err != nil {
			return boot, len(boot.code), err
		}

		ins := &Instruction{op: line[0], arg: arg}
		boot.code = append(boot.code, ins)
		if ins.op == "nop" {
			boot.repairs[len(boot.code)-1] = "jmp"
		} else if ins.op == "jmp" {
			boot.repairs[len(boot.code)-1] = "nop"
		}
	}

	return boot, len(boot.code), nil
}
