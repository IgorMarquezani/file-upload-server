package files

type File struct {
	Name     string
	Size     float64
	Modified string
}

type ServerFiles struct {
	Ip    string
	Port  string
	Files []File
}
