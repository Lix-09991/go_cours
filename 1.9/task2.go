package main

import (
	"fmt"
)

// Структура Student
type Student struct {
	Name   string
	Grades []float64
}

// Метод AverageGrade
func (s Student) AverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0.0
	for _, grade := range s.Grades {
		sum += grade
	}
	return sum / float64(len(s.Grades))
}

func main() {
	// Создаем студента и оценки
	student := Student{
		Name:   "Иван Иванов",
		Grades: []float64{4.5, 3.7, 4.2, 5.0},
	}

	// Выводим средний балл
	fmt.Printf("Средний балл студента %s: %.2f\n", student.Name, student.AverageGrade())
}
