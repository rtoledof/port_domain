package models

type Port struct {
	ID          string
	Code        string
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float32
	Province    string
	Timezone    string
	Unlocs      []string
}
