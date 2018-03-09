package questions

import "testing"

func TestVerifyAnswer_IsRight(t *testing.T) {
	q := &Question{
		Text:   "Not relevant",
		Answer: "AAA",
	}

	a := "AAA"

	if !q.VerifyAnswer(a) {
		fail(t, "answers should match")
	}
}

func TestVerifyAnswer_IsRightDespiteCase(t *testing.T) {
	q := &Question{
		Text:   "Not relevant",
		Answer: "aAa",
	}

	a := "AAA"

	if !q.VerifyAnswer(a) {
		fail(t, "answers should match")
	}
}

func TestVerifyAnswer_IsWrong(t *testing.T) {
	q := &Question{
		Text:   "Not relevant",
		Answer: "AAA",
	}

	a := "BBB"

	if q.VerifyAnswer(a) {
		fail(t, "answers should not match")
	}
}

func fail(t *testing.T, msg string) {
	t.Log("Fail: " + msg)
	t.Fail()
}
