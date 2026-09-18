// Лабораторная работа №1
// Тема: «Типы данных в Go»
// Программный прототип: работа с композитным типом Employee ,срезом структур, фильтрацией и сортировкой.

package main

import (
	"fmt"
	"sort"
)

// Employee — композитный тип данных (структура), описывающий сотрудника.
type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

// ToString возвращает строковое представление сотрудника в удобном для чтения виде.
func (e Employee) ToString() string {
	return fmt.Sprintf("ID: %-3d | %-18s | Отдел: %-12s | Зарплата: %8.2f",
		e.ID, e.Name, e.Department, e.Salary)
}

// AddEmployee добавляет нового сотрудника в срез и возвращает обновлённый срез.
//  append может выделить новую область памяти при переполнении ёмкости, поэтому результат  присваиваиваем обратно.

func AddEmployee(employees []Employee, e Employee) []Employee {
	return append(employees, e)
}

// FilterByDepartment возвращает новый срез, содержащий только сотрудников указанного отдела. Исходный срез не изменяется.
func FilterByDepartment(employees []Employee, department string) []Employee {
	result := make([]Employee, 0)
	for _, e := range employees {
		if e.Department == department {
			result = append(result, e)
		}
	}
	return result
}

// AverageSalary считает среднюю зарплату по срезу сотрудников.
func AverageSalary(employees []Employee) float64 {
	if len(employees) == 0 {
		return 0
	}
	var total float64
	for _, e := range employees {
		total += e.Salary
	}
	return total / float64(len(employees))
}

// FilterAboveAverageSalary возвращает сотрудников, чья зарплата выше средней
func FilterAboveAverageSalary(employees []Employee) []Employee {
	avg := AverageSalary(employees)
	result := make([]Employee, 0)
	for _, e := range employees {
		if e.Salary > avg {
			result = append(result, e)
		}
	}
	return result
}

// SortBySalaryDesc сортирует срез сотрудников по убыванию зарплаты.
// Функция изменяет исходный срез на месте (sort.Slice работаетс переданным слайсом напрямую, без копирования).
func SortBySalaryDesc(employees []Employee) {
	sort.Slice(employees, func(i, j int) bool {
		return employees[i].Salary > employees[j].Salary
	})
}

// SortByName сортирует срез сотрудников по имени в алфавитном порядке.
func SortByName(employees []Employee) {
	sort.Slice(employees, func(i, j int) bool {
		return employees[i].Name < employees[j].Name
	})
}

// printEmployees выводит список сотрудников построчно, используя метод ToString.
func printEmployees(employees []Employee) {
	if len(employees) == 0 {
		fmt.Println("  (список пуст)")
		return
	}
	for _, e := range employees {
		fmt.Println(" ", e.ToString())
	}
}

func main() {
	// 1. Создание среза структур и заполнение через AddEmployee
	var employees []Employee
	employees = AddEmployee(employees, Employee{ID: 1, Name: "Иванов И.И.", Department: "IT", Salary: 85000})
	employees = AddEmployee(employees, Employee{ID: 2, Name: "Петрова А.С.", Department: "IT", Salary: 95000})
	employees = AddEmployee(employees, Employee{ID: 3, Name: "Сидоров К.П.", Department: "Бухгалтерия", Salary: 60000})
	employees = AddEmployee(employees, Employee{ID: 4, Name: "Кузнецова М.В.", Department: "HR", Salary: 55000})
	employees = AddEmployee(employees, Employee{ID: 5, Name: "Смирнов Д.А.", Department: "IT", Salary: 120000})

	fmt.Println("=== Полный список сотрудников ===")
	printEmployees(employees)

	fmt.Println("\n=== Сотрудники отдела IT ===")
	itEmployees := FilterByDepartment(employees, "IT")
	printEmployees(itEmployees)

	fmt.Printf("\n=== Средняя зарплата по всем сотрудникам: %.2f ===\n", AverageSalary(employees))

	fmt.Println("\n=== Сотрудники с зарплатой выше средней ===")
	aboveAvg := FilterAboveAverageSalary(employees)
	printEmployees(aboveAvg)

	fmt.Println("\n=== Список, отсортированный по убыванию зарплаты ===")
	sortedBySalary := make([]Employee, len(employees))
	copy(sortedBySalary, employees)
	SortBySalaryDesc(sortedBySalary)
	printEmployees(sortedBySalary)

	fmt.Println("\n=== Список, отсортированный по имени ===")
	sortedByName := make([]Employee, len(employees))
	copy(sortedByName, employees)
	SortByName(sortedByName)
	printEmployees(sortedByName)
}
