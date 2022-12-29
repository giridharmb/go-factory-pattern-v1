# go-factory-pattern

Steps

Generate API Key From : https://openweathermap.org/api

Clone The Repo

```bash
git clone https://github.com/giridharmb/go-factory-pattern-v1 go-factory-pattern
```

Building

```bash
cd go-factory-pattern
go build -o gfc
```

Running

```bash
./gfc -apikey <REDACTED> -city london
```

Output

```bash
2022/12/29 12:43:43 @ provider : path : /weather?q=london&appid=<REDACTED>&units=metric
2022/12/29 12:43:43 @ provider : completeURL : https://api.openweathermap.org/data/2.5/weather?q=london&appid=<REDACTED>&units=metric

{
    "Temp": 6.29,
    "Pressure": 1005,
    "MinTemp": 4.87,
    "MaxTemp": 7.36
}
```
