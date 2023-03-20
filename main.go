package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "/Users/robsonpetinari/Documents/EFD_00021_00001.txt"
	arquivo, err := lerLinhasArquivo(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	indexMap := make(map[string]int, len(arquivo))
	for i, line := range arquivo {
		indexMap[line] = i
	}

	erros := 0

	for i, line := range arquivo {
		if strings.Contains(line, "|0200|") {
			linha0200 := strings.Split(line, "|")
			for j := i + 1; j < len(arquivo); j++ {
				if strings.Contains(arquivo[j], "|C170|") {
					linhaC170 := strings.Split(arquivo[j], "|")
					if linha0200[2] == linhaC170[3] && linha0200[6] != linhaC170[6] {
						linhaC170[6] = linha0200[6]
						indexC170, ok := indexMap[arquivo[j]]
						if !ok {
							fmt.Printf("Erro: índice não encontrado para linha %q\n", arquivo[j])
							continue
						}
						arquivo[indexC170] = strings.Join(linhaC170, "|")
						erros++
					}
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
