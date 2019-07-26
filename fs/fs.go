package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type FileService interface {
	Save([]int) error
	Read() ([]int, error)
}

type service struct {
	name      string
	delimeter string
}

func (s service) sliceToString(slc []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(slc), " ", s.delimeter, -1), "[]")
}

func (s service) stringToSlice(str string) []int {
	arrStr := strings.Split(str, s.delimeter)

	arr := []int{}

	for _, i := range arrStr {
		j, err := strconv.Atoi(i)
		if err != nil {
			arr = append(arr, 0)
		}
		arr = append(arr, j)
	}
	return arr
}

func (s service) Save(slc []int) error {
	f, err := os.OpenFile(s.name, os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()
	text := s.sliceToString(slc)

	if _, err = f.WriteString(text); err != nil {
		return err
	}

	return nil
}

func (s service) Read() ([]int, error) {
	bufFile, err := ioutil.ReadFile(s.name)
	if err != nil {
		return []int{}, err
	}
	str := string(bufFile)

	arr := s.stringToSlice(str)

	return arr, nil
}

func NewFileService(name string) FileService {
	return &service{
		name:      name,
		delimeter: " ",
	}
}
