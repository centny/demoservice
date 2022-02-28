package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: item2md <item title> <item file>\n")
		os.Exit(1)
		return
	}
	allData, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		fmt.Printf("read file %v fail with %v\n", os.Args[2], err)
		os.Exit(1)
	}
	allText := fmt.Sprintf("## <a name=\"metadata\"></a>%v\n\n```.go\n"+string(allData)+"```\n", os.Args[1])
	metaReg := regexp.MustCompile(`/.*metadata:.*/`)
	outText := metaReg.ReplaceAllStringFunc(allText, func(s string) string {
		s = strings.Trim(s, "/*")
		s = strings.TrimSpace(s)
		s = strings.TrimPrefix(s, "metadata:")
		s = strings.TrimSpace(s)
		s = fmt.Sprintf("```\n\n#### <a name=\"metadata-%v\"></a>%v\n\n```.go\n", s, s)
		return s
	})
	os.Stdout.WriteString(outText)
}
