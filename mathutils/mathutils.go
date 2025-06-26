// файл: mathutils/mathutils.go
package mathutils

import (
	"errors"
	//"fmt" // Может понадобиться для ошибок
)

// Add возвращает сумму двух целых чисел.
func Add(a, b int) int {
	return a + b
}

// Subtract возвращает разность двух целых чисел (a - b).
func Subtract(a, b int) int {
	return a - b
}

// Factorial вычисляет факториал неотрицательного целого числа n.
// Для отрицательных чисел должна возвращаться ошибка.
// TODO: Реализуйте эту функцию. Учтите граничные случаи (0! = 1).
// Используйте errors.New или fmt.Errorf для создания ошибки.
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("факториал отрицательного числа не определен") // Пример ошибки
	}
	if n == 0 || n == 1 {
		return 1, nil
	}
	res := 1
	for num := 2; num <= n; num++ {
		res *= num
	}
	return res, nil
}
