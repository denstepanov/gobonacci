# gobonacci

Конкурентный расчёт числа Фибоначчи.

## Стек

1. GO 1.19.3
2. chi
3. twirp
4. cleanenv

## Описание работы

Шаги, выполняемые приложением:

1. Приём числового параметра на REST эндпоинт первого сервиса.
2. Генерация горутин на основе переданного числового параметра. В каждой горутине происходит начальный расчёт числа Фибоначчи и передаёт результат по gRPC во второй сервис.
3. Второй сервис производит расчёт и отправляет данные в ответ на запрос по gRPC.
4. В хэндлере первого сервиса происходит бесконечная отправка новых данных

## Команды

grpc генерация (выполнять из корня):

```shell
protoc --twirp_out=. --go_out=. shared/proto/count.proto
```
