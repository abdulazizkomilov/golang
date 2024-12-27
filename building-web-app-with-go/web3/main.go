// import (
// 	"bufio"
// 	"crypto/rand"
// 	"crypto/sha1"
// 	"database/sql"
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/gorilla/mux"
// 	"github.com/gorilla/sessions"
// 	"github.com/streadway/amqp"
// 	"html/template"
// 	"io"
// 	"log"
// 	"net/http"
// 	"regexp"
// 	"text/template"
// 	"time"
// )

// var WelcomeTitle = "You've successfully registered!"
// var WelcomeEmail = "Welcome to our CMS, {{Email}}! We're glad you could join us."

// const (
// 	DBHost  = "127.0.0.1"
// 	DBPort  = ":3306"
// 	DBUser  = "root"
// 	DBPass  = ""
// 	DBDbase = "cms"
// 	PORT    = ":8080"
// 	MQHost  = "127.0.0.1"
// 	MQPort  = ":5672"
// )

// func MQConnect() (*amqp.Connection, *amqp.Channel, error) {
// 	url := "amqp://" + MQHost + MQPort
// 	conn, err := amqp.Dial(url)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	channel, err := conn.Channel()
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	if _, err := channel.QueueDeclare("", false, true, false, false,
// 		nil); err != nil {
// 		return nil, nil, err
// 	}
// 	return conn, channel, nil
// }

// type RegistrationData struct {
// 	Email string `json:"email"`
// 	Message string `json:"message"`
// }

// res, err := database.Exec("INSERT INTO users SET user_name=?, user_guid=?, user_email=?, user_password=?", name, guid, email, password)
// if err != nil {
// 	fmt.Fprintln(w, err.Error)
// } else {
// 	Email := RegistrationData{Email: email, Message: ""}
// 	message, err := template.New("email").Parse(WelcomeEmail)
// 	var mbuf bytes.Buffer
// 	message.Execute(&mbuf, Email)
// 	MQPublish(json.Marshal(mbuf.String()))
// 	http.Redirect(w, r, "/page/"+pageGUID, 301)
// }

// func MQPublish(message []byte) {
// 	err = channel.Publish(
// 	"email", // exchange
// 	"", // routing key
// 	false, // mandatory
// 	false, // immediate
// 	amqp.Publishing{
// 	ContentType: "text/plain",
// 	Body: []byte(message),
// 	})
// }