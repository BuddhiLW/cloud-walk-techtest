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
