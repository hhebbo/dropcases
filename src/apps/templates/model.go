package templates

type template struct {
	Name       string       `json:"Name"`
	FolderTree []FolderTree `json:"FolderTree"`
}

type FolderTree struct {
	FolderName string       `json:"FolderName"`
	Childern   []FolderTree `json:"Childern"`
}
