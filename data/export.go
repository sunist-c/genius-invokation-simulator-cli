package data

type ExportedSkillDescription struct {
	DescriptionType  string `yaml:"description_type"`
	SkillName        string `yaml:"skill_name"`
	SkillDescription string `yaml:"skill_description"`
}

type ExportedCharacterDescription struct {
	DescriptionType      string `yaml:"description_type"`
	CharacterName        string `yaml:"character_name"`
	CharacterDescription string `yaml:"character_description"`
	CharacterTitle       string `yaml:"character_title"`
	CharacterStory       string `yaml:"character_story"`
}

type ExportedModDescription struct {
	DescriptionType string `yaml:"description_type"`
	ModName         string `yaml:"mod_name"`
	ModDescription  string `yaml:"mod_description"`
}

type ExportedSkillContext struct {
	ShortID             uint16                              `yaml:"short_id"`
	OwnCharacter        string                              `yaml:"own_character"`
	OwnCharacterShortID uint16                              `yaml:"own_character_short_id"`
	SkillName           string                              `yaml:"skill_name"`
	SkillType           string                              `yaml:"skill_type"`
	Descriptions        map[string]ExportedSkillDescription `yaml:"descriptions"`
}

type ExportedCharacterContext struct {
	ShortID              uint16                                  `yaml:"short_id"`
	CharacterName        string                                  `yaml:"character_name"`
	CharacterAffiliation string                                  `yaml:"character_affiliation"`
	CharacterVision      string                                  `yaml:"character_vision"`
	CharacterWeapon      string                                  `yaml:"character_weapon"`
	CharacterHP          uint                                    `yaml:"character_hp"`
	CharacterMP          uint                                    `yaml:"character_mp"`
	CharacterSkills      []uint16                                `yaml:"character_skills"`
	Descriptions         map[string]ExportedCharacterDescription `yaml:"descriptions"`
}

type ExportedModContext struct {
	ModID        uint64                            `yaml:"mod_id"`
	PackagePath  string                            `yaml:"package_path"`
	Descriptions map[string]ExportedModDescription `yaml:"descriptions"`
}

type ExportedEventContext struct{}

type ExportedCardContext struct{}

type ExportedSummonContext struct{}

type ExportedModifierContext struct{}

type ExportedContext struct {
	Mod        ExportedModContext
	Paths      map[string]string
	Characters map[uint64]ExportedCharacterContext
	Skills     map[uint64]ExportedSkillContext
	Events     map[uint64]ExportedEventContext
	Cards      map[uint64]ExportedCardContext
	Summons    map[uint64]ExportedSummonContext
	Modifiers  map[uint64]ExportedModifierContext
}
