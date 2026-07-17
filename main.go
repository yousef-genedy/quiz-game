package main

func main() {
	file, limit, shuffle := ReadFlags()
	problems := ReadProblems(*file)
	if *shuffle {
		problems = Shuffle(problems)
	}
	ok := Interrupt()
	Timer(*limit, problems, ok)
}
