package main

import (
	"flag"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"playground/client/kinohod"
	"sort"
	"strconv"
	"strings"
)

func main() {
	telegramKey := flag.String("t_key", "", "Telegram API key")
	flag.Parse()

	bot, err := tgbotapi.NewBotAPI(*telegramKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60

	updates, err := bot.GetUpdatesChan(update)

	for update := range updates {
		fmt.Print(update)

		if update.CallbackQuery == nil {
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}

			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать")

				var keyboard = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Что сейчас в кино?"),
					),
				)

				msg.ReplyMarkup = keyboard

				_, err := bot.Send(msg)
				if err != nil {
					log.Fatal(err)
				}

				break
			case "Что сейчас в кино?":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Сеансы:")

				movies, err := kinohod.GetMovies()
				if err != nil {
					log.Fatal(err)
				}

				var InlineKeyboard = tgbotapi.NewInlineKeyboardMarkup()

				sort.Slice(movies, func(i, j int) bool {
					return movies[i].Attributes.ImdbRating > movies[j].Attributes.ImdbRating
				})

				for i := 0; i < 6; i++ {
					movie := movies[i]

					if movie.Attributes.ImdbRating != "" {
						InlineKeyboard.InlineKeyboard = append(InlineKeyboard.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonData(movie.Attributes.Title+" "+movie.Attributes.ImdbRating, "movie_"+movie.Id),
						))
					}
				}

				msg.ReplyMarkup = InlineKeyboard

				_, err = bot.Send(msg)
				if err != nil {
					log.Fatal(err)
				}

				break
			default:
				log.Println(update.Message.Text)
				break
			}
		} else {
			data := update.CallbackQuery.Data
			switch strings.Split(data, "_")[0] {
			case "1":
				// TODO
				break
			case "movie":
				id, err := strconv.Atoi(strings.Split(data, "_")[1])
				if err != nil {
					log.Fatalln(err)
				}

				movie, err := kinohod.GetMovie(id)
				if err != nil {
					log.Fatalln(err)
				}

				title := movie.Attributes.Title
				year := movie.Attributes.ProductionYear
				genres := movie.Attributes.Genres[0].Name
				if len(movie.Attributes.Genres) > 1 {
					for i := 1; i < len(movie.Attributes.Genres); i++ {
						genres += ", " + movie.Attributes.Genres[i].Name
					}
				}
				annotation := movie.Attributes.AnnotationFull

				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, ""+
					"Вы выбрали <b>"+title+" "+"("+year+")</b> - "+genres+".\n"+
					annotation)
				msg.ParseMode = "HTML"

				_, err = bot.Send(msg)
				if err != nil {
					log.Fatalln(err)
				}
			default:
				break

				//time, path := service.GetPath(47, 152)
				//
				//fmt.Println("Time", time, "minutes")
				//fmt.Println("Path", path)
			}

			_, err := bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
			if err != nil {
				log.Fatalln(err)
			}
		}

		// 1. Список фильмов (ТОП по рейтингу)
		// 2. Местоположение? Не решил, на каком этапе спрашивать
		// 3. Список ближайших сеансов с учетом дороги на метро
		// 		1. Список сеансов.
		// 		2. Координаты кинотеатров и расстояние до метро.
		// 		3. Расстояние между польователем и ближайшей станцией.
		//		4. Время на метро.
		//			(https://www.moscowmap.ru/_ajax/metro-p2p-m.php?st1=233&st2=176, придется парсить страницу для id станций)
		//			(хуяндекс карты)
		//		5. Упорядочить по убыванию суммарного времени пути.
		//
		// Далее:
		// 		Показывать загруженность залов
		// 		Бронировать сеанс
	}
}
