package cli

type Command = string

const (
	CommandInitMod        Command = "init"
	CommandNewEntity      Command = "new"
	CommandModifyEntity   Command = "modify"
	CommandDeleteEntity   Command = "delete"
	CommandAddDescription Command = "description"
	CommandList           Command = "list"
	CommandExit           Command = "exit"
)

var (
	AllCommand = []Command{
		CommandInitMod, CommandNewEntity, CommandModifyEntity, CommandDeleteEntity, CommandAddDescription, CommandList,
		CommandExit,
	}
)

type EntityType = string

const (
	EntityCharacter EntityType = "character"
	EntitySkill     EntityType = "skill"
	EntityEvent     EntityType = "event"
	EntitySummon    EntityType = "summon"
	EntityCard      EntityType = "card"
	EntityRule      EntityType = "rule"
)

var (
	AllEntityType = []EntityType{
		EntityCharacter, EntitySkill, EntityEvent, EntitySummon, EntityCard, EntityRule,
	}
)
