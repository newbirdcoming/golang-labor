package person

//go:generate mockgen -destination=../mocks/mock_person.go -package=mocks -source=person.go
type Person interface {
	Eat() string
	Sleep(name string) string
}
