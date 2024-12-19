package main

import (
	"fmt"
	"math"
)

func countInRange(numbers []float64, min, max float64) (count int) {
	count = 0
	for _, num := range numbers {
		if num >= min && num <= max {
			count++
		}
	}
	return count
}

func findMin(numbers []float64) float64 {
	if len(numbers) == 0 {
		return math.Inf(1)
	}
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

func findMax(numbers []float64) float64 {
	if len(numbers) == 0 {
		return math.Inf(-1)
	}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

type Tuple struct {
	GranInterval     []float64
	Average          []float64
	List             string
	Chast            int
	OtnosChast       float64
	NakoplChast      int
	NakoplOtnosChast float64
}

func Itog(mas []float64, g int) (tuple []Tuple) {
	len := len(mas)
	minX := findMin(mas)
	maxX := findMax(mas)
	deltaX := maxX - minX
	interval := math.Ceil(deltaX / float64(g))

	min := minX - 0.5
	max := min + interval
	nakoplChast := 0
	nakoplOtnosChast := 0.0
	for i := 0; i < g; i++ {
		if i != 0 {
			min = max
			max = min + interval
		}

		average := (max + min) / 2
		count := countInRange(mas, min, max)
		otnosChast := float64(count) / float64(len)
		nakoplChast += count
		nakoplOtnosChast += otnosChast

		bars := ""
		for j := 0; j < count; j++ {
			if j > 0 && j%5 == 0 {
				bars += " "
			}
			bars += "|"
		}

		list := bars

		tuple = append(tuple, Tuple{
			GranInterval:     []float64{min, max},
			Average:          []float64{average},
			List:             list,
			Chast:            count,
			OtnosChast:       otnosChast,
			NakoplChast:      nakoplChast,
			NakoplOtnosChast: nakoplOtnosChast,
		})
	}

	return tuple
}

type Tuple1 struct {
	Average []float64
	Mi      float64
	Ui      float64
	MiUi    float64
	MiUiKV  float64
	Control float64
}

func Itog1(c float64, k float64, mainTuple []Tuple) (tuple1 []Tuple1) {
	for _, mt := range mainTuple {
		tuple1 = append(tuple1, Tuple1{
			Average: mt.Average,
			Mi:      float64(mt.Chast),
			Ui:      (mt.Average[0] - c) / k,
			MiUi:    float64(mt.Chast) * (mt.Average[0] - c) / k,
			MiUiKV:  float64(mt.Chast) * ((mt.Average[0] - c) / k) * ((mt.Average[0] - c) / k),
			Control: float64(mt.Chast) * math.Pow((((mt.Average[0]-c)/k)+1), 2),
		})
	}

	return tuple1
}

type TeorChast struct {
	IntervalGrup []float64
	Pi           float64
	MiTeor       float64
	Person       float64
}

func TeorChastFunc(tup []Tuple, sg float64, n float64) (teorChast []TeorChast) {
	for _, t := range tup {
		// Вычисляем вероятности с использованием нормального распределения
		p1 := math.Exp(-math.Pow(t.GranInterval[0], 2) / (2 * math.Pow(sg, 2)))
		p2 := math.Exp(-math.Pow(t.GranInterval[1], 2) / (2 * math.Pow(sg, 2)))

		// Delta P (вероятность попадания в интервал)
		Pi := p1 - p2

		// Теоретическая частота
		MiTeor := n * Pi

		person := math.Pow((float64(t.Chast)-MiTeor), 2) / MiTeor

		teorChast = append(teorChast, TeorChast{
			IntervalGrup: t.GranInterval,
			Pi:           Pi,
			MiTeor:       MiTeor,
			Person:       person,
		})
	}
	return teorChast
}

func main() {
	massive := []float64{26.7, 94.2, 74.8, 88.7, 93.2, 78.7, 90.5, 73.3, 76.3, 71.9, 80.3, 27.3,
		73.3, 69.8, 69.1, 81.9, 67.7, 57.7, 68.4, 96.1, 67.0, 64.4, 92.3, 67.0,
		39.9, 53.8, 79.5, 74.1, 63.8, 77.1, 86.9, 87.8, 81.1, 61.3, 97.0, 5.5,
		41.5, 48.7, 95.1, 71.2, 58.3, 53.3, 49.2, 55.4, 50.7, 47.7, 52.7, 60.0,
		13.5, 50.2, 77.9, 60.6, 45.4, 98.0, 100.0, 72.6, 44.9, 59.5, 56.5, 56.0,
		16.5, 42.7, 70.5, 43.2, 41.9, 85.2, 38.7, 48.2, 39.1, 44.5, 9.5, 39.5,
		26.1, 49.7, 99.0, 45.8, 40.3, 82.7, 86.1, 51.7, 83.5, 43.6, 52.2, 51.2,
		22.3, 30.2, 89.6, 39.9, 33.3, 91.4, 38.3, 26.2, 37.5, 36.8, 28.3, 37.9,
		65.0, 13.5, 84.4, 27.3, 24.7, 66.4, 58.9, 54.9, 46.8, 61.9, 47.2, 65.7,
		30.0, 42.3, 75.6, 63.1, 62.5, 40.7, 41.1, 46.3, 44.0, 37.2, 57.1, 54.9}

	// data := []float64{
	// 	95.0, 93.0, 89.0, 100.0, 94.0, 95.0, 94.0, 101.0, 90.0, 95.0,
	// 	103.0, 98.0, 99.0, 91.0, 95.0, 94.0, 95.0, 94.0, 89.0, 93.0,
	// 	98.0, 95.0, 93.0, 89.0, 100.0, 107.0, 100.0, 98.0, 101.0, 97.0,
	// 	90.0, 95.0, 103.0, 98.0, 99.0, 91.0, 94.0, 95.0, 94.0, 89.0,
	// 	93.0, 98.0, 93.0, 96.0, 101.0, 97.0, 102.0, 97.0, 106.0, 101.0,
	// 	96.0, 96.0, 94.0, 100.0, 95.0, 92.0, 93.0, 96.0, 97.0, 98.0,
	// 	99.0, 97.0, 104.0, 101.0, 98.0, 109.0, 98.0, 104.0, 95.0, 100.0,
	// 	102.0, 92.0, 95.0, 99.0, 93.0, 92.0, 97.0, 99.0, 98.0, 102.0,
	// 	98.0, 94.0, 98.0, 97.0, 94.0, 90.0, 95.0, 97.0, 103.0, 100.0,
	// 	97.0, 91.0, 96.0, 108.0, 100.0, 91.0, 93.0, 106.0, 93.0, 97.0,
	// 	93.0, 90.0, 95.0, 97.0, 97.0, 99.0, 93.0, 96.0, 101.0, 96.0,
	// 	100.0, 106.0, 105.0, 94.0, 102.0, 91.0, 94.0, 106.0, 98.0, 100.0,
	// }

	g := 12
	result := Itog(massive, g)

	fmt.Printf("%-25s %-10s %-10s\t\t   %-15s %-15s %-20s %-20s\n", "Границы интервала", "Среднее", "Лист", "Частота (mi)", "Отн. частота", "Накоп. Частота", "Накоп. Отн. Частота")
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------")

	// Выводим строки таблицы
	for _, t := range result {
		granInterval := fmt.Sprintf("[%.2f; %.2f)", t.GranInterval[0], t.GranInterval[1])
		average := fmt.Sprintf("%.2f", t.Average[0])
		fmt.Printf("%-25s %-10s %-25s %-15d %-20.6f %-20d %-20.5f\n",
			granInterval,
			average,
			t.List,
			t.Chast,
			t.OtnosChast,
			t.NakoplChast,
			t.NakoplOtnosChast)
	}

	fmt.Printf("%-10s %-10s %-25s %-10s %-10s %-20s\n", "Среднее", "Частота", "Условные варианты", "MiUi", "MiUi^2", "Mi(Ui+1)^2")
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------")

	result1 := Itog1(41, 8, result)
	// result1 := Itog1(97.50, 2, result)

	sumMi := 0.0
	sumMiUi := 0.0
	sumMiUiKV := 0.0
	control := 0.0
	for _, t1 := range result1 {
		average := t1.Average[0]
		fmt.Printf("%-10f %-12f %-20f %-12f %-12f %-10f\n",
			average,
			t1.Mi,
			t1.Ui,
			t1.MiUi,
			t1.MiUiKV,
			t1.Control,
		)

		sumMi += t1.Mi
		sumMiUi += t1.MiUi
		sumMiUiKV += t1.MiUiKV
		control += t1.Control
	}

	fmt.Printf("Итого:\t   %-12f\t\t\t     %-12f %-12f %-10f\n", sumMi, sumMiUi, sumMiUiKV, control)

	sumPi := 0.0
	sumMiTeor := 0.0
	sumPer := 0.0
	teorChast := TeorChastFunc(result, 46.22413, 120)

	fmt.Printf("%-30s %-10s %-25s %-10s\n", "Интервал", "Pi", "Mi~", "Пирсон")
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------")
	for _, t2 := range teorChast {
		granInterval := fmt.Sprintf("[%.2f; %.2f)", t2.IntervalGrup[0], t2.IntervalGrup[1])
		fmt.Printf("%-25s %-12f %-20.5f %-12.5f\n", granInterval, t2.Pi, t2.MiTeor, t2.Person)
		sumPi += t2.Pi
		sumMiTeor += t2.MiTeor
		sumPer += t2.Person
	}

	fmt.Printf("Итого:\t   %-12f\t\t\t     %-12f %-12f\n", sumPi, sumMiTeor, sumPer)
}

func sumWithForLoop(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
