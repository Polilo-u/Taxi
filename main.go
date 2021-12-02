package main

import (
	"Tax/FreTax"
	jsonTax "Tax/Json"
	"Tax/PasTax"
	"Tax/Passenger"
	"Tax/Sql"
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	//_ "github.com/go-sql-driver/mysql"
)

type Taxi interface { //Интерфейс Taxi для обхединения PasTax и FreTax
	Price_count()
	Get() (string, string, int)
}

var str []interface{}

var Taxs []Taxi //Создание слайса со всеми Taxi

func sliceRemoveItem(slicep interface{}, i int) { //Функция для удаления элемента из слайса
	v := reflect.ValueOf(slicep).Elem()
	v.Set(reflect.AppendSlice(v.Slice(0, i), v.Slice(i+1, v.Len())))
}

func main() {
	url := "https://postman-echo.com/post"
	//Создание легковых автомобилей
	pastax1 := PasTax.Pas_Tax{}
	pastax2 := PasTax.Pas_Tax{}
	//Объявление значений у легковых такси
	pastax1.Set("Аскар", "КА127Л")
	pastax2.Set("Мухамед", "РА756Е")

	//Создание грузовых автомобилей
	fretax1 := FreTax.Fre_Tax{}
	fretax2 := FreTax.Fre_Tax{}

	//Объявление значений у грузовых такси
	fretax1.Set("Жылкызбек", "ДЛ654K")
	fretax2.Set("Алексей", "ТИ423П")

	//Добавление всех такси в слайс
	Taxs = append(Taxs, &pastax1, &pastax2, &fretax1, &fretax2)

	// Создание пассажиров
	pas1 := Passenger.Passenger{}
	pas2 := Passenger.Passenger{}
	pas3 := Passenger.Passenger{}
	pas4 := Passenger.Passenger{}
	pas5 := Passenger.Passenger{}
	//Объявление значений у каждого пассажира
	pas1.Set(12, 13, 2, "Vasya")
	pas2.Set(120, 20, 1, "Petya")
	pas3.Set(50, 20, 2, "Slava")
	pas4.Set(21, 80, 1, "Masha")
	pas5.Set(19, 50, 1, "Gleb")

	//Добавление пассажиров в слайс
	Passenger.Passengers = append(Passenger.Passengers, pas1, pas2, pas3, pas4, pas5)
	//Создание вспомогательной переменной
	a := len(Taxs)

	for _, element := range Passenger.Passengers { // Цикл, который пробегает по слайсу Пасажиров
		kg, km, kol_vo, namepas := element.Get()              //Получение данных из пассажира
		if len(Taxs) != 0 && len(Passenger.Passengers) != 0 { //Проверка на то, что слайсы не пусты
			if kg < 50 && kol_vo <= 4 { // Удовлятворяют ли запросы пассажира легковому такси
				for i, value := range Taxs { // Цикл, который пробегает по слайсу Такси

					value.Price_count() //Расчёт цены
					name, number, kmch := value.Get()
					fmt.Println("К ", namepas, "подъедет:", name, "Номер автомобиля:", number, "Цена составит:", kmch*km)
					str = append(str, name, number, kmch*km)
					byt := jsonTax.Marsh(str)
					req, err := http.NewRequest("POST", url, bytes.NewBuffer(byt))
					if err != nil {
						panic(err)
					}
					taxis := &http.Client{}
					resp, err := taxis.Do(req)
					if err != nil {
						panic(err)
					}
					defer resp.Body.Close()
					fmt.Println("response Status:", resp.Status)
					dat := jsonTax.UnMarsh(byt)
					fmt.Println(dat)
					Sql.AdedSql(name, number, kmch*km) //Добавление в БД
					sliceRemoveItem(&Taxs, i)          // Удаление элемента из Такси
					break                              //Break цикла Taxi
				}
			} else if kol_vo <= 2 { // Если запросы пассажира не удовлетворяют легковому такси, то проверка на грузовое
				for i, value := range Taxs { //Пробегаем по слайсу такси
					value.Price_count()
					name, number, kmch := value.Get()
					fmt.Println("К ", namepas, "подъедет:", name, "Номер автомобиля:", number, "Цена составит:", kmch*km)
					str = append(str, name, number, kmch*km)
					byt := jsonTax.Marsh(str)
					req, err := http.NewRequest("POST", url, bytes.NewBuffer(byt))
					if err != nil {
						panic(err)
					}
					taxis := &http.Client{}
					resp, err := taxis.Do(req)
					if err != nil {
						panic(err)
					}
					defer resp.Body.Close()
					fmt.Println("response Status:", resp.Status)
					dat := jsonTax.UnMarsh(byt)
					fmt.Println(dat)
					Sql.AdedSql(name, number, kmch*km)
					sliceRemoveItem(&Taxs, i)
					break
				}
			} else { //Если запрос пассажира не удовлетворяет ни грузовому, ни легковому такси
				fmt.Println("К сожалению у нас нет свободного автомобиля под ваш запрос")
			}
		}
	}
	if len(Taxs) == 0 && len(Passenger.Passengers) != 0 {
		for i, element := range Passenger.Passengers {
			_, _, _, name := element.Get()
			if i >= a {
				fmt.Println("Идет поиск автомобиля для ", name)
			}
		}
	}
}
