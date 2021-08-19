// Пакет datafile предназначен для чтения данных из файлов
package datafile

import (
  "bufio"
  "os"
)

//GetStrings читает строку из каждой строки данных файла
func GetStrings (filename string) ([]string, error){
  var lines []string //в переменной хранится срез строк
  file, err := os.Open(filename) //открывает файл с переданным именем
  if err != nil {
    return nil, err
  }
  scanner := bufio.NewScanner(file) //для файла созд новое значение Scanner
  for scanner.Scan() { //читает строку из файла
    line := scanner.Text() //добавление строки в файл через срез
    lines = append(lines, line)
  }
  err = file.Close() //закрывает файл для освобрждения ресурсов
  if err != nil {
    return nil, err
  }
  if scanner.Err() != nil {
    return nil, scanner.Err()
  }
  return lines, nil
}
