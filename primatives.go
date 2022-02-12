package acme

import (
	"errors"
	"fmt"
)

type RWL string

const (
	R RWL = "read"
	W RWL = "write"
)

type Command struct {
	BaseCommand string
	Rwl RWL
	Path string
}

func NewCommand(Rwl RWL, Path string) (comm Command, e error) {
	switch Rwl {
		case R, W:
			comm.Rwl = Rwl
		default:
			return comm, errors.New("Invalid RWL")
		}
	comm.BaseCommand = "9p"
	comm.Path = Path
	return comm, nil
}


type Acme struct {
	Consctl string
	Index string
	Log string
	New string
	Windows []Win
}

func (self *Acme)GetWindow(i int) (w Win, e error) {
	for _, value := range self.Windows {
		if value.WinID == i {
			return value, nil
		}
	}
	return w, errors.New("Window not found")
}

func NewAcme() (a Acme) {
	a.Consctl = "acme/consctl"
	a.Index = "acme/index"
	a.Log = "acme/log"
	a.New = "acme/new"
	return
}


type Win struct {
	WinID int
	Addr string
	Body string
	Ctl string
	Data string
	Errors string
	Event string
	Tag string
}

func NewWin(ID int) (w Win) {
	w.WinID = ID
	w.Addr = fmt.Sprintf("acme/%d/addr", ID)
	w.Body = fmt.Sprintf("acme/%d/body", ID)
	w.Ctl = fmt.Sprintf("acme/%d/ctl", ID)
	w.Data = fmt.Sprintf("acme/%d/data", ID)
	w.Errors = fmt.Sprintf("acme/%d/errors", ID)
	w.Event = fmt.Sprintf("acme/%d/event", ID)
	w.Tag = fmt.Sprintf("acme/%d/tag", ID)
	return
}