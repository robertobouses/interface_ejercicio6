package user

type Disc struct {
	Title   string
	Tipe    string
	Creator Artist
}

func (D Disc) Data() (string, string, Artist) {
	return D.Title, D.Tipe, D.Creator
}
