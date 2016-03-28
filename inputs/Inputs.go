package inputs

type Input interface {
	ReadLine() (string, error)
	Close()
	GetType() string
	GetName() string
}
