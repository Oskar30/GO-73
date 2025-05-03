package human

import (
	"fmt"
	"math/rand"
)

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func NewMan(name, last, gen string, age, crim int) Man {
	return Man{
		Name:     name,
		LastName: last,
		Age:      age,
		Gender:   gen,
		Crimes:   crim,
	}
}

func (m Man) Stringer() string {
	return fmt.Sprintf("Досье подозреваемого:\n Фамилия: %s\n Имя: %s\n Пол: %s\n Возраст: %d\n Количество совершённых преступлений: %d\n", m.LastName, m.Name, m.Gender, m.Age, m.Crimes)
}

func RandMan(m map[string]Man) map[string]Man {
	nameM := []string{"Иван", "Пётр", "Михаил", "Александр", "Евгений", "Артём", "Сергей", "Дмитрий", "Андрей", "Максим"}
	lastnameM := [10]string{"Иванов", "Петров", "Смирнов", "Кузнецов", "Зайцев", "Попов", "Романов", "Морозов", "Орлов", "Белов"}

	nameW := []string{"Евгения", "Юлия", "Кира", "Анна", "Стефания", "Елизавета", "Екатерина", "Ирина", "Алиса", "Пелагея"}
	lastnameW := [10]string{"Макаровна", "Трошина", "Пантелеева", "Петрова", "Гусева", "Николаева", "Чернышева", "Федорова", "Леонова", "Чистякова"}

	for i := rand.Intn(5) + 5; i > 0; i-- {
		if rand.Intn(2) > 0 {
			m[nameM[i]] = NewMan(nameM[i], lastnameM[rand.Intn(len(lastnameM))], "M", rand.Intn(72)+18, rand.Intn(4)+1)
		} else {
			m[nameM[i]] = NewMan(nameW[i], lastnameW[rand.Intn(len(lastnameM))], "Ж", rand.Intn(72)+18, rand.Intn(4)+1)
		}
	}
	return m
}
