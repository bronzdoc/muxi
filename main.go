package main

import (
	"fmt"
	"os"

	"github.com/bronzdoc/muxi/tmux"
)

func main() {
	layout := tmux.NewLayout("./test.yml")
	err := layout.Create()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("success")

	//s := tmux.NewSession("test-doc2")

	//w := tmux.NewWindow("")
	//s.AddWindow(w)

	//w1 := tmux.NewWindow("")
	//s.AddWindow(w1)

	//p := tmux.NewPane()
	//p.AddCommand("ls")
	//p2 := tmux.NewPane()
	//p.AddCommand("zshrc")
	//p3 := tmux.NewPane()
	////p2 := tmux.NewPane()
	//w.AddPane(p)
	//w.AddPane(p2)
	//w.AddPane(p3)
	////w.AddPane(p2)

	//s.Create()

	//cmd := exec.Command("tmux", "select-window", "-t", "2")
	//cmd.Stdin = os.Stdin
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr

	//out, err := cmd.Output()

	//if err != nil {
	//	fmt.Println(err)
	//}

	//if err := cmd.Start(); err != nil {
	//	fmt.Println(err)
	//}

	//if err := cmd.Wait(); err != nil {
	//	fmt.Println(err)
	//}

	//fmt.Println(string(out))
}
