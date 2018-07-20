package model


type Weapon struct {
	Name 				string	`json:"name,omitempty"`
	Rarity 				string	`json:"rarity,omitempty"`
	FireRate			float64	`json:"firerate:omitempty"`
	MagSize				int		`json:"mag_size:omitempty"`
	ReloadTime       	float64 `json:"reloadtime,omitempty"`
	Multiplier       	int     `json:"multiplier,omitempty"`
	DamageBody       	int     `json:"damagebody,omitempty"`
	DamageHead       	int     `json:"damagehead,omitempty"`
	FallOff          	string  `json:"fallof,omitempty"`
	DPSBody          	int     `json:"dpsbody,omitempty"`
	DPSHead          	int     `json:"dpshead,omitempty"`
	ShotsKillAvgBody	float64 `json:"shotskillavgbody,omitempty"`
	ShotsKillAvgHead	float64 `json:"shotskillavghead,omitempty"`
}