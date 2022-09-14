package main

type Peep struct {
	message string `json: "message"`
	uuid    string `json: "uuid"`
	date    string `json: "date"`
	user    string `json: "user"`
}

func newPeep(m string, d string, u string, id string) Peep {
	return Peep{message: m, date: d, user: u, uuid: id}
}

func (p Peep) get() string {
	return p.message + ", posted by: " + p.user + ", on date: " + p.date
}

func (p Peep) checkUuid(uuid string) bool {
	if p.uuid == uuid {
		return true
	}; return false
}