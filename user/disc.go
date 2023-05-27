package user

type Disc struct {
	Title   string
	Tipe    string
	Creator Artist
	Price   int
	Units   int
}

func (D Disc) Data() (string, string, Artist) {
	return D.Title, D.Tipe, D.Creator
}

func (D Disc) DataArtist() Artist {
	return D.Creator
}

func (D Disc) Amount() int {
	amount := D.Price * D.Units
	return amount
}
