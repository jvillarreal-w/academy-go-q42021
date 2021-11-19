package model

type Pokemon struct {
	ID            uint64 `json:"id"`
	Name          string `json:"name"`
	PrimaryType   string `json:"primary_type"`
	SecondaryType string `json:"secondary_type"`
	Generation    uint64 `json:"generation"`
	Stats         Stats  `json:"stats"`
}

type Stats struct {
	HP             uint64 `json:"hp"`
	Attack         uint64 `json:"attack"`
	Defense        uint64 `json:"defense"`
	SpecialAttack  uint64 `json:"special_attack"`
	SpecialDefense uint64 `json:"special_defense"`
	Speed          uint64 `json:"speed"`
	BaseStatTotal  uint64 `json:"base_stat_total"`
}
