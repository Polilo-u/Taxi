package FreTax

import (
	"Tax/Passenger" //подключение пакета Passenger для использования слайса Passengers
)

type Fre_Tax struct { //Структура грузовое такси
	name   string
	number string
	price  int
}

func (f *Fre_Tax) Get() (string, string, int) { //Функция для получения значений из структуры грузового такси
	return f.name, f.number, f.price
}

func (f *Fre_Tax) Set(a, b string) { //Функция для установки значений в структуру грузового такси
	f.name = a
	f.number = b
}

func (f *Fre_Tax) Price_count() { //Функция для расчета цены у грузового такси
	if len(Passenger.Passengers) > 3 {
		f.price = 30
	} else if len(Passenger.Passengers) <= 3 && len(Passenger.Passengers) > 1 {
		f.price = 23
	} else {
		f.price = 19
	}
}
