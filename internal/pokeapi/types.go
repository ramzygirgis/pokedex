package pokeapi

type LocationAreaResult struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationArea struct {
	Count    int                  `json:"count"`
	Next     string               `json:"next"`
	Previous string               `json:"previous"`
	Results  []LocationAreaResult `json:"results"`
}

type ExploreResult struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int                   `json:"game_index"`
	ID                   int                   `json:"id"`
	Location             NamedAPIResource      `json:"location"`
	Name                 string                `json:"name"`
	Names                []LocalizedName       `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource `json:"encounter_method"`
	VersionDetails  []VersionDetail  `json:"version_details"`
}

type VersionDetail struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type LocalizedName struct {
	Language NamedAPIResource `json:"language"`
	Name     string           `json:"name"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource       `json:"pokemon"`
	VersionDetails []VersionEncounterInfo `json:"version_details"`
}

type VersionEncounterInfo struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int               `json:"max_chance"`
	Version          NamedAPIResource  `json:"version"`
}

type EncounterDetail struct {
	Chance         int                `json:"chance"`
	ConditionValues []NamedAPIResource `json:"condition_values"`
	MaxLevel       int                `json:"max_level"`
	Method         NamedAPIResource   `json:"method"`
	MinLevel       int                `json:"min_level"`
}
