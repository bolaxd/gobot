package plugin 

import (
	"os"
	"os/exec"
	"fmt"
	x "gobot/utils/message"
	y "gobot/utils/simple"
	"gobot/utils/helper"
	a "gobot/constanta"
	)
	
func Update(ball *y.S, m *x.Parse)  {
	if !m.IsOwn { ball.Reply(a.FOwner, true); return }
	email := os.Getenv("EMAIL_GH")
	username := os.Getenv("USERNAME_GH")
	caption := helper.GenerateID()
	res, err := exec.Command("bash", "-c", `git config --global user.email ` +email+ ` && git config --global user.name `+username+` && git add . && git commit -m `+caption+` && git push`).Output()
	if err != nil {
		ball.Reply(fmt.Sprintf("%v", err), true)
		return
	}
	ball.Reply(string(res), true)
}