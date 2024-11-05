package event

type Form struct {
	Name      string
	Questions []Question
}

type Question struct {
	Name    string
	Answers []Answer
}

type Answer struct {
	Text      string
	IsCorrect bool
}
