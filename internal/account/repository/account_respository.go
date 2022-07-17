package repository

type AccountRepo struct {
	InputFile  string
	OutputFile string
}

const (
	delimiterCSV = ';'
)

func NewAccountRepo(inputFile, outputFile string) *AccountRepo {
	return &AccountRepo{
		InputFile:  inputFile,
		OutputFile: outputFile,
	}
}
