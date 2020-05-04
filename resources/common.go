package resources

type Error struct {
	Key         string `json:"key"`
	Description string `json:"description"`
	Trace       string `json:"trace"`
	ErrorData   []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
}
