package virloc

var (
	Commands map[CommandName]CommandMessage = map[CommandName]CommandMessage{
		TURN_OFF_OUT0: TURN_OFF_OUT0MSG,
		TURN_ON_OUT0:  TURN_ON_OUT0MSG,
		TURN_OFF_OUT1: TURN_OFF_OUT1MSG,
		TURN_ON_OUT1:  TURN_ON_OUT1MSG,
	}
)
