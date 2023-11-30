package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Название канала ля NATS-STREAMING
var channelName = "WBChannel"

// Идентификаторы кластера и клиента
var clusterID = "WB-cluster"
var clientID = "plan9t-client"

// URL NATS Streaming сервера
var natsURL = "nats://localhost:4222"

// Подключение к NATS Streaming
var sc stan.Conn

// var err error

var MyCache = NewCache()

func init() {
	// Инициализация подключения к NATS Streaming
	var err error
	sc, err = stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL(natsURL),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Успешная инициализация")
}

func main() {

	fmt.Println("PROGRAMM STARTED")

	// Работа с КЭШ

	MyCache.AddOrders(GetOrdersFromPostgreSQL())

	// Запуск HTTP сервера на порту и старт прослушки порта 3333 в горутине
	//go startHTTPServer()

	// Настройка обработчика для POST запросов
	//http.HandleFunc("/forpost", handlePostRequest)

	// Проверка подключения к серверу NATS-streaming
	if sc.NatsConn().IsConnected() {
		fmt.Println("ПОДКЛЮЧЕНО К СЕРВЕРУ NATS-STREAMING")
	} else {
		fmt.Println("Не подключено к серверу NATSSTREAMING")
	}

	// Создание канала (темы) при публикации сообщения
	//err = sc.Publish(channelName, []byte("Канал "+channelName+" был создан"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Подписка на канал и прослушка канала (здесь мы парсим JSON и записываем данные в PostgreSQL
	_, err := createDurableSubscription(sc, channelName, clientID)
	if err != nil {
		log.Fatal(err)
	}

	//subscription, err := sc.Subscribe(channelName, func(msg *stan.Msg) {
	//	fmt.Printf("Получено сообщение из канала '%s': %s\n", channelName, string(msg.Data))
	//
	//	// Вызывай функцию для парсинга JSON и сохранения в базу данных
	//	orderData, err := parseJSON(msg.Data)
	//	if err != nil {
	//		fmt.Println("Ошибка при парсинге JSON:", err)
	//		return
	//	}
	//	fmt.Println("JSON спашрен")
	//
	//	err = SaveToPostgreSQL(orderData)
	//	if err != nil {
	//		fmt.Println("Ошибка при сохранении в PostgreSQL:", err)
	//		return
	//	}
	//	fmt.Println("Данные из темы ", channelName, " сохранены в БД")
	//	// Запись в кэш
	//	MyCache.AddOrder(orderData)
	//	fmt.Println("Данные из темы ", channelName, " записаны в КЭШ")
	//}, stan.DurableName(clientID))
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	// После отработки функции main() отписка и закрытие соединения с сервером NATS-streaming

	// Создание канала для ожидания завершения работы программы
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Ожидание сигнала завершения работы
	select {
	case <-done:
		// Закрытие подключения к NATS Streaming
		//subscription.Unsubscribe()
		sc.Close()

		// вывод кэша
		fmt.Println(MyCache.Orders)
		fmt.Println("Программа завершена.")
	}

}

func createDurableSubscription(sc stan.Conn, channelName, clientID string) (stan.Subscription, error) {
	subscription, err := sc.Subscribe(channelName, func(msg *stan.Msg) {
		fmt.Printf("Получено сообщение из канала '%s': %s\n", channelName, string(msg.Data))

		// Вызывай функцию для парсинга JSON и сохранения в базу данных
		orderData, err := parseJSON(msg.Data)
		if err != nil {
			fmt.Println("Ошибка при парсинге JSON:", err)
			return
		}
		fmt.Println("JSON спашрен")

		err = SaveToPostgreSQL(orderData)
		if err != nil {
			fmt.Println("Ошибка при сохранении в PostgreSQL:", err)
			return
		}
		fmt.Println("Данные из темы ", channelName, " сохранены в БД")
		// Запись в кэш
		MyCache.AddOrder(orderData)
		fmt.Println("Данные из темы ", channelName, " записаны в КЭШ")
	}, stan.DurableName(clientID))

	if err != nil && err == stan.ErrBadSubscription {
		// подписка с таким прочным именем уже существует, отписываемся от нее
		if err := subscription.Unsubscribe(); err != nil {
			return nil, err
		}

		// создаем новую подписку с тем же прочным именем
		subscription, err = sc.Subscribe(channelName, func(msg *stan.Msg) {
			fmt.Printf("Получено сообщение из канала '%s': %s\n", channelName, string(msg.Data))

			// Вызывай функцию для парсинга JSON и сохранения в базу данных
			orderData, err := parseJSON(msg.Data)
			if err != nil {
				fmt.Println("Ошибка при парсинге JSON:", err)
				return
			}
			fmt.Println("JSON спашрен")

			err = SaveToPostgreSQL(orderData)
			if err != nil {
				fmt.Println("Ошибка при сохранении в PostgreSQL:", err)
				return
			}
			fmt.Println("Данные из темы ", channelName, " сохранены в БД")
			// Запись в кэш
			MyCache.AddOrder(orderData)
			fmt.Println("Данные из темы ", channelName, " записаны в КЭШ")
		}, stan.DurableName(clientID))
	}

	return subscription, err
}

//func startHTTPServer() {
//	err := http.ListenAndServe(":3333", nil)
//	if err != nil {
//		fmt.Println("Ошибка при запуске HTTP сервера:", err)
//	}
//	fmt.Println("HTTP сервер запущен")
//}
//
//func handlePostRequest(w http.ResponseWriter, r *http.Request) {
//	// Чтение тела запроса
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		http.Error(w, "Ошибка при чтении тела запроса", http.StatusInternalServerError)
//		return
//	}

// Вывод содержимого тела запроса в консоль
//fmt.Println("Получен POST-запрос:")
//fmt.Println(string(body))

// Отправка ответа клиенту
//w.WriteHeader(http.StatusOK)
//w.Write([]byte("Запрос успешно обработан"))

//publishMessage(body, channelName, sc)
//}

// Функция для отправки сообщения в канал NATS Streaming
//func publishMessage(message []byte, channelName string, sc stan.Conn) {
//
//	err := sc.Publish(channelName, message)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Опубликовано сообщение на канале '%s': %s\n", channelName, message)
//}
