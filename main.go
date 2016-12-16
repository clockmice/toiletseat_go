package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)

	fmt.Println(policy1(s))
	fmt.Println(policy2(s))
	fmt.Println(policy3(s))
}

func policy1(s string) int{
	count := 0

	if s[0] != s[1] {
		count++
	}
	if s[1] == 'D' {
		count++
	}
	if len(s) > 2 {
		for i := 2; i < len(s); i++ {
			if s[i] == 'D' {
				count = count + 2;
			}
		}
	}

	return count
}

func policy2(s string) int{
	count := 0

	if s[0] != s[1] {
		count++;
	}

	if s[1] == 'U' {
		count++;
	}

	if len(s) > 2 {
		for i := 2; i < len(s); i++ {
			if s[i] == 'U' {
				count = count + 2;
			}
		}
	}

	return count
}

func policy3(s string) int{
	count := 0;

	for i := 0; i < len(s) - 1; i++ {
		if s[i] != s[i + 1] {
			count++;
		}
	}
	return count
}

