package basic

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func regexp_compile_demo() {

	words := [...]string{"Seven", "even", "Maven", "Amen", "eleven"}

	re, err := regexp.Compile(".even")
	// re := regexp.MustCompile(".even")

	if err != nil {
		log.Fatal(err)
	}

	for _, word := range words {
		found := re.MatchString(word)

		if found {
			fmt.Printf("%s matches\n", word)
		} else {
			fmt.Printf("%s does not match\n", word)
		}
	}
}

func regexp_findallstring_demo() {

	var content = `Foxes are omnivorous mammals belonging to several genera 
of the family Canidae. Foxes have a flattened skull, upright triangular ears, 
a pointed, slightly upturned snout, and a long bushy tail. Foxes live on every 
continent except Antarctica. By far the most common and widespread species of 
fox is the red fox.`

	re := regexp.MustCompile("(?i)fox(es)?")
	found := re.FindAllString(content, -1)

	fmt.Printf("%q\n", found)

	if found == nil {
		fmt.Printf("no match found\n")
		os.Exit(1)
	}
	fmt.Println()

	for _, word := range found {
		fmt.Printf("%s\n", word)
	}
	fmt.Println()

	idx := re.FindAllStringIndex(content, -1)
	for _, j := range idx {
		match := content[j[0]:j[1]]
		fmt.Printf("%s at %d:%d\n", match, j[0], j[1])
	}

}

func regexp_findstring_demo() {

	str := "Today is Tuesday!"
	regexp, _ := regexp.Compile("^T([a-z]+)y")

	fmt.Println(regexp.FindString(str))
}

func regexp_find_submatch_demo() {

	websites := [...]string{"webcode.me", "zetcode.com", "freebsd.org", "netbsd.org"}

	re := regexp.MustCompile("(\\w+)\\.(\\w+)")

	for _, website := range websites {

		parts := re.FindStringSubmatch(website)

		for i, _ := range parts {
			fmt.Println(parts[i])
		}
		fmt.Println("---------------------")
	}
}

func regexp_match_string_demo() {

	words := [...]string{"Seven", "even", "Maven", "Amen", "eleven"}

	for _, word := range words {

		found, err := regexp.MatchString(".even", word)

		if err != nil {
			log.Fatal(err)
		}

		if found {

			fmt.Printf("%s matches\n", word)
		} else {

			fmt.Printf("%s does not match\n", word)
		}
	}
}
