package system

import "time"

const (
	UpdateInterval = time.Second * 5
)

type System struct {
	State State
	Exit  chan bool
}

func NewSystem() System {
	return System{
		State: NewState(),
		Exit:  make(chan bool),
	}
}

func Process(sys *System) {

	ticker := time.NewTicker(UpdateInterval)
	defer ticker.Stop()

	for {

		select {

		case <-sys.Exit:
			return

		case <-ticker.C:
			//...

		}

	}

}
