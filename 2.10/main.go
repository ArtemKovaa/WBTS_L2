package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
	"github.com/spf13/pflag"
)

var (
	columnFlag  = pflag.IntP("k", "k", 1, "Column number to sort by (1-based)")
	numericFlag = pflag.BoolP("n", "n", false, "Sort by numeric value")
	reverseFlag = pflag.BoolP("r", "r", false, "Reverse the sort order")
	uniqueFlag  = pflag.BoolP("u", "u", false, "Output only unique lines")
	monthFlag   = pflag.BoolP("M", "M", false, "Sort by month name (Jan..Dec)")
	ignoreBlank = pflag.BoolP("b", "b", false, "Ignore trailing blanks in key")
	checkSorted = pflag.BoolP("c", "c", false, "Check whether input is sorted")
	humanFlag   = pflag.BoolP("h", "h", false, "Sort human-readable sizes (e.g. 2K, 1G)")
	fieldSep    = "\t"
	monthNames  = []string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	monthMap map[string]int
)

func init() {
	monthMap = make(map[string]int, len(monthNames))
	for i, m := range monthNames {
		monthMap[strings.ToLower(m)] = i + 1
	}
}

type record struct {
	line   string
	key    string
	num    float64
	month  int
	human  float64
	parsed bool
}

func parseKey(line string, col int, ignoreTrailingBlanks bool) (string, error) {
	fields := strings.Fields(line)
	if col < 1 || col > len(fields) {
		return "", errors.New("column number out of range")
	}
	key := fields[col-1]
	if ignoreTrailingBlanks {
		key = strings.TrimRightFunc(key, unicode.IsSpace)
	}
	return key, nil
}

func parseMonth(key string) int {
	return monthMap[strings.ToLower(key)]
}

var humanSizeRegexp = regexp.MustCompile(`^([0-9]*\.?[0-9]+)([KMGTPE]?)(i?)B?$`)

func parseHuman(s string) (float64, error) {
	matches := humanSizeRegexp.FindStringSubmatch(strings.TrimSpace(s))
	if len(matches) < 2 {
		return 0, fmt.Errorf("invalid human size: %s", s)
	}
	base, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, err
	}
	unit := strings.ToUpper(matches[2])
	mult := 1.0
	switch unit {
	case "K":
		mult = 1 << 10
	case "M":
		mult = 1 << 20
	case "G":
		mult = 1 << 30
	case "T":
		mult = 1 << 40
	case "P":
		mult = 1 << 50
	case "E":
		mult = 1 << 60
	}
	return base * mult, nil
}

func extractNumberFromString(s string) float64 {
	numExtractRegexp := regexp.MustCompile(`\d+`)
	match := numExtractRegexp.FindString(s)
	if match == "" {
		return 0
	}
	n, _ := strconv.ParseFloat(match, 64)
	return n
}

func naturalLess(a, b string) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		ra, sizea := rune(a[i]), 1
		rb, sizeb := rune(b[j]), 1

		if ra >= utf8.RuneSelf {
			ra, sizea = utf8.DecodeRuneInString(a[i:])
		}
		if rb >= utf8.RuneSelf {
			rb, sizeb = utf8.DecodeRuneInString(b[j:])
		}

		if unicode.IsDigit(ra) && unicode.IsDigit(rb) {
			na, ni := extractNumberString(a, i)
			nb, nj := extractNumberString(b, j)
			if na != nb {
				return na < nb
			}
			i = ni
			j = nj
			continue
		}

		if ra != rb {
			return ra < rb
		}
		i += sizea
		j += sizeb
	}
	return len(a) < len(b)
}

func extractNumberString(s string, pos int) (int, int) {
	start := pos
	for pos < len(s) {
		r, size := utf8.DecodeRuneInString(s[pos:])
		if !unicode.IsDigit(r) {
			break
		}
		pos += size
	}
	n, _ := strconv.Atoi(s[start:pos])
	return n, pos
}

func makeRecords(lines []string, col int, ignoreTrailingBlanks bool) ([]record, error) {
	records := make([]record, 0, len(lines))
	for _, line := range lines {
		key, err := parseKey(line, col, ignoreTrailingBlanks)
		if err != nil {
			key = ""
		}
		rec := record{line: line, key: key}

		if *numericFlag {
			rec.num = extractNumberFromString(key)
			rec.parsed = true
		} else if *monthFlag {
			mon := parseMonth(key)
			if mon > 0 {
				rec.month = mon
				rec.parsed = true
			}
		} else if *humanFlag {
			if f, err := parseHuman(key); err == nil {
				rec.human = f
				rec.parsed = true
			}
		}
		records = append(records, rec)
	}
	return records, nil
}

type sorter struct {
	records []record
	reverse bool
}

func (s sorter) Len() int           { return len(s.records) }
func (s sorter) Swap(i, j int)      { s.records[i], s.records[j] = s.records[j], s.records[i] }
func (s sorter) Less(i, j int) bool {
	a, b := s.records[i], s.records[j]
	var cmp int

	switch {
	case *numericFlag:
		if a.parsed && b.parsed {
			switch {
			case a.num < b.num:
				cmp = -1
			case a.num > b.num:
				cmp = 1
			default:
				cmp = 0
			}
		} else {
			cmp = strings.Compare(a.key, b.key)
		}
	case *monthFlag:
		if a.parsed && b.parsed {
			switch {
			case a.month < b.month:
				cmp = -1
			case a.month > b.month:
				cmp = 1
			default:
				cmp = 0
			}
		} else {
			cmp = strings.Compare(a.key, b.key)
		}
	case *humanFlag:
		if a.parsed && b.parsed {
			switch {
			case a.human < b.human:
				cmp = -1
			case a.human > b.human:
				cmp = 1
			default:
				cmp = 0
			}
		} else {
			cmp = strings.Compare(a.key, b.key)
		}
	default:
		if naturalLess(a.key, b.key) {
			cmp = -1
		} else if naturalLess(b.key, a.key) {
			cmp = 1
		} else {
			cmp = 0
		}
	}

	if s.reverse {
		return cmp > 0
	}
	return cmp < 0
}

func printSorted(lines []string, unique bool) {
	var prev string
	for i, line := range lines {
		if unique && i > 0 && line == prev {
			continue
		}
		prev = line
		fmt.Println(line)
	}
}

func checkSortedLines(records []record, reverse bool) int {
	for i := 1; i < len(records); i++ {
		less := sorter{records: records, reverse: reverse}.Less(i, i-1)
		if less {
			return i
		}
	}
	return -1
}

func main() {
	pflag.Parse()

	in := os.Stdin
	if flag.NArg() > 0 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		in = f
	}

	reader := bufio.NewReader(in)

	lines := make([]string, 0, 1024)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
			os.Exit(1)
		}
		line = strings.TrimRight(line, "\r\n")
		if len(line) > 0 || (err == nil) {
			lines = append(lines, line)
		}
		if err == io.EOF {
			break
		}
	}

	records, err := makeRecords(lines, *columnFlag, *ignoreBlank)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error processing input: %v\n", err)
		os.Exit(1)
	}

	if *checkSorted {
		pos := checkSortedLines(records, *reverseFlag)
		if pos < 0 {
			fmt.Println("Input is sorted")
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "Input is not sorted at line %d\n", pos+1)
		os.Exit(1)
	}

	sorterInst := sorter{records: records, reverse: *reverseFlag}
	sort.Sort(sorterInst)

	outLines := make([]string, len(records))
	for i, r := range sorterInst.records {
		outLines[i] = r.line
	}

	printSorted(outLines, *uniqueFlag)
}
