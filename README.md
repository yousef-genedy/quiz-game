# Gophercises Quiz Game

A command-line quiz game written in Go. Questions and answers are loaded from a CSV file, and the user is scored on how
many they get right — with an optional timed mode.

This is Exercise #1 from [gophercises.com](https://gophercises.com/).

## Features

- Loads quiz questions from a CSV file (defaults to `problems.csv`)
- Customizable CSV file path via flag
- Tracks and reports correct/incorrect answers
- Timed mode: quiz stops as soon as the time limit is reached, even mid-question
- Customizable time limit via flag
- Optional shuffling of question order

## CSV Format

Each row is a question and its answer, separated by a comma:

```
5+5,10
7+3,10
1+1,2
```

## Usage

```bash
go run .
```

With flags:

```bash
go run . -csv=problems.csv -limit=30
```

| Flag     | Description                         | Default        |
|----------|-------------------------------------|----------------|
| `-csv`   | Path to the quiz CSV file           | `problems.csv` |
| `-limit` | Time limit for the quiz, in seconds | `30`           |

At the end of the quiz, your score is printed as:

```
You scored X out of Y.
```

## About

Part of my ongoing practice with [gophercises.com](https://gophercises.com/) exercises while learning Go.
