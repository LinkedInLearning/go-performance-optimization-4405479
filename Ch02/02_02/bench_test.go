package tokenizer

import "testing"

var text = `
TO THE RED-HEADED LEAGUE: On account of the bequest of the late
Ezekiah Hopkins, of Lebanon, Pennsylvania, U.S.A., there is now another
vacancy open which entitles a member of the League to a salary of £ 4 a
week for purely nominal services. All red-headed men who are sound in
body and mind and above the age of twenty-one years, are eligible.
Apply in person on Monday, at eleven o’clock, to Duncan Ross, at the
offices of the League, 7 Pope’s Court, Fleet Street.
`

func BenchmarkTokenize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tokens := Tokenize(text)
		if len(tokens) != 88 {
			b.Fatal(len(tokens))
		}
	}
}
