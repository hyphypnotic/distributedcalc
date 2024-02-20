не забудьте переименовать папку на distributedcalc

установите все зависимости:

0. go get github.com/gorilla/sessions


запустите код:


0. run in terminal redis-server --daemonize yes 
2. run in terminal "go mod download"
3. run in terminal "go build backend/cmd/agent/main.go"
4. run in terminal "go build backend/cmd/orchestrator/main.go"
4.что бы протестировать нужно зайти на http://localhost:8080/

Я не сделал список операций.

работает это так Оркестратору приходит арифметическое выражение он посылает его агенту через Redis и получает ответ
