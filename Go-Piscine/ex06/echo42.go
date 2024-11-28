package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var noNewline bool
	var separator string
	var helpFlag bool

	flag.BoolVar(&noNewline, "n", false, "最後の改行を省略")
	flag.StringVar(&separator, "s", " ", "-sの次に渡される文字列の内容で出力引数を区切る。(デフォルトは空白)")
	flag.BoolVar(&helpFlag, "help", false, "helpを表示する")
	flag.Parse()
	if helpFlag {
		fmt.Println("Usage of ./echo42:")
		fmt.Println("  -n    omit trailing newline")
		fmt.Println("  -s string")
		fmt.Println("        separator (default \" \")")
		return
	}
	args := flag.Args()
	output := strings.Join(args, separator)
	if noNewline {
		fmt.Print(output)
	} else {
		fmt.Println(output)
	}
}
