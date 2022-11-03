package pwc_arg

import "time"

type PwcArg struct {
	ClientAddr string
	StartTime  time.Time
	EndTime    time.Time
}

func (PwcArg pa) AddClient(string) bool {

}

func (PwcArg pa) ExistClient(string) bool {

}
