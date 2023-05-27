package user

type Tape struct {
	Title   string
	Tipe    string
	Creator Artist
}

func (T Tape) Data() (string, string, Artist) {
	return T.Title, T.Tipe, T.Creator
}
