This project contain normal Golang Gin Framework with client that consuming Websocket is NodeJS
We have normal WebSocket and another WebSocket that consumed message from RabbitMQ
Guides:
go get -u github.com/gin-gonic/gin
go get github.com/gorilla/websocket
go get 

Example set up RabbitmQ
RabbitmQ URL: http://localhost:15672/#/

##to set a new admin user
rabbitmqctl add_user newadmin s0m3p4ssw0rd
rabbitmqctl set_user_tags newadmin administrator
rabbitmqctl set_permissions -p / newadmin ".*" ".*" ".*"
