// файл: mathutils/mathutils_test.go
package mathutils

import (
	"strings"
	"testing"
	// "errors" // Может понадобиться для сравнения ошибок
)

func TestAdd(t *testing.T) {
	// TODO: Определите тестовые случаи в виде таблицы (слайс структур)
	testCases := []struct {
		name     string // Имя под-теста
		a        int    // Входное значение 1
		b        int    // Входное значение 2
		expected int    // Ожидаемый результат
	}{
		// Добавьте несколько тестовых случаев:
		{"Сумма положительных", 2, 3, 5},
		{"Сумма отрицательных", -2, -13, -15},
		{"Сложение с нулем", 0, 3, 3},
		{"Сложение положительного и отрицательного", -2, 3, 1},
		{"Сложение нулей", 0, 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// TODO: Вызовите тестируемую функцию Add
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("ожидали %d, получили %d", tc.expected, result)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	// TODO: Напишите тесты для функции Subtract, используя табличный подход.
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Вычитание положительных", 5, 3, 2},
		{"Вычитание отрицательных", -5, -7, 2},
		{"Вычитание нуля", 5, 0, 5},
		{"Вычитание из нуля", 0, -1, 1},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Subtract(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("ожидали %d, получили %d", tc.expected, result)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	// TODO: Напишите тесты для функции Factorial, используя табличный подход.
	testCases := []struct {
		name             string
		n                int
		expectedVal      int
		expectedErr      bool   // true, если ожидается ошибка, иначе false
		errorMsgContains string // Опционально: подстрока для проверки в тексте ошибки
	}{
		{"Факториал нуля", 0, 1, false, ""},
		{"Факториал единицы", 1, 1, false, ""},
		{"Факториал пяти", 5, 120, false, ""},
		{"Факториал отрицательного", -1, 0, true, "факториал отрицательного числа не определен"}, // Ожидаем ошибку
		{"Факториал десяти", 10, 3628800, false, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// TODO: Вызовите Factorial(tc.n)
			result, err := Factorial(tc.n)
			if tc.expectedErr {
				if err == nil {
					t.Fatalf("ожидали ошибку, но получили nil")
				}
				// Опционально: проверьте текст ошибки, если нужно
				if !strings.Contains(err.Error(), tc.errorMsgContains) {
					t.Fatalf("Ожидалась ошибка '%s', получена '%s'", tc.errorMsgContains, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("не ожидали ошибку, но получили: %v", err)
				}
				// Если ошибки не ожидалось, сравните результат (result) с ожидаемым (tc.expectedVal)
				if result != tc.expectedVal {
					t.Errorf("ожидали %d, получили %d", tc.expectedVal, result)
				}
			}
		})
	}
}
