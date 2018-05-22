package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/alediaferia/prefixmap"
)

var similarity float64
var datasource *prefixmap.PrefixMap

type Match struct {
	Value      string
	Similarity float64
}

func (match *Match) Print() {
	fmt.Printf("match: \t%s\tsimilarity: %.2f\t\n", match.Value, match.Similarity)
}

func init() {
	flag.Float64Var(&similarity, "similarity", 0.3, "the similarity target to use when searching")

	datasource = prefixmap.New()
}

func main() {
	flag.Parse()

	input := flag.Arg(0)
	if input == "" {
		fmt.Println("Please, specify an input string")
		os.Exit(1)
	}

	// here we populate the datasource
	for _, country := range Countries {
		//把国家做切分，不同部分代表相同的国家
		parts := strings.Split(strings.ToLower(country), " ")
		for _, part := range parts {
			datasource.Insert(part, country)
		}
	}

	values := datasource.GetByPrefix(strings.ToLower(input))
	results := make([]*Match, 0, len(values))
	for _, v := range values {
		// 反射获取到真正的值
		value := v.(string)
		//根据此算法计算相似度
		s := ComputeSimilarity(len(value), len(input), LevenshteinDistance(value, input))
		if s >= similarity {
			m := &Match{value, s}
			results = append(results, m)
		}
	}

	fmt.Printf("Result for target similarity: %.2f\n", similarity)
	PrintMatches(results)
}

func PrintMatches(matches []*Match) {
	for _, m := range matches {
		m.Print()
	}
}
