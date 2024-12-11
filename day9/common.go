package day9

type Block struct {
	Id          string
	TotalSize   int
	FilledSpace int
	FreeSpace   int
}

type Block2 struct {
	Used    []string
	NewUsed []string
	Free    []string
}
