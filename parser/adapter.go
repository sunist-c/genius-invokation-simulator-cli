package parser

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/adapter"
)

type ElementTypeToStringAdapter struct {
	dictionary map[enum.ElementType]string
}

func (adapter *ElementTypeToStringAdapter) Convert(source enum.ElementType) (success bool, result string) {
	stringResult, exist := adapter.dictionary[source]
	return exist, stringResult
}

func NewElementTypeToStringAdapter() adapter.Adapter[enum.ElementType, string] {
	return &ElementTypeToStringAdapter{
		dictionary: map[enum.ElementType]string{
			enum.ElementCurrency:  "ElementCurrency",
			enum.ElementAnemo:     "ElementAnemo",
			enum.ElementCryo:      "ElementCryo",
			enum.ElementDendro:    "ElementDendro",
			enum.ElementElectro:   "ElementElectro",
			enum.ElementGeo:       "ElementGeo",
			enum.ElementHydro:     "ElementHydro",
			enum.ElementPyro:      "ElementPyro",
			enum.ElementNone:      "ElementNone",
			enum.ElementUndefined: "ElementUndefined",
		},
	}
}

type StringToElementTypeAdapter struct {
	dictionary map[string]enum.ElementType
}

func (adapter *StringToElementTypeAdapter) Convert(source string) (success bool, result enum.ElementType) {
	elementTypeResult, exist := adapter.dictionary[source]
	return exist, elementTypeResult
}

func NewStringToElementTypeAdapter() adapter.Adapter[string, enum.ElementType] {
	return &StringToElementTypeAdapter{
		dictionary: map[string]enum.ElementType{
			"ElementCurrency":  enum.ElementCurrency,
			"ElementAnemo":     enum.ElementAnemo,
			"ElementCryo":      enum.ElementCryo,
			"ElementDendro":    enum.ElementDendro,
			"ElementElectro":   enum.ElementElectro,
			"ElementGeo":       enum.ElementGeo,
			"ElementHydro":     enum.ElementHydro,
			"ElementPyro":      enum.ElementPyro,
			"ElementNone":      enum.ElementNone,
			"ElementSame":      enum.ElementSame,
			"ElementUndefined": enum.ElementUndefined,
		},
	}
}

type LanguageToFullStringAdapter struct {
	dictionary map[enum.Language]string
}

func (adapter *LanguageToFullStringAdapter) Convert(source enum.Language) (success bool, result string) {
	stringResult, exist := adapter.dictionary[source]
	return exist, stringResult
}

func NewLanguageToFullStringAdapter() adapter.Adapter[enum.Language, string] {
	return &LanguageToFullStringAdapter{
		dictionary: map[enum.Language]string{
			enum.ChineseSimplified:  "ChineseSimplified",
			enum.ChineseTraditional: "ChineseTraditional",
			enum.English:            "English",
			enum.French:             "French",
			enum.German:             "German",
			enum.Japanese:           "Japanese",
			enum.Korean:             "Korean",
			enum.Russian:            "Russian",
			enum.Unknown:            "Unknown",
		},
	}
}

type LanguageToShortStringAdapter struct {
	dictionary map[enum.Language]string
}

func (adapter *LanguageToShortStringAdapter) Convert(source enum.Language) (success bool, result string) {
	stringResult, exist := adapter.dictionary[source]
	return exist, stringResult
}

func NewLanguageToShortStringAdapter() adapter.Adapter[enum.Language, string] {
	return &LanguageToShortStringAdapter{
		map[enum.Language]string{
			enum.ChineseSimplified:  "zh-hans",
			enum.ChineseTraditional: "zh-hant",
			enum.English:            "en",
			enum.French:             "fr",
			enum.German:             "de",
			enum.Japanese:           "ja",
			enum.Korean:             "ko",
			enum.Russian:            "ru",
			enum.Unknown:            "unknown",
		},
	}
}

type FullStringToLanguageAdapter struct {
	dictionary map[string]enum.Language
}

func (adapter *FullStringToLanguageAdapter) Convert(source string) (success bool, result enum.Language) {
	languageResult, exist := adapter.dictionary[source]
	return exist, languageResult
}

