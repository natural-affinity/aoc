package halting

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	op  string
	arg int
}

type Bootloader struct {
	Acc  int
	code []*Instruction
}

func (b *Bootloader) Lines() int {
	return len(b.code)
}

func (b *Bootloader) Execute(i *Instruction) (jmp int) {
	switch {
	case i.op == "acc":
		b.Acc += i.arg
	case i.op == "jmp":
		return i.arg
	}

	return 1
}

func (b *Bootloader) RunOnce() {
	done := map[int]struct{}{}
	ip := 0
	for {
		if _, run := done[ip]; run {
			return
		}

		jmp := b.Execute(b.code[ip])
		done[ip] = struct{}{}
		ip += jmp
	}
}

func Load(path string) (*Bootloader, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	boot := &Bootloader{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		arg, err := strconv.Atoi(line[1])
		if err != nil {
			return boot, err
		}

		ins := &Instruction{op: line[0], arg: arg}
		boot.code = append(boot.code, ins)
	}

	return boot, nil
}
