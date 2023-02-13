package types

type Report struct {
	Name       string
	Count      int
	Note       string
	Applicable []*Entry
}
