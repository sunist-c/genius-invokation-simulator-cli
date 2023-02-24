package parser

type Command string

const (
	CommandNewEntity      Command = "new"
	CommandModifyEntity   Command = "modify"
	CommandDeleteEntity   Command = "delete"
	CommandAddDescription Command = "description"
)

type EntityType string

const (
	EntityCharacter EntityType = "character"
	EntitySkill     EntityType = "skill"
	EntityEvent     EntityType = "event"
	EntitySummon    EntityType = "summon"
	EntityCard      EntityType = "card"
	EntityRule      EntityType = "rule"
)
