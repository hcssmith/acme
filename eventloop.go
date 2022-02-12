package acme

import (
	"strings"
	"strconv"
)

const (
	E OriginTypes = iota // Writes body / tags
	F // Actions  on other files / windows
	K // Keyboard
	M // Mouse
)

const (
	D ActionTypes = iota // deleted (body)
	d // deleted (tags)
	I // Inserted (body)
	i // inserted (tags)
	L // button 3 (body)
	l // button 3 (tags)
	X // button 2 (body)
	x // button 2 (tags)
)

type OriginTypes int64
type ActionTypes int64

type Event struct {
	Origin OriginTypes // orgin see const
	Action ActionTypes // action see const
	CAddr1 int // start address of action
	CAddr2 int // end address of action
	Flag int // flag
	NumChars int // count of chars in text
	Text string // Optional Text
}

type Arg struct {
	Loc string
	Arg string
}

type EvCommand struct {
	C string
	A Arg
	W *Win
}

type Evm map[string]func(*Win, Arg)

func parse_event(arr []byte) (ev Event) {
	switch arr[0] {
		case 'E':
			ev.Origin = E
		case 'F':
			ev.Origin = F
		case 'K':
			ev.Origin = K
		case 'M':
			ev.Origin = M
	}
	switch arr[1] {
		case 'D':
			ev.Action = D
		case 'd':
			ev.Action = d
		case 'I':
			ev.Action = I
		case 'i':
			ev.Action = i
		case 'L':
			ev.Action = L
		case 'l':
			ev.Action = l
		case 'X':
			ev.Action = X
		case 'x':
			ev.Action = x
	}
	arr2 := string(arr[2:])
	spl := strings.Split(arr2, " ")
	ev.CAddr1, _ = strconv.Atoi(spl[0])
	ev.CAddr2, _ = strconv.Atoi(spl[1])
	ev.Flag, _ = strconv.Atoi(spl[2])
	ev.NumChars, _ = strconv.Atoi(spl[3])
	ev.Text = strings.Join(spl[4:], " ")
	return
}

// Window event loop
func (self *Win)Event_Loop(mp Evm) {
	c, _ := NewCommand(R, self.Event)
	r, _ := ReadCommand(c)

	for {
		var ev1, ev2, ev3, ev4 Event
		var evnum int
		var c EvCommand
		c.W = self
		line, _, _ := r.ReadLine()

		ev1 = parse_event(line)
		evnum = 1

		if ev1.Flag&2 != 0 {
			line2, _, _ := r.ReadLine()
			ev2 = parse_event(line2)
			evnum = 2
		}

		if ev1.Flag&8 != 0 {
			line3, _, _ := r.ReadLine()
			ev3 = parse_event(line3)
			line4, _, _ := r.ReadLine()
			ev4 = parse_event(line4)
			evnum = 4
		}

		if evnum == 2 {
			c.C = ev2.Text
			c.A.Arg = ""
			c.A.Loc = ""
		}

		if evnum == 4 {
			c.C = ev2.Text
			c.A.Arg = ev3.Text
			c.A.Loc = ev4.Text
		}

		for comm, fun := range mp {
			if comm == c.C && (ev1.Action == X || ev1.Action == x){
				fun(c.W, c.A)
			}
		}
	}
}