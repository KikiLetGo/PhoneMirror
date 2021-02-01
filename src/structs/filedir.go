package structs
type FileDir struct {
	Name string `json:"name"`
	Documents []string `json:"documents"`
	Dirs []FileDir `json:"dirs"`
}