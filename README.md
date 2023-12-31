Guides:
go get -u github.com/gin-gonic/gin
go get github.com/gorilla/websocket
go get 

Example set up RabbitmQ
RabbitmQ URL: http://localhost:15672/#/

rabbitmqctl add_user newadmin s0m3p4ssw0rd
rabbitmqctl set_user_tags newadmin administrator
rabbitmqctl set_permissions -p / newadmin ".*" ".*" ".*"
