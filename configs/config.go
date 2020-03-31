package config

import "fmt"

//rabbitUserName is the name of the user in Rabbit MQ
const rabbitUserName = "admin"
//rabbitPassword is the password of the user
const rabbitPassword = "admin"
//RabbitConnection is the connection string to the Rabbit MQ
var RabbitConnection = fmt.Sprintf("amqp://%s:%s@46.101.138.224:5672/", rabbitUserName, rabbitPassword)