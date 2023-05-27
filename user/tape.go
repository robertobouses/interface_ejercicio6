package user

type Tape struct {
	Title   string
	Tipe    string
	Creator Artist
	Price   int
	Units   int
}

func (T Tape) Data() (string, string, Artist) {
	return T.Title, T.Tipe, T.Creator
}

func (T Tape) DataArtist() Artist {
	return T.Creator
}

func (T Tape) Amount() int {
	amount := T.Price * T.Units
	return amount
}
