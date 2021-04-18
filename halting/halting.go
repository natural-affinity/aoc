package halting

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	op  string
	arg string
}

type Bootloader struct {
	code []*Instruction
	acc  int
}

func (b *Bootloader) Lines() int {
	return len(b.code)
}

func (b *Bootloader) Execute(i *Instruction) (jump int, err error) {
	arg, err := strconv.Atoi(i.arg)
	if err != nil {
		return 0, err
	}

	switch {
	case i.op == "acc":
		b.acc += arg
	case i.op == "jmp":
		return arg, nil
	}

	return 1, nil
}

func (b *Bootloader) RunOnce() (int, error) {
	done := map[int]struct{}{}
	ip := 0
	for {
		if _, run := done[ip]; run {
			return b.acc, nil
		}

		jump, err := b.Execute(b.code[ip])
		if err != nil {
			return -1, err
		}

		done[ip] = struct{}{}
		ip += jump
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
		ins := &Instruction{op: line[0], arg: line[1]}
		boot.code = append(boot.code, ins)
	}

	return boot, nil
}
