package eos

// Extensions extensions in trx, abi or actions
type Extensions struct {
	Exts []*Extension `json:"datas"`
}

func NewExtensions() *Extensions {
	return &Extensions{
		Exts: make([]*Extension, 0, 8),
	}
}
