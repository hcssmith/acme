package main

import (
	. "github.com/hcssmith/acme"
	"os"
)


func main() {

	fm := Evm{
		"test1":test1,
		"Close":Close,
		"textarg":textarg,
	}

	a := NewAcme()

	text := "This is some text\nand some furter text\n\tand a bit more"

	i := a.Create_win("Notification")
	w, _ := a.GetWindow(i)

	w.Set_Tags("Refresh Close textarg")
	w.Set_Text(text)



	w.Event_Loop(fm)

}

func test1(w *Win, a Arg) {
	w.Set_Text("How about this instead")
}

func textarg(w *Win, a Arg) {
	w.Set_Text(a.Arg)
}

func Close(w *Win, a Arg) {
	w.Send_Ctl("clean")
	os.Exit(0)
}