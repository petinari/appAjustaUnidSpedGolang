package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "/Users/robsonpetinari/Documents/EFD_00018_00001.txt"
	arquivo, err := lerLinhasArquivo(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	var _0200List []string
	var _C170List []string
	erros := 0

	for _, line := range arquivo {
		if strings.Contains(line, "|0200|") {
			_0200List = append(_0200List, line)
		}
		if strings.Contains(line, "|C170|") {
			_C170List = append(_C170List, line)
		}
	}

	for _, item0200 := range _0200List {
		linha0200 := strings.Split(item0200, "|")
		for _, itemC170 := range _C170List {
			linhaC170 := strings.Split(itemC170, "|")
			if linha0200[2] == linhaC170[3] {
				if linha0200[6] != linhaC170[6] {
					linhaC170Novo := make([]string, len(linhaC170))
					copy(linhaC170Novo, linhaC170)
					linhaC170Novo[6] = linha0200[6]
					indexC170, err := indiceItemArquivo(itemC170, arquivo)
					if err != nil {
						fmt.Println(err)
						return
					}
					arquivo[indexC170] = strings.Join(linhaC170Novo, "|")
					erros++
				}
			}
		}
	}

	fmt.Println(erros)
	err = escreverLinhasArquivo(path, arquivo)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func lerLinhasArquivo(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var linhas []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}
	return linhas, scanner.Err()
}

func escreverLinhasArquivo(path string, linhas []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, linha := range linhas {
		fmt.Fprintln(writer, linha)
	}
	return writer.Flush()
}

func indiceItemArquivo(item string, arquivo []string) (int, error) {
	for i, linha := range arquivo {
		if linha == item {
			return i, nil
		}
	}
	return -1, fmt.Errorf("item %s não encontrado no arquivo", item)
}
