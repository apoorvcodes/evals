package questions



type QuestionType int

const (
	EQ QuestionType = iota // Emotional Intelligence questions
	IQ                     // Intelligence questions
	Learning               // Learning-focused questions  
	Safety                 // Safety-related questions
)

// String returns the string representation of QuestionType
func (qt QuestionType) String() string {
	switch qt {
	case EQ:
		return "EQ"
	case IQ:
		return "IQ" 
	case Learning:
		return "Learning"
	case Safety:
		return "Safety"
	default:
		return "Unknown"
	}
}

// FromString converts a string to QuestionType
func FromString(s string) QuestionType {
	switch s {
	case "EQ":
		return EQ
	case "IQ":
		return IQ
	case "Learning":
		return Learning
	case "Safety":
		return Safety
	default:
		return QuestionType(-1) // Invalid question type
	}
}

type QuestionGenerator interface {
	GenerateQuestions() []string
}

var QuestionGenerators = map[QuestionType]QuestionGenerator{
	EQ:       GetEngagementQuestions(),
	IQ:       NewIQGenerator(), 
	Learning: NewLearningGenerator(),
	Safety:   NewSafetyGenerator(),
}



