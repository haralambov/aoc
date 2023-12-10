# [🎄🎁 Advent of Code 🎁🎄](https://adventofcode.com/)

My solutions to the Advent of Code puzzles, written in Go

<sub><em>Disclaimer: I am not a professional Go developer.  
The goal of solving AoC problems is to gain knowledge in that language.</em></sub>

## Progress

| Day\Year | 2023 | 2022 | 2021 | 2020 | 2019 | 2018 | 2017 | 2016 | 2015 |
| :------: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
|     1    |  ⭐  |      |      |      |      |      |      |      |  ⭐  |
|     2    |  ⭐  |      |      |      |      |      |      |      |  ⭐  |
|     3    |  ⭐  |      |      |      |      |      |      |      |      |
|     4    |  ⭐  |      |      |      |      |      |      |      |      |
|     5    |      |      |      |      |      |      |      |      |      |
|     6    |  ⭐  |      |      |      |      |      |      |      |      |
|     7    |  ⭐  |      |      |      |      |      |      |      |      |
|     8    |  ⭐  |      |      |      |      |      |      |      |      |
|     9    |  ⭐  |      |      |      |      |      |      |      |      |
|    10    |      |      |      |      |      |      |      |      |      |
|    11    |      |      |      |      |      |      |      |      |      |
|    12    |      |      |      |      |      |      |      |      |      |
|    13    |      |      |      |      |      |      |      |      |      |
|    14    |      |      |      |      |      |      |      |      |      |
|    15    |      |      |      |      |      |      |      |      |      |
|    16    |      |      |      |      |      |      |      |      |      |
|    17    |      |      |      |      |      |      |      |      |      |
|    18    |      |      |      |      |      |      |      |      |      |
|    19    |      |      |      |      |      |      |      |      |      |
|    20    |      |      |      |      |      |      |      |      |      |
|    21    |      |      |      |      |      |      |      |      |      |
|    22    |      |      |      |      |      |      |      |      |      |
|    23    |      |      |      |      |      |      |      |      |      |
|    24    |      |      |      |      |      |      |      |      |      |
|    25    |      |      |      |      |      |      |      |      |      |

<strong>Legend:</strong>  
🚧 First part  
⭐ Both parts

## Building and running

```
go build
./aoc --help
```

```
Usage of ./aoc:
  -day int
        Selected day. Must be between 1 and 25 (default 1)
  -part int
        Selected part. Must be either 1 or 2 (default 1)
  -year int
        Selected year. Must be between 2015 and 2023 (default 2023)
```

### Examples

#### Running for a specific year, day and part with my input

```
# Command
./aoc --year 2023 --day 1 --part 1
```

```
# Ouotput
53334
```

#### Running for a specific year, day and part with external input

##### With input redirection

```
# Sample command
./aoc --year 2023 --day 1 --part 1 < /path/to/input_file

# e.g. external file containing the input for that day
./aoc --year 2023 --day 1 --part 1 < /home/zlatomir/Documents/external_input

# Output
53334
```

##### With a pipe

```
# Sample command
cat /path/to/input_file | ./aoc --year 2023 --day 1 --part 1

# e.g. external file containing the input for that day
cat /home/zlatomir/Documents/external_input | ./aoc --year 2023 --day 1 --part 1

# Output
53334
```

