package plugin 

import (
	"os/exec"
	"fmt"
	x "gobot/utils/message"
	y "gobot/utils/simple"
	a "gobot/constanta"
	)
	
func Execute(ball *y.S, m *x.Parse)  {
	if !m.IsOwn { ball.Reply(a.FOwner, true); return }
	res, err := exec.Command("bash", "-c", m.Query).Output()
	if err != nil {
		ball.Reply(fmt.Sprintf("%v", err), true)
		return
	}
	ball.Reply(string(res), true)
}