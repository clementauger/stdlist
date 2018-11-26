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

	allPaths := map[string][]string{}
	orderedPaths := []string{}

	for _, fp := range res {
		v := strings.TrimSuffix(fp, ".txt")
		v = strings.TrimPrefix(v, "data/")

		(func(fp, v string) {
			f, err := os.Open(fp)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			sc := bufio.NewScanner(f)
			for sc.Scan() {
				l := sc.Text()
				if _, ok := allPaths[l]; !ok {
					allPaths[l] = []string{}
					orderedPaths = append(orderedPaths, l)
				}
				allPaths[l] = append(allPaths[l], v)
			}
		})(fp, v)
	}

	var out bytes.Buffer

	fmt.Fprintln(&out, "package stdlist")
	fmt.Fprintln(&out, "")
	fmt.Fprintln(&out, "var rawdata = map[string][]string{")
	sort.Strings(orderedPaths)
	for _, p := range orderedPaths {
		vs := allPaths[p]
		for i, v := range vs {
			vs[i] = fmt.Sprintf("%q", v)
		}
		fmt.Fprintf(&out, "    %q: []string{%v},\n", p, strings.Join(vs, ","))
	}
	fmt.Fprintln(&out, "}")
	fmt.Fprintln(&out, "")
	fmt.Fprintln(&out, "func init(){")
	fmt.Fprintln(&out, "  for k,v:=range rawdata{")
	fmt.Fprintln(&out, "    data[k]=v")
	fmt.Fprintln(&out, "  }")
	fmt.Fprintln(&out, "}")

	o, err := format.Source(out.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(o)
}
