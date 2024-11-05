package event

type FormCreateRequest struct {
	Name      string
	Questions []*Question
}

type Question struct {
	Text    string
	Answers []*Answer
}

type Answer struct {
	Text      string
	IsCorrect bool
}
