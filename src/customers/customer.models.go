package customers

type Customer struct {
	id      string `json:"id" xml:"id"`
	Name    string `json:"name" xml:"name"`
	Address string `json:"address" xml:"address"`
	Zip     string `json:"zip" xml:"zip"`
}
