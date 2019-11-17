package main

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	"playground/client/imdb"
	"playground/client/kinohod"
	"playground/client/seance"
	"strconv"
)

func main() {

	Test()

	bot, err := tgbotapi.NewBotAPI("")
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
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

				var InlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Фильм 1", "1"),
					),
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Фильм 2", "2"),
					),
				)

				msg.ReplyMarkup = InlineKeyboard

				_, err := bot.Send(msg)
				if err != nil {
					log.Fatal(err)
				}

				break
			}
		} else {
			switch update.CallbackQuery.Data {
			case "1":
				bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Вы выбрали фильм "+update.CallbackQuery.Data))
				break
			case "2":
				bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Вы же выбрали фильм "+update.CallbackQuery.Data))
				break
			}

			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
		}
	}

	//1. Запросить список фильмов.

	movies, err := kinohod.GetMovies()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(movies)

	movieId, err := strconv.Atoi(movies[0].Id)
	if err != nil {
		log.Fatalln(err)
	}

	imdbId, err := strconv.Atoi(movies[0].Attributes.ImdbId)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(movies[0])

	client2 := imdb.Client{}
	result2 := client2.Get(imdbId)
	fmt.Println(result2.Ratings[0].Value)

	client3 := seance.Client{}

	fmt.Println(client3.Get(1, movieId))

	//3. Для самого крутого показать кинотеатры с координатами.

}
