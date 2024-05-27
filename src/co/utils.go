package co

import (
	"bytes"
	"io"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"golang.org/x/exp/constraints"
)

/*

go get golang.org/x/exp/slices
slices.Contains(things, "foo") // true

*/

// .
func Contains[E comparable](s []E, x E) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

// 특정 리스트에 targets 리스트가 모두 포함되어 있는지 확인.
func ContainsAll(list []string, targets []string) bool {
	for _, target := range targets {
		found := false
		for _, s := range list {
			if s == target {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func IsIncludedInList(numbers []string, targetValue string) bool {
	for _, number := range numbers {
		if number == targetValue {
			return true
		}
	}
	return false
}

/*

    fmt.Println(min([]int{10, 2, 4, 1, 6, 8, 2}))
    fmt.Println(max([]float64{3.2, 5.1, 6.2, 7.6, 8.2, 1.5, 4.8}))



func max[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return m
}

func min[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m > v {
			m = v
		}
	}
	return m
}

*/

func MinMax[T constraints.Ordered](s []T) (T, T) {
	if len(s) == 0 {
		var zero T
		return zero, zero
	}
	m := s[0]
	m2 := s[0]

	for _, v := range s {
		if m < v {
			m = v
		}
		if m2 > v {
			m2 = v
		}
	}
	return m2, m
}

// .
func MinMax2__deplicated(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func MinMaxStr__deplicated(array []string) (string, string) {
	var max string = array[0]
	var min string = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func MinMaxStrInt(array []string) (int, int) {
	var max int = Str2int(array[0])
	var min int = Str2int(array[0])
	for _, value := range array {
		if max < Str2int(value) {
			max = Str2int(value)
		}
		if min > Str2int(value) {
			min = Str2int(value)
		}
	}
	return min, max
}

func SortDirectionCheck(x int) bool {
	s := []int{1, -1}
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

func ItemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func EmptyString(str string) bool {

	return len(strings.TrimSpace(str)) == 0

}

func NotEmptyString(str string) bool {

	return len(strings.TrimSpace(str)) > 0

}

// 익명함수 정의
// func Celldatas(cells []*xlsx.Cell) []string {
// 	var celldatas_ []string
// 	for _, cell := range cells {
// 		celldatas_ = append(celldatas_, cell.String())
// 	}
// 	return celldatas_
// }

func Transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func Transpose2(slice [][]string) []string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	result2 := make([]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}

	for i := 0; i < xl; i++ {
		result2[i] = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strconv.Itoa(i)+"_"+strings.Join(result[i], "_"), "\n", ""), " ", ""), ")", ""), "(", "")
	}
	return result2
}

func Str2int(str string) int {

	intVal, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		// log.Print(err, "  ... [", str, "]")
		return 0
	}
	return intVal

}

// .
func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func FileDelete(filePath string) bool {
	if err := os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// .
func FileCopy(originalFile string, newFile string) error {
	original, err := os.Open(originalFile)
	if err != nil {
		return err
	}
	defer original.Close()

	copy, err := os.Create(newFile)
	if err != nil {
		return err
	}
	defer copy.Close()

	_, err = io.Copy(copy, original)
	if err != nil {
		return err
	}
	err = copy.Sync()
	if err != nil {
		panic(err)
	}

	return err
}

// .
func FolderExistsAndMkdir(path string) (err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, os.ModePerm); err != nil {
			return err
		}
	}
	return err
}

// .
func UNUSED(x ...interface{}) {}

// https://github.com/dustin/go-humanize/blob/master/comma.go
func Commaf(vv float64, round_ uint) string {
	v := roundFloat(vv, round_)
	buf := &bytes.Buffer{}
	if v < 0 {
		buf.Write([]byte{'-'})
		v = 0 - v
	}

	comma := []byte{','}

	parts := strings.Split(strconv.FormatFloat(v, 'f', -1, 64), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(comma)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte{'.'})
		buf.WriteString(parts[1])
		if len(parts[1]) < int(round_) {
			buf.WriteString(strings.Repeat("0", int(round_)-len(parts[1])))
		} else if len(parts[1]) > int(round_) {
			buf.Truncate(buf.Len() - (len(parts[1]) - int(round_)))
		}
	} else {
		buf.WriteString("." + strings.Repeat("0", int(round_)))
	}

	return buf.String()
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{Data: bh.Data, Len: bh.Len} //reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))                 //*(*string)(unsafe.Pointer(&sh))
}
