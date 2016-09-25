package server

func (p *ClientHandler) HandleUser() {
	p.user = p.param
	p.writeMessage(331, "User name ok, password required")
}

func (p *ClientHandler) HandlePass() {
	// think about using https://developer.bitium.com
	if p.daddy.driver.CheckUser(p.user, p.param, &p.userInfo) {
		p.writeMessage(230, "Password ok, continue")
	} else {
		p.writeMessage(530, "Incorrect password, not logged in")
		p.conn.Close()
		delete(p.daddy.ConnectionMap, p.cid)
	}
}
