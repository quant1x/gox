package progressbar

import (
	"fmt"
	"gitee.com/quant1x/gox/api"
	"strings"
	"sync"
)

var (
	mu           sync.Mutex
	gSrcLine     = 0 //起点行
	gCurrentLine = 0 //当前行
	gMaxLine     = 0 //最大行
)

func barMove(line int) {
	//fmt.Println("\n\n\n\n", gCurrentLine, line)
	fmt.Printf("\033[%dA\033[%dB", gCurrentLine, line)
	gCurrentLine = line
}

func barPrint(line int, args ...any) {
	mu.Lock()
	defer mu.Unlock()

	barMove(line)
	var realArgs []any
	realArgs = append(realArgs, "\r")
	realArgs = append(realArgs, args...)
	fmt.Print(realArgs...)
	barMove(gMaxLine)
}

func barPrintf(line int, format string, args ...any) {
	mu.Lock()
	defer mu.Unlock()

	barMove(line)
	fmt.Printf("\r"+format, args...)
	barMove(gMaxLine)
}

func barPrintln(line int, args ...any) {
	mu.Lock()
	defer mu.Unlock()

	barMove(line)
	var realArgs []any
	realArgs = append(realArgs, "\r")
	realArgs = append(realArgs, args...)
	fmt.Print(realArgs...)
	barMove(gMaxLine)
}

func Print(args ...any) {
	mu.Lock()
	lf := countLF("", args...)
	if gMaxLine == 0 {
		gMaxLine += lf + 1
	} else {
		gMaxLine += lf
	}
	mu.Unlock()

	barPrint(gMaxLine, args...)
}

func Printf(format string, args ...any) {
	mu.Lock()

	lf := countLF(format, args...)
	if gMaxLine == 0 {
		gMaxLine += lf + 1
	} else {
		gMaxLine += lf
	}
	mu.Unlock()

	barPrintf(gMaxLine, format, args...)
}

func Println(args ...any) {
	mu.Lock()

	lf := countLF("", args...)
	lf++
	if gMaxLine == 0 {
		gMaxLine += lf + 1
	} else {
		gMaxLine += lf
	}
	mu.Unlock()

	barPrintln(gMaxLine, args...)
}

func countLF(format string, args ...any) int {
	var count int
	count = strings.Count(format, "\n")
	for _, arg := range args {
		count += strings.Count(api.ToString(arg), "\n")
	}

	return count
}
