package Passenger

type Passenger struct { //Структура пасажир
	kg     int
	km     int
	kol_vo int
	name   string
}

func (p *Passenger) Set(kg, km, kol_vo int, name string) { //Функция для установки значений в пасажира
	p.kg = kg
	p.km = km
	p.kol_vo = kol_vo
	p.name = name
}

func (p *Passenger) Get() (int, int, int, string) { //Функция для получения значений из пасажира
	return p.kg, p.km, p.kol_vo, p.name
}

var Passengers []Passenger //Слайс с пасажирами
