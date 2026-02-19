package waterbill

// Пакет waterbill предоставляет функции для расчета стоимости потребления воды
// и формирования отчетов. Содержит функции для вычисления объема потребления,
// стоимости, применения штрафов и форматирования итогового отчета.

import (
	"fmt"
)

// WaterUsage вычисляет объем потребленной воды на основе предыдущих (prev)
// и текущих (curr) показаний счетчика.
// Возвращает ошибку, если показания отрицательные или текущее показание
// меньше предыдущего (что невозможно при нормальной работе счетчика).
func WaterUsage(prev, curr float64) (float64, error) {
	if prev < 0 || curr < 0 {
		return 0, fmt.Errorf("показания счетчика не могут быть отрицательными: prev=%.2f, curr=%.2f", prev, curr)
	}
	if curr < prev {
		return 0, fmt.Errorf("текущее показание (%.2f) не может быть меньше предыдущего (%.2f)", curr, prev)
	}
	return curr - prev, nil
}

// WaterCost вычисляет стоимость потребленной воды на основе объема (cubic)
// и тарифа (tariff) за единицу объема.
// Возвращает ошибку, если объем или тариф отрицательные.
func WaterCost(cubic, tariff float64) (float64, error) {
	if cubic < 0 {
		return 0, fmt.Errorf("объем потребления не может быть отрицательным: %.2f", cubic)
	}
	if tariff < 0 {
		return 0, fmt.Errorf("тариф не может быть отрицательным: %.2f", tariff)
	}
	return cubic * tariff, nil
}

// ApplyPenalty применяет штраф в размере penaltyPercent процентов к стоимости,
// на которую указывает cost. Изменяет значение по указателю.
// Возвращает ошибку, если cost равен nil, или если процент штрафа
// находится за пределами диапазона (0, 100] (штраф должен быть положительным и не превышать 100%).
func ApplyPenalty(cost *float64, penaltyPercent float64) error {
	if cost == nil {
		return fmt.Errorf("указатель на стоимость не может быть nil")
	}
	if penaltyPercent <= 0 {
		return fmt.Errorf("процент штрафа должен быть положительным: %.2f", penaltyPercent)
	}
	if penaltyPercent > 100 {
		return fmt.Errorf("процент штрафа не может превышать 100: %.2f", penaltyPercent)
	}
	*cost = *cost * (1 + penaltyPercent/100)
	return nil
}

// FormatWaterReport формирует строку отчета о потреблении воды.
// Принимает имя владельца (owner), объем потребления (cubic) и стоимость (cost).
// Возвращает отформатированную строку или ошибку, если имя владельца пустое,
// или если объем/стоимость отрицательные.
func FormatWaterReport(owner string, cubic, cost float64) (string, error) {
	if owner == "" {
		return "", fmt.Errorf("имя владельца не может быть пустым")
	}
	if cubic < 0 {
		return "", fmt.Errorf("объем потребления не может быть отрицательным: %.2f", cubic)
	}
	if cost < 0 {
		return "", fmt.Errorf("стоимость не может быть отрицательной: %.2f", cost)
	}

	report := fmt.Sprintf("Отчет по воде для %s\n", owner)
	report += fmt.Sprintf("Объем потребления: %.2f куб. м\n", cubic)
	report += fmt.Sprintf("Итого к оплате: %.2f руб.", cost)
	return report, nil
}
