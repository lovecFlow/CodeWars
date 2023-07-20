//Функция, суммирующая все цифры в строке (если цифра не заканчивается на единицах, то берётся как отдельное число) 
//=============================================================================================================================
//Моё решение: 
package main

import (
	"log"
	"regexp"
	"strconv"
)

func main() {
	SumOfIntsInString("The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog")
}

func SumOfIntsInString(s string) int {
	re := regexp.MustCompile("[0-9]+") //Регулярное выражение для поиска цифр в строке
	newr := re.FindAllString(s, -1) //Поиск цифр
	result := 0 //Переменная для результата
	for _, val := range newr { //Цикл для прохода по newr (там лежат цифры) 
		vl, err := strconv.Atoi(val) /
		if err != nil {
			log.Fatal(err)
		} //Конвертируем цифры из стрингов в инты 
		result += vl //Суммируем цифры
	}
	return result
}
//=============================================================================================================================
//Более изящное и простое решение: 

func SumOfIntegersInString(strng string) int {
  r := 0
  for _, s := range regexp.MustCompile(`\d+`).FindAllString(strng, -1) {
    x, _ := strconv.Atoi(s)
    r += x
  }
  return r
}
