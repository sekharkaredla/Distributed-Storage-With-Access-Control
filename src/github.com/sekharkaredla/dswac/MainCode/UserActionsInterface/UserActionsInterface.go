package UserActions

type UserOptions interface {
	AddNewFileToIPFS(string) (string, error)
}
