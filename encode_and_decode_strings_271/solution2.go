package encode_and_decode_strings_271

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func (codec *Codec) Encode(strs []string) string {
	// 没什么用
	total := 0
	for _, x := range strs {
		total += len(x)
	}

	output := strings.Builder{}
	output.Grow(total)
	for _, str := range strs {
		output.WriteString(strconv.Itoa(len(str)))
		output.WriteString("_")
		output.WriteString(str)
	}
	return output.String()
}

func (codec *Codec) Decode(str string) []string {
	var output []string
	for i := 0; i < len(str); {
		counter := strings.Builder{}
		for str[i] != '_' {
			counter.WriteByte(str[i])
			i++
		}
		i++
		num, _ := strconv.Atoi(counter.String())

		data := strings.Builder{}
		data.Grow(num)
		for j := 0; j < num; j++ {
			data.WriteByte(str[i])
			i++
		}
		output = append(output, data.String())
	}
	return output
}
