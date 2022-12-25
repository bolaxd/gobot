package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func Join(ball *y.S, m *x.Parse)  {
	if !m.IsOwn {
		ball.Reply("fitur ini khusus owner", true)
		return
	}
	if m.Query == "" {
		ball.Reply("Masukan link group", true)
		return
	}
	ball.Joining(m.Query);
}