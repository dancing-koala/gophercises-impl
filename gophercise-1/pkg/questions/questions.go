package questions

import (
	"bytes"
	"encoding/csv"
	"io"
)

type Question struct {
	Text   string
	Answer string
}

func (q *Question) VerifyAnswer(a string) bool {
	return q.Answer == a
}

func ReadCsv(csvData []byte) []*Question {

	list := make([]*Question, 0)

	csvReader := csv.NewReader(bytes.NewReader(csvData))

	for {
		row, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		list = append(list, &Question{
			Text:   row[0],
			Answer: row[1],
		})
	}

	return list
}
