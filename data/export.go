package data

type ExportCharacterDescription struct {
	Type        string `yaml:"type" json:"type"`
	ShortID     uint16 `yaml:"short_id" json:"short_id"`
	LongID      uint64 `yaml:"long_id" json:"long_id"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
	Title       string `yaml:"title" json:"title"`
	Story       string `yaml:"story" json:"story"`
}

type ExportSkillDescription struct {
	Type        string `yaml:"type" json:"type"`
	ShortID     uint16 `yaml:"short_id" json:"short_id"`
	LongID      uint64 `yaml:"long_id" json:"long_id"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
}

type ExportEventDescription struct {
	Type        string `yaml:"type" json:"type"`
	ShortID     uint16 `yaml:"short_id" json:"short_id"`
	LongID      uint64 `yaml:"long_id" json:"long_id"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
}

type ExportCardDescription struct {
	Type        string `yaml:"type" json:"type"`
	ShortID     uint16 `yaml:"short_id" json:"short_id"`
	LongID      uint64 `yaml:"long_id" json:"long_id"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
}

type ExportSummonDescription struct {
	Type        string `yaml:"type" json:"type"`
	ShortID     uint16 `yaml:"short_id" json:"short_id"`
	LongID      uint64 `yaml:"long_id" json:"long_id"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
}

type ExportModifierDescription struct {
	Type        string `yaml:"type" json:"type"`
	ShortID     uint16 `yaml:"short_id" json:"short_id"`
	LongID      uint64 `yaml:"long_id" json:"long_id"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
}

type ExportModDescription struct {
	Type        string `yaml:"type" json:"type"`
	ID          uint64 `yaml:"id" json:"id"`
	PackagePath string `yaml:"package_path" json:"package_path"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
}
