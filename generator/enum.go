package generator

type FileType byte

const (
	Main FileType = iota
	Character
	Skill
	Event
)
