package PasTax

import (
	"Tax/Passenger" //подключение пакета Passtnger для использования слайса Passengers
)

type Pas_Tax struct { //Структура пасажирское такси
	name   string
	number string
	price  int
}

func (p *Pas_Tax) Get() (string, string, int) { //Функция для получения значений из структуры легкового такси
	return p.name, p.number, p.price
}

func (p *Pas_Tax) Set(a, b string) { //Функция для записи значений в структуру легкового такси
	p.name = a
	p.number = b
}

func (p *Pas_Tax) Price_count() { //расчет цены для легкового такс
	if len(Passenger.Passengers) > 3 {
		p.price = 25
	} else if len(Passenger.Passengers) <= 3 && len(Passenger.Passengers) > 1 {
		p.price = 20
	} else {
		p.price = 15
	}
}
