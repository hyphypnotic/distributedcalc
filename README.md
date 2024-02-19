0. run in terminal redis-server --daemonize yes
1. run in terminal "go mod download"
2. run in terminal "go build backend/cmd/agent/main.go"
3. run in terminal "go build backend/cmd/orchestrator/main.go"
Я не сделал список операций.

работает это так Оркестратору приходит арифметическое выражение он посылает его агенту через Redis и получает ответ
