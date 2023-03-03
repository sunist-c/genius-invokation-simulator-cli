package data

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
)

type CharacterDescription struct {
	Type        enum.DescriptionType
	Name        string
	Description string
	Title       string
	Story       string
}

type SkillDescription struct {
	Type        enum.DescriptionType
	Name        string
	Description string
}

type EventDescription struct {
	Type        enum.DescriptionType
	Name        string
	Description string
}

type CardDescription struct {
	Type        enum.DescriptionType
	Name        string
	Description string
}

type SummonDescription struct {
	Type        enum.DescriptionType
	Name        string
	Description string
}

type ModifierDescription struct {
	Type        enum.DescriptionType
	Name        string
	Description string
}

type ModDescription struct {
	Type        enum.DescriptionType
	Name        string
	Description string
}
