package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Yeralkhan06/Waterbill/waterbill"
	"github.com/google/uuid" // Добавленный внешний пакет
)

func main() {
	// Входные данные
	prevReading := 120.5
	currReading := 145.3
	tariff := 35.50
	penalty := 10.0 // 10% штраф
	owner := "Иванов И.И."

	// Генерируем уникальный ID для квитанции
	receiptID := uuid.New().String()
	receiptShortID := strings.Split(receiptID, "-")[0] // Берем первую часть для краткости

	fmt.Println("=== Демонстрация работы пакета waterbill ===\n")
	fmt.Printf("Квитанция №: %s\n", receiptShortID)
	fmt.Printf("Полный UUID: %s\n\n", receiptID)

	// 1. Вызов F1 (WaterUsage) - вычислительная функция
	fmt.Println("1. Расчет объема потребления:")
	usage, err := waterbill.WaterUsage(prevReading, currReading)
	if err != nil {
		log.Fatalf("Ошибка при расчете объема: %v", err)
	}
	fmt.Printf("Предыдущие показания: %.2f\n", prevReading)
	fmt.Printf("Текущие показания: %.2f\n", currReading)
	fmt.Printf("Объем потребления: %.3f куб. м\n\n", usage)

	// 2. Вызов F1 (WaterCost) - еще одна вычислительная функция
	fmt.Println("2. Расчет стоимости без штрафа:")
	cost, err := waterbill.WaterCost(usage, tariff)
	if err != nil {
		log.Fatalf("Ошибка при расчете стоимости: %v", err)
	}
	fmt.Printf("Тариф: %.2f руб./куб. м\n", tariff)
	fmt.Printf("Стоимость без штрафа: %.2f руб.\n\n", cost)

	// 3. Вызов F2 (ApplyPenalty) - функция с указателем
	fmt.Println("3. Применение штрафа:")
	costWithPenalty := cost
	err = waterbill.ApplyPenalty(&costWithPenalty, penalty)
	if err != nil {
		log.Fatalf("Ошибка при применении штрафа: %v", err)
	}
	fmt.Printf("Штраф: %.0f%%\n", penalty)
	fmt.Printf("Стоимость до штрафа: %.2f руб.\n", cost)
	fmt.Printf("Стоимость после штрафа: %.2f руб.\n\n", costWithPenalty)

	// 4. Вызов F3 (FormatWaterReport) - формирование строки отчета
	fmt.Println("4. Формирование отчета:")
	report, err := waterbill.FormatWaterReport(owner, usage, costWithPenalty)
	if err != nil {
		log.Fatalf("Ошибка при формировании отчета: %v", err)
	}
	fmt.Println(report)

	// Добавляем UUID в отчет
	fmt.Printf("\nЭлектронная подпись: %s\n", receiptID)

	// Демонстрация обработки ошибок
	fmt.Println("\n=== Демонстрация обработки ошибок ===")

	// Попытка с отрицательными показаниями
	fmt.Println("\nПопытка 1: Отрицательные показания")
	_, err = waterbill.WaterUsage(-10, 100)
	if err != nil {
		fmt.Printf("Ошибка перехвачена: %v\n", err)
	}

	// Попытка с некорректным процентом штрафа
	fmt.Println("\nПопытка 2: Штраф > 100%%")
	testCost := 1000.0
	err = waterbill.ApplyPenalty(&testCost, 150)
	if err != nil {
		fmt.Printf("Ошибка перехвачена: %v\n", err)
	}

	// Попытка с пустым именем владельца
	fmt.Println("\nПопытка 3: Пустое имя владельца")
	_, err = waterbill.FormatWaterReport("", 10, 500)
	if err != nil {
		fmt.Printf("Ошибка перехвачена: %v\n", err)
	}
}
