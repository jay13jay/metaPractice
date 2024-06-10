package wronganswers

func GetWrongAnswers(N int32, C string) string {
  a := make(map[rune]rune)
	a['A'] = 'B'
	a['B'] = 'A'
	var answers []rune
  for _, v := range C {
		answers = append(answers, a[(v)])
  }
	answer := string(answers)
  return string(answer)
}
