package main

import (
	"fmt"
	"log"
	"os"

	chatbot "github.com/green-api/whatsapp-chatbot-golang"
	"github.com/joho/godotenv"
)

func sendMessage(bot *chatbot.Bot, phoneNumber, message string) {
	response, err := bot.GreenAPI.Methods().Sending().SendMessage(map[string]interface{}{
		"chatId":  phoneNumber + "@c.us", // Formato correto para WhatsApp
		"message": message,
	})

	if err != nil {
		log.Printf("Erro ao enviar mensagem para %s: %v\n", phoneNumber, err)
	} else {
		fmt.Printf("Mensagem enviada para %s: %v\n", phoneNumber, response)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	apiToken := os.Getenv("APITOKEN")
	instanceID := os.Getenv("INSTANCE_ID")
	os.Getenv("")
	bot := chatbot.NewBot(instanceID, apiToken)

	numbers := []string{"5521973946969", "5521991202558"}
	message := `🌟 Transforme sua vida com a Intera! 🌟

Aqui, oferecemos um atendimento personalizado e integrado. Se você está pronto para dar o próximo passo rumo a uma vida extraordinária, estamos aqui para ajudar!

📞 Entre em contato conosco pelo WhatsApp e agende. Juntos, podemos transformar sua vida!


👉 2743-1176`

	for _, number := range numbers {
		sendMessage(bot, number, message)
	}
}
