package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {

	g := filepath.Join("data", "*.txt")
	res, err := filepath.Glob(g)
	if err != nil {
		log.Fatal(err)
	}

	pathOccurences := map[string]int{}

	allVersions := map[string][]string{}

	for _, fp := range res {
		v := strings.TrimSuffix(fp, ".txt")
		v = strings.TrimPrefix(v, "data/")
		allVersions[v] = []string{}
		(func(fp, v string) {
			f, err := os.Open(fp)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			sc := bufio.NewScanner(f)
			for sc.Scan() {
				l := sc.Text()
				allVersions[v] = append(allVersions[v], l)
				pathOccurences[l]++
			}
		})(fp, v)
	}

	var redundants []string
	for p, c := range pathOccurences {
		if c == len(allVersions) {
			redundants = append(redundants, p)
			for v, ap := range allVersions {
				n := []string{}
				for i := 0; i < len(ap); i++ {
					if ap[i] != p {
						n = append(n, ap[i])
					}
				}
				allVersions[v] = n
			}
		}
	}

	sort.Strings(redundants)

	var out bytes.Buffer

	fmt.Fprintln(&out, "package stdlist")
	fmt.Fprintln(&out, "")
	fmt.Fprintln(&out, "import (\"strings\")")
	fmt.Fprintln(&out, "")
	fmt.Fprintln(&out, "var rawdata = map[string]string{")
	fmt.Fprintf(&out, "   %q: %q,\n", "common", strings.Join(redundants, ","))
	for v, a := range allVersions {
		sort.Strings(a)
		fmt.Fprintf(&out, "    %q: %q,\n", v, strings.Join(a, ","))
	}
	fmt.Fprintln(&out, "}")
	fmt.Fprintln(&out, "")
	fmt.Fprintln(&out, "func init(){")
	fmt.Fprintln(&out, "  for k,v:=range rawdata{")
	fmt.Fprintln(&out, "    data[k]=strings.Split(v, \",\")")
	fmt.Fprintln(&out, "  }")
	fmt.Fprintln(&out, "}")

	o, err := format.Source(out.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(o)
}
