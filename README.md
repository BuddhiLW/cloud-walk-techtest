# Cloud Walk Technical Test

## How to run this project

### Requirements 
- go

### How to produce the binaries (Install)

``` sh
git clone https://github.com/BuddhiLW/cloud-walk-techtest.git
bash ./build.sh
```

Now you should have the new binary file in the project root called `qrep` (qrep stands for *quake report*).

### How to use the binaries

Make sure the permissions are given to read, write and execute,

``` sh
chmod +rwx ./qrep
```

Then, use `qrep` as follows,

#### To get an explanation of the command (help like a `man`ual page)
``` sh
./qrep data help
```

``` output
NAME
       qrep - qrep is a quake-report binary. It's goal is to output JSON data for the chosen match, arg should be a <number: integer>

ALIASES
       qrep (data|json)

SYNOPSIS
       qrep qrep <number>

COMMANDS
       help             - display help similar to man page format
       r|rank
       stats|statistics
```

#### To get all games data (json format)

Call `qrep data` without arguments: 
``` sh
./qrep data 
```

The expected output looks like this:

 ```  output
All games (N):  22
---------------------------------

{
    "game_1": {
        "total_kills": 0,
        "players": [
            "Isgalamido"
        ],
        "kills": {
            "Isgalamido": 0
        }
    },
    "game_10": {
        "total_kills": 60,
        "players": [
            "Mal",
            "Assasinu Credi",
            "Isgalamido",
            "Oootsimo",
            "Dono da Bola",
            "Zeh",
            "Chessus"
        ],
        "kills": {
(...)
```

Notes:
- The output is not ordered like "game_1", "game_2", etc., but contains all games.
- The message "All games (N):  22\n ------------" is only in `log` (debug channel);
- Therefore, the output can be "isolated" to only the json data by pipping the result to a file. E.g., `./qrep data > tmp.txt`

#### To get only the <nth> game data (json format)

Call `qrep data n`, with `n` as the argument: 
``` sh
./qrep data <n>
```

The expected output looks like this:

``` sh
./qrep data 10
```

``` output
2024/03/13 19:35:04 JSON data for Match:
2024/03/13 19:35:04 Match chosen: 10
{
    "total_kills": 60,
    "players": [
        "Chessus",
        "Mal",
        "Assasinu Credi",
        "Isgalamido",
        "Oootsimo",
        "Dono da Bola",
        "Zeh"
    ],
    "kills": {
        "Assasinu Credi": 3,
        "Chessus": 5,
        "Dono da Bola": 3,
        "Isgalamido": 6,
        "Mal": 1,
        "Oootsimo": -1,
        "Zeh": 7
    }
}
```


---------------
# Introduction to the **test** specification

The test consists in parsing a log file and outputting a json as a respose.

## Relevant data

Log to be parsed: https://gist.github.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8

## Goal (Parse)

Transform log into this structure, per game-match

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

## Goal (Script Report and Ranking)

``` quote
Create a script that prints a report (grouped information) for each match and a player ranking.
```

### Report about the match (Implementation)
``` sh
./main report 10
```

``` output
2024/03/13 21:05:44 JSON data for Match:
2024/03/13 21:05:44 Match chosen: 10
{
    "total_kills": 60,
    "players": [
        "Mal",
        "Assasinu Credi",
        "Isgalamido",
        "Oootsimo",
        "Dono da Bola",
        "Zeh",
        "Chessus"
    ],
    "kills": {
        "Assasinu Credi": 3,
        "Chessus": 5,
        "Dono da Bola": 3,
        "Isgalamido": 6,
        "Mal": 1,
        "Oootsimo": -1,
        "Zeh": 7
    }
}
```

### Player ranking for the match (Implementation)

The compose-command `rank` defaults to a pretty print
``` sh
./main report rank 10
```

``` output
2024/03/13 21:05:16 Rank of match number << 10 >>
2024/03/13 21:05:16 --------- (by kills) -----------
2024/03/13 21:05:16 In text:

Name: Zeh,            Kills: 7,  Position: 1
Name: Isgalamido,     Kills: 6,  Position: 2
Name: Chessus,        Kills: 5,  Position: 3
Name: Assasinu Credi, Kills: 3,  Position: 4
Name: Dono da Bola,   Kills: 3,  Position: 5
Name: Mal,            Kills: 1,  Position: 6
Name: Oootsimo,       Kills: -1, Position: 7
```

But, you can also specify a json output format, like so:

``` sh
./main report rank json 10
```

``` output
2024/03/13 21:07:48 In json:

{
    "Assasinu Credi": {
        "name": "Assasinu Credi",
        "kills": 1,
        "position": 3
    },
    "Dono da Bola": {
        "name": "Dono da Bola",
        "kills": 0,
        "position": 4
    },
    "Isgalamido": {
        "name": "Isgalamido",
        "kills": 2,
        "position": 1
    },
    "Zeh": {
        "name": "Zeh",
        "kills": 1,
        "position": 2
    }
}
```



## Goal (Report by Death type)

``` 
Generate a report of deaths grouped by death cause for each match.
```


