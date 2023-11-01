package main

import (
	"bufio"
	"fmt"
	"github.com/pborman/getopt"
	"os"
	"sort"
	"strconv"
)

func getFile(path string, reader bool) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	data := make([]string, 0)
	set := make(map[string]struct{})

	for scanner.Scan() {
		if reader {
			if _, ok := set[scanner.Text()]; !ok {
				set[scanner.Text()] = struct{}{}
			} else {
				continue
			}
		}
		data = append(data, scanner.Text())
	}

	return data, nil
}

func sortFile(data []string, flag bool) []string {

	if flag {
		sort.Slice(data, func(i, j int) bool {
			vi, _ := strconv.Atoi(data[i])
			vj, _ := strconv.Atoi(data[j])
			return vi < vj
		})
	} else {
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
	}

	return data
}

func main() {
	filepath := getopt.String('f', "task.go", "file")
	n := *getopt.Bool('n', "сортировать по числовому значению")
	r := *getopt.Bool('r', "сортировать в обратном порядке")
	u := *getopt.Bool('u', "не выводить повторяющиеся строки")
	getopt.Parse()
	file, err := getFile(*filepath, u)
	if err != nil {
		panic(err)
	}

	file = sortFile(file, n)

	if r {
		for i := len(file) - 1; i > 0; i-- {
			fmt.Println(file[i])
		}
	} else {
		for _, value := range file {
			fmt.Println(value)
		}
	}
}

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
