# Cloud Walk Technical Test

## Introduction

The test consists in parsing a log file and outputting a json as a respose.

## Relevant data

Log to be parsed: https://gist.github.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8

## Goal (Parse)

Transform log into this structure, per match

``` json
"game_1": {
"total_kills": 45,
"players": ["Dono da bola", "Isgalamido", "Zeh"],
"kills": {
  "Dono da bola": 5,
  "Isgalamido": 18,
  "Zeh": 20
  }
}
```

*Print it on console*

## Goal (Script Report)

``` quote
Create a script that prints a report (grouped information) for each match and a player ranking.
```

``` sh

./main report 10
```

``` output
2024/03/13 14:19:02 JSON data for Match:
2024/03/13 14:19:02 Match chosen: 10
{"total_kills":60,"players":["Dono da Bola","Zeh","Chessus","Mal","Assasinu Credi","Isgalamido","Oootsimo"],"kills":{"Assasinu Credi":3,"Chessus":5,"Dono da Bola":3,"Isga
lamido":6,"Mal":1,"Oootsimo":-1,"Zeh":7}}
```



## Goal (Report by Death type)

``` 
Generate a report of deaths grouped by death cause for each match.
```