func NewFullStringToLanguageAdapter() adapter.Adapter[string, enum.Language] {
	return &FullStringToLanguageAdapter{
		dictionary: map[string]enum.Language{
			"ChineseSimplified":  enum.ChineseSimplified,
			"ChineseTraditional": enum.ChineseTraditional,
			"English":            enum.English,
			"French":             enum.French,
			"German":             enum.German,
			"Japanese":           enum.Japanese,
			"Korean":             enum.Korean,
			"Russian":            enum.Russian,
			"Unknown":            enum.Unknown,
		},
	}
}

type ShortStringToLanguageAdapter struct {
	dictionary map[string]enum.Language
}

func (adapter *ShortStringToLanguageAdapter) Convert(source string) (success bool, result enum.Language) {
	languageResult, exist := adapter.dictionary[source]
	return exist, languageResult
}

func NewShortStringToLanguageAdapter() adapter.Adapter[string, enum.Language] {
	return &ShortStringToLanguageAdapter{
		dictionary: map[string]enum.Language{
			"zh-cn":   enum.ChineseSimplified,
			"zh-sg":   enum.ChineseSimplified,
			"zh-hans": enum.ChineseSimplified,
			"zh-tw":   enum.ChineseTraditional,
			"zh-hk":   enum.ChineseTraditional,
			"zh-mo":   enum.ChineseTraditional,
			"zh-hant": enum.ChineseTraditional,
			"en":      enum.English,
			"fr":      enum.French,
			"de":      enum.German,
			"ja":      enum.Japanese,
			"ko":      enum.Korean,
			"ru":      enum.Russian,
			"unknown": enum.Unknown,
		},
	}
}

type ReactionToStringAdapter struct {
	dictionary map[enum.Reaction]string
}

func (adapter *ReactionToStringAdapter) Convert(source enum.Reaction) (success bool, result string) {
	stringResult, exist := adapter.dictionary[source]
	return exist, stringResult
}

func NewReactionToStringAdapter() adapter.Adapter[enum.Reaction, string] {
	return &ReactionToStringAdapter{
		dictionary: map[enum.Reaction]string{
			enum.ReactionNone:              "ReactionNone",
			enum.ReactionMelt:              "ReactionMelt",
			enum.ReactionVaporize:          "ReactionVaporize",
			enum.ReactionOverloaded:        "ReactionOverloaded",
			enum.ReactionSuperconduct:      "ReactionSuperconduct",
			enum.ReactionFrozen:            "ReactionFrozen",
			enum.ReactionElectroCharged:    "ReactionElectroCharged",
			enum.ReactionBurning:           "ReactionBurning",
			enum.ReactionBloom:             "ReactionBloom",
			enum.ReactionQuicken:           "ReactionQuicken",
			enum.ReactionCryoSwirl:         "ReactionCryoSwirl",
			enum.ReactionElectroSwirl:      "ReactionElectroSwirl",
			enum.ReactionHydroSwirl:        "ReactionHydroSwirl",
			enum.ReactionPyroSwirl:         "ReactionPyroSwirl",
			enum.ReactionCryoCrystalize:    "ReactionCryoCrystalize",
			enum.ReactionElectroCrystalize: "ReactionElectroCrystalize",
			enum.ReactionHydroCrystalize:   "ReactionHydroCrystalize",
			enum.ReactionPyroCrystalize:    "ReactionPyroCrystalize",
		},
	}
}

type StringToReactionAdapter struct {
	dictionary map[string]enum.Reaction
}

func (adapter *StringToReactionAdapter) Convert(source string) (success bool, result enum.Reaction) {
	reactionResult, exist := adapter.dictionary[source]
	return exist, reactionResult
}

func NewStringToReactionAdapter() adapter.Adapter[string, enum.Reaction] {
	return &StringToReactionAdapter{
		dictionary: map[string]enum.Reaction{
			"ReactionNone":              enum.ReactionNone,
			"ReactionMelt":              enum.ReactionMelt,
			"ReactionVaporize":          enum.ReactionVaporize,
			"ReactionOverloaded":        enum.ReactionOverloaded,
			"ReactionSuperconduct":      enum.ReactionSuperconduct,
			"ReactionFrozen":            enum.ReactionFrozen,
			"ReactionElectroCharged":    enum.ReactionElectroCharged,
			"ReactionBurning":           enum.ReactionBurning,
			"ReactionBloom":             enum.ReactionBloom,
			"ReactionQuicken":           enum.ReactionQuicken,
			"ReactionCryoSwirl":         enum.ReactionCryoSwirl,
			"ReactionElectroSwirl":      enum.ReactionElectroSwirl,
			"ReactionHydroSwirl":        enum.ReactionHydroSwirl,
			"ReactionPyroSwirl":         enum.ReactionPyroSwirl,
			"ReactionCryoCrystalize":    enum.ReactionCryoCrystalize,
			"ReactionElectroCrystalize": enum.ReactionElectroCrystalize,
			"ReactionHydroCrystalize":   enum.ReactionHydroCrystalize,
			"ReactionPyroCrystalize":    enum.ReactionPyroCrystalize,
		},
	}
}

