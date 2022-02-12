package acme

import (
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

// Get list of current windows (populates self.Windows)
func (self *Acme) Get_acme_windows() {
	self.Windows = nil
	cmd, _ := NewCommand(R, self.Index)
	index, _ := ReadCommand(cmd)
	for {
		line, _, err := index.ReadLine()
		if err != nil {
			break
		}
		arr := []byte(line)
		ID, _ := strconv.Atoi(strings.TrimSpace(string(arr[:11])))
		w := NewWin(ID)
		self.Windows = append(self.Windows, w)
	}
}

// create a new Acme window, returns the ID of the new window
func (self *Acme) Create_win(title string) int {
	cmd, _ := NewCommand(R, self.Log)
	b, _ := ReadCommand(cmd)
	ic := make(chan int)
	tc := make(chan int)
	go func (){
		var bt bool
		bt = false
		for {
			line, _, _ := b.ReadLine()
			if bt == false {
				tc<-1
				bt = true
			}
			entry := strings.Split(string(line), " ")
			if len(entry) > 2 {
				if entry[1] == "new" {
					ID, _ := strconv.Atoi(entry[0])
					ic <- ID
					break
				}
			}
		}
	}()

	for {
		t := <-tc
		if t != 1 {
			continue
		}
		break
	}
	pl := fmt.Sprintf("name %s\n", title)
	cmd2, _ := NewCommand(W, self.New + "/ctl")
	_ = SendCommand(cmd2, pl)
	i := <-ic
	w := NewWin(i)
	self.Windows = append(self.Windows, w)
	return i
}

// Clears all text from a given window
func (self *Win)Clear_Window() {
	self.Set_Addr(",")
	c2, _ := NewCommand(W, self.Data)
	_ = SendCommand(c2, " ")
}


// Clear window then set text
func (self *Win)Set_Text(text string) {
	self.Clear_Window()
	c, _ := NewCommand(W, self.Body)
	_ = SendCommand(c, text)
}

// Appened text
func (self *Win)Append_Text(text string) {
	c, _ := NewCommand(W, self.Body)
	_ = SendCommand(c, text)
}

// Append Tags
func (self *Win)Append_Tags(tags string) {
	c, _ := NewCommand(W, " " + self.Tag)
	_ = SendCommand(c, tags)
}

// Set Tags
func (self *Win)Set_Tags(tags string) {
	self.Send_Ctl("cleartag")
	c, _ := NewCommand(W, self.Tag)
	_ = SendCommand(c, " " +tags)
}

func (self *Win)Set_Addr(addr string) {
	c, _ := NewCommand(W, self.Addr)
	_ = SendCommand(c, addr)
}

func (self *Win)Send_Ctl(command string) {
	c, _ := NewCommand(W, self.Ctl)
	_ = SendCommand(c, command)
}

func (self *Acme)Main_Log() (r *bufio.Reader) {
	c, _ := NewCommand(R, self.Log)
	r, _ = ReadCommand(c)
	return
}

func (self *Win)Event_Log() (r *bufio.Reader) {
	c, _ := NewCommand(R, self.Event)
	r, _ = ReadCommand(c)
	return
}

func (self *Acme)Get_Index() (r *bufio.Reader) {
	c, _ := NewCommand(R, self.Index)
	r, _ = ReadCommand(c)
	return r
}