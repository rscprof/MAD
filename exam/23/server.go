package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	ShelfLife int     `json:"shelf_life"`
	IsOrganic bool    `json:"is_organic"`
}

var products = []Product{
	{"1", "Яблоко", 120.5, 30, true},
	{"2", "Груша", 110.75, 25, false},
	{"3", "Банан", 50.3, 7, true},
	{"4", "Апельсин", 90.25, 15, false},
	{"5", "Персик", 150.0, 20, true},
	{"6", "Киви", 160.0, 10, true},
	{"7", "Вишня", 200.0, 15, false},
	{"8", "Гранат", 250.0, 30, true},
	{"9", "Арбуз", 60.0, 5, false},
	{"10", "Малина", 180.0, 7, true},
	{"11", "Клубника", 250.0, 10, false},
	{"12", "Черешня", 220.0, 12, true},
	{"13", "Черника", 210.0, 14, false},
	{"14", "Ежевика", 300.0, 18, true},
	{"15", "Мандарин", 120.0, 30, false},
	{"16", "Лимон", 70.0, 20, true},
	{"17", "Болгарский перец", 80.0, 10, false},
	{"18", "Помидор", 60.0, 7, true},
	{"19", "Огурец", 40.0, 5, false},
	{"20", "Картофель", 30.0, 60, true},
	{"21", "Морковь", 40.0, 40, false},
	{"22", "Свекла", 50.0, 50, true},
	{"23", "Лук", 30.0, 90, false},
	{"24", "Чеснок", 60.0, 180, true},
	{"25", "Брокколи", 100.0, 7, false},
	{"26", "Цветная капуста", 120.0, 10, true},
	{"27", "Шпинат", 90.0, 5, false},
	{"28", "Руккола", 150.0, 5, true},
	{"29", "Кукуруза", 80.0, 50, false},
	{"30", "Горошек", 70.0, 30, true},
	{"31", "Кабачок", 60.0, 20, false},
	{"32", "Тыква", 110.0, 60, true},
	{"33", "Ананас", 200.0, 15, false},
	{"34", "Манго", 300.0, 20, true},
	{"35", "Кокос", 350.0, 60, false},
	{"36", "Арбуз", 80.0, 7, true},
	{"37", "Болгарский перец", 75.0, 10, false},
	{"38", "Капуста", 40.0, 60, true},
	{"39", "Редис", 30.0, 7, false},
	{"40", "Петрушка", 20.0, 14, true},
	{"41", "Укроп", 25.0, 10, false},
	{"42", "Мята", 50.0, 5, true},
	{"43", "Сельдерей", 90.0, 15, false},
	{"44", "Огурец", 45.0, 6, true},
	{"45", "Помидоры", 55.0, 12, false},
	{"46", "Перец чили", 80.0, 30, true},
	{"47", "Шампиньоны", 150.0, 7, false},
	{"48", "Лисички", 200.0, 5, true},
	{"49", "Боровики", 250.0, 6, false},
	{"50", "Маслята", 170.0, 8, true},
	{"51", "Трюфели", 1000.0, 1, false},
	{"52", "Сморчки", 400.0, 3, true},
	{"53", "Сливки", 200.0, 7, false},
	{"54", "Молоко", 60.0, 14, true},
	{"55", "Сыр", 300.0, 20, false},
	{"56", "Творог", 120.0, 10, true},
	{"57", "Яйца", 50.0, 30, false},
	{"58", "Кефир", 80.0, 14, true},
	{"59", "Крем-сыр", 180.0, 15, false},
	{"60", "Йогурт", 90.0, 20, true},
	{"61", "Масло сливочное", 150.0, 60, false},
	{"62", "Майонез", 80.0, 90, true},
	{"63", "Гречка", 50.0, 365, false},
	{"64", "Рис", 60.0, 180, true},
	{"65", "Пшено", 40.0, 360, false},
	{"66", "Макароны", 30.0, 365, true},
	{"67", "Киноа", 200.0, 180, false},
	{"68", "Чечевица", 100.0, 365, true},
	{"69", "Фасоль", 120.0, 365, false},
	{"70", "Кускус", 150.0, 180, true},
	{"71", "Манка", 40.0, 365, false},
	{"72", "Перловка", 30.0, 360, true},
	{"73", "Ячмень", 50.0, 360, false},
	{"74", "Овсянка", 45.0, 365, true},
	{"75", "Сахар", 60.0, 365, false},
	{"76", "Соль", 20.0, 730, true},
	{"77", "Перец черный", 70.0, 180, false},
	{"78", "Чесночный порошок", 90.0, 365, true},
	{"79", "Тмин", 100.0, 180, false},
	{"80", "Корица", 120.0, 365, true},
	{"81", "Лавровый лист", 60.0, 180, false},
	{"82", "Гвоздика", 150.0, 365, true},
	{"83", "Карри", 80.0, 180, false},
	{"84", "Имбирь", 90.0, 180, true},
	{"85", "Куркума", 70.0, 180, false},
	{"86", "Саус", 100.0, 30, true},
	{"87", "Кетчуп", 90.0, 90, false},
	{"88", "Масло оливковое", 350.0, 730, true},
	{"89", "Масло растительное", 120.0, 365, false},
	{"90", "Соевый соус", 180.0, 180, true},
	{"91", "Кокосовое масло", 500.0, 730, false},
	{"92", "Горчица", 70.0, 180, true},
	{"93", "Песто", 150.0, 30, false},
	{"94", "Томатная паста", 80.0, 180, true},
	{"95", "Соус барбекю", 100.0, 180, false},
	{"96", "Вино красное", 700.0, 730, true},
	{"97", "Вино белое", 600.0, 730, false},
	{"98", "Пиво", 120.0, 180, true},
	{"99", "Виски", 1500.0, 730, false},
	{"100", "Ром", 1000.0, 730, true},
}

func main() {
	http.HandleFunc("/products", getProductsByShelfLife)

	port := ":9080"
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// Функция для поиска продуктов по сроку годности
func getProductsByShelfLife(w http.ResponseWriter, r *http.Request) {
	minShelfLife := r.URL.Query().Get("min_shelf_life")
	maxShelfLife := r.URL.Query().Get("max_shelf_life")

	if minShelfLife == "" || maxShelfLife == "" {
		http.Error(w, "Both min_shelf_life and max_shelf_life query parameters are required", http.StatusBadRequest)
		return
	}

	// Преобразуем строковые значения в целые числа
	minShelfLifeVal, err := strconv.Atoi(minShelfLife)
	if err != nil {
		http.Error(w, "Invalid min_shelf_life value", http.StatusBadRequest)
		return
	}

	maxShelfLifeVal, err := strconv.Atoi(maxShelfLife)
	if err != nil {
		http.Error(w, "Invalid max_shelf_life value", http.StatusBadRequest)
		return
	}

	// Ищем продукты в указанном интервале срока годности
	var result []Product
	for _, product := range products {
		if product.ShelfLife >= minShelfLifeVal && product.ShelfLife <= maxShelfLifeVal {
			result = append(result, product)
		}
	}

	// Если продукты не найдены
	if len(result) == 0 {
		http.Error(w, "No products found within the specified shelf life range", http.StatusNotFound)
		return
	}

	// Отправляем результат в виде JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
