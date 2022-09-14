package main

type Peep struct {
	message string
	uuid    string
	date    string
	user    string
}

func newPeep(m string, d string, u string) Peep {
	return Peep{message: m, date: d, user: u}
}

func (p Peep) get() string {
	return p.message + ", posted by: " + p.user + ", on date: " + p.date
}