type WeaponTypeToStringAdapter struct {
	dictionary map[enum.WeaponType]string
}

func (adapter *WeaponTypeToStringAdapter) Convert(source enum.WeaponType) (success bool, result string) {
	stringResult, exist := adapter.dictionary[source]
	return exist, stringResult
}

func NewWeaponTypeToStringAdapter() adapter.Adapter[enum.WeaponType, string] {
	return &WeaponTypeToStringAdapter{
		dictionary: map[enum.WeaponType]string{
			enum.WeaponSword:    "WeaponSword",
			enum.WeaponClaymore: "WeaponClaymore",
			enum.WeaponBow:      "WeaponBow",
			enum.WeaponCatalyst: "WeaponCatalyst",
			enum.WeaponPolearm:  "WeaponPolearm",
			enum.WeaponOthers:   "WeaponOthers",
		},
	}
}

type StringToWeaponTypeAdapter struct {
	dictionary map[string]enum.WeaponType
}

func (adapter *StringToWeaponTypeAdapter) Convert(source string) (success bool, result enum.WeaponType) {
	weaponTypeResult, exist := adapter.dictionary[source]
	return exist, weaponTypeResult
}

func NewStringToWeaponTypeAdapter() adapter.Adapter[string, enum.WeaponType] {
	return &StringToWeaponTypeAdapter{
		dictionary: map[string]enum.WeaponType{
			"WeaponSword":    enum.WeaponSword,
			"WeaponClaymore": enum.WeaponClaymore,
			"WeaponBow":      enum.WeaponBow,
			"WeaponCatalyst": enum.WeaponCatalyst,
			"WeaponPolearm":  enum.WeaponPolearm,
			"WeaponOthers":   enum.WeaponOthers,
		},
	}
}

type AffiliationToStringAdapter struct {
	dictionary map[enum.Affiliation]string
}

func (adapter *AffiliationToStringAdapter) Convert(source enum.Affiliation) (success bool, result string) {
	stringResult, exist := adapter.dictionary[source]
	return exist, stringResult
}

func NewAffiliationToStringAdapter() adapter.Adapter[enum.Affiliation, string] {
	return &AffiliationToStringAdapter{
		dictionary: map[enum.Affiliation]string{
			enum.AffiliationMondstadt: "AffiliationMondstadt",
			enum.AffiliationLiyue:     "AffiliationLiyue",
			enum.AffiliationInazuma:   "AffiliationInazuma",
			enum.AffiliationSumeru:    "AffiliationSumeru",
			enum.AffiliationFatui:     "AffiliationFatui",
			enum.AffiliationHilichurl: "AffiliationHilichurl",
			enum.AffiliationMonster:   "AffiliationMonster",
			enum.AffiliationUndefined: "AffiliationUndefined",
		},
	}
}

type StringToAffiliationAdapter struct {
	dictionary map[string]enum.Affiliation
}

func (adapter *StringToAffiliationAdapter) Convert(source string) (success bool, result enum.Affiliation) {
	affiliationResult, exist := adapter.dictionary[source]
	return exist, affiliationResult
}

func NewStringToAffiliationAdapter() adapter.Adapter[string, enum.Affiliation] {
	return &StringToAffiliationAdapter{
		dictionary: map[string]enum.Affiliation{
			"AffiliationMondstadt": enum.AffiliationMondstadt,
			"AffiliationLiyue":     enum.AffiliationLiyue,
			"AffiliationInazuma":   enum.AffiliationInazuma,
			"AffiliationSumeru":    enum.AffiliationSumeru,
			"AffiliationFatui":     enum.AffiliationFatui,
			"AffiliationHilichurl": enum.AffiliationHilichurl,
			"AffiliationMonster":   enum.AffiliationMonster,
			"AffiliationUndefined": enum.AffiliationUndefined,
		},
	}
}
