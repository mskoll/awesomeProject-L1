**1. Какой самый эффективный способ конкатенации строк?**

В Go строки представляют собой неизменяемый последовательный набор байт фиксированного размера.
При любом изменении строки - создается новая строка, что может привести к проблемам с памятью.
Наиболее эффективным считается использование strings.Builder

[7 Ways to Concatenate Strings in Golang](https://golangdocs.com/concatenate-strings-in-golang#:~:text=6.)


**2. Что такое интерфейсы, как они применяются в Go?**

Интерфейс — это абстрактный тип, который описывает поведение, но не реализовывает его.
Интерфейсы описывают абстракцию (обобщают) поведение других типов. Они не реализовывают, а лишь определяют набор
методов, которые структуры данного типа должны реализовать.
Структура соответсвует интерфейсу, если обладает всеми методами, перечисленными в интерфейсе.
В Go интерфейс реализуется неявно. Для реализации, структуре достаточно реализовать методы, которые определяет интерфейс.

В Go также можно использовать пустой интерфейс (interface{}). Это значит что можно использовать объект любого типа.
Например:
```go
    person := make(map[string]interface{}, 0)
    person["name"] = "Alice"
    person["age"] = 21
```


**3. Чем отличаются RWMutex от Mutex?**

RWMutex предоставляет дополнительные методы RLock, RUnlock, поторые позволяют управлять блокирокой чтения ресурсов


**4. Чем отличаются буферизированные и не буферизированные каналы?**

Каналы - инструмент коммуникации между горутинами.

Для создания небуферизированного канала вызывается функция make() без указания емкости канала
Если канал пустой, то горутина-получатель блокируется, пока в канале не окажутся данные.
Когда горутина-отправитель посылает данные, горутина-получатель получает эти данные и возобновляет работу.
Горутина-отправитель может отправлять данные только в пустой канал.
Горутина-отправитель блокируется до тех пор, пока данные из канала не будут получены.

Буферизированные каналы также создаются с помощью функции make(), только в качестве второго аргумента в функцию передается емкость канала.
Если канал пуст, то получатель ждет, пока в канале появится хотя бы один элемент.
При отправке данных горутина-отправитель ожидает, пока в канале не освободится место для еще одного элемента и
отправляет элемент, только тогда, когда в канале освобождается для него место.
Пока в канале есть место - горутина не блокируется.


**5. Какой размер у структуры struct{}{}?**

Такая структура занимает 0 байт
Это можно проверить с помощью:
```go
st := struct{}{}
fmt.Printf("Struct size: %v", unsafe.Sizeof(st))
```


**6. Есть ли в Go перегрузка методов или операторов?**

Нет
Чтобы избежать путаницы и уязвимости из-за большого количества методов с одинаковыми именами,
но разными сигнатурами, в Go нет перегрузки методов/операторов.

[Frequently Asked Questions (FAQ)](https://golang.org/doc/faq#overloading)


**7. В какой последовательности будут выведены элементы map[int]int?
Пример:**
```go
m[0]=1
m[1]=124
m[2]=281
```

Неизвестно. В хэш-таблицах данные хранятся неупорядоченно, поэтому при каждом выводе порядок может быть разным.


**8. В чем разница make и new?**

new(T) выделяет память память под объект типа Т и возвращает его "адрес" - указатель *Т
Значение объекта инициализируется "нулем"

make используется только для выделения и инициализации объектов типов срезов, мап и каналов.

[new](https://go.dev/doc/effective_go#allocation_new)
[make](https://go.dev/doc/effective_go#allocation_make)


**9. Сколько существует способов задать переменную типа slice или map?**

Slice: 4

Используя []<type>{}:
```go
s := []int // s := []int{1,2}
```

Используя фрагмент другого среза или массива:
```go
numbers := [5]int{1, 2, 3, 4, 5} // numbers := []int{1, 2, 3, 4, 5}
num1 := numbers[2:4]
```

make():
```go
numbers := make([]int, 5)
```

new():
```go
numbers := new([]int)
```

Map: 2

Используя map[<key_type>]<value_type>:
```go
employeeSalary := map[string]int{}
```

make():
```go
employeeSalary := make(map[string]int)
```


**10. Что выведет данная программа и почему?**
```go
func update(p *int) {
	b := 2
	p = &b
}
func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}
```
Выведется 1, 1
Так выводится потому, что в функциию приходит копия адреса переменной a.
Чтобы программа работала корректно - нужно разыменовать переменную р и уже в нее записать значение
*p = b
Стоит учитывать, что при этом значение переменной а тоже изменится


**11. Что выведет данная программа и почему?**
```go
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) { //go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i) // }(&wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
```
Выведутся числа от 0 до 4, в случайном порядке и произойдет дедлок.
Так происходит из-за того что в горутину WG передается по значению, а не по ссылке
WG копируется и вызов wg.Done() не изменяет счетчик wg в мэйне
Чтобы программа работала корректно - в горутину нужно передать указатель на wg


**12. Что выведет данная программа и почему?**
```go
func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
```
Эта программа выведет 0, т.к. n := 0 и n := 1 создаются в разных областях видимости
Если бы fmt.Println(n) находился внутри блока if, либо вместо n := 1 было n = 1, то вывелось бы значение 2


**13. Что выведет данная программа и почему?**
```go
func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}
func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
```
Программа выведет [100 2 3 4 5]
В данной программе срез передается по значению, т.е. в функию попадает копия среза, но с тем же указателем.
В свою очередь, функция append возвращает новый обновленный срез.


14. Что выведет данная программа и почему?
```go
func main() {
	slice := []string{"a", "a"}
	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
```
Программа выведет [b b a][a a]
Функция append возвращает новый срез, поэтому все изменения, происходящие после добавления элемента в срез,
не затрагивают оригинальный срез.

