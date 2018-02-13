package flashCard

import (
	"math/rand"
	"time"
	"strings"
	"strconv"
)

const (
	plus  = "+"
	minus = "-"
)

type MathProblem struct {
	N1       int
	N2       int
	Operator string
	Answer   int
}

var r1 *rand.Rand

func init() {
	r1 = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// GenerateDiff produces the data for a subtraction flash card.
func GenerateDiff() (MathProblem) {
	n1 := r1.Intn(19)
	n2 := r1.Intn(10)
	if n1 < n2 {
		n1, n2 = n2, n1
	}
	diff := n1 - n2
	for diff >= 10 {
		n1 -= 4
		diff = n1 - n2
	}
	return MathProblem{n1, n2, minus, diff}
}

// GenerateAdd produces the data for an addition flash card.
func GenerateAdd() (MathProblem) {
	n1 := r1.Intn(10)
	n2 := r1.Intn(10)
	sum := n1 + n2
	return MathProblem{n1, n2, plus, sum}
}

// ParsePrevious takes the form data string and converts it back into a MathProblem
func ParsePrevious(str string) (MathProblem) {
	str = strings.Trim(str, "{} ")
	xs := strings.Fields(str)
	n1, _ := strconv.Atoi(xs[0])
	n2, _ := strconv.Atoi(xs[1])
	opr := xs[2]
	ans, _ := strconv.Atoi(xs[3])
	return MathProblem{n1, n2, opr, ans}
}
