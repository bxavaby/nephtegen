package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// gen functions

func GenOneWord(styles []string, wordFormat string, minLen, maxLen int, prefixStyle, suffixStyle string) {

	randomLetter := fetchRandomLetter()
	randomLastLetter := fetchRandomLetter()

	prompt := fmt.Sprintf(
		"Generate a %s name with styles %s, length %d-%d characters, prefix style %s, suffix style %s. Start the name with the following letter(s) '%s'. Find the best way to fit in the following letter(s) as well: '%s'. Be unusual; take whatever concept you thought of first and recycle it multiple times, improve on it; deepening your sense of word-crafting.",
		wordFormat, formatStyles(styles), minLen, maxLen, prefixStyle, suffixStyle, randomLetter, randomLastLetter,
	)

	RunSpinner("Crafting Ren", 3*time.Second)

	result := callFabric(prompt)

	firstLine := strings.Split(result, "\n")[0]
	logCreation(firstLine)

	fmt.Println(PrimaryStyle.Render(fmt.Sprintf("\nEvoked Singular Ren: %s\n", result)))
}

func GenTwoWord(styles []string, wordFormat string, minLen, maxLen int, firstWordStyle, secondWordStyle string) {

	randomLetter := fetchRandomLetter()
	prompt := fmt.Sprintf(
		"Generate a %s name with styles %s, length %d-%d characters. The first word should have a %s style, and the second word should have a %s style. It is imperative that it starts with the following letter(s) '%s'.",
		wordFormat, formatStyles(styles), minLen, maxLen, firstWordStyle, secondWordStyle, randomLetter,
	)

	RunSpinner("Crafting Dual Ren", 3*time.Second)

	result := callFabricDual(prompt)

	firstLine := strings.Split(result, "\n")[0]
	logCreation(firstLine)

	fmt.Println(PrimaryStyle.Render(fmt.Sprintf("\nEvoked Dual Ren: %s\n", result)))
}

func GenThreeWord(styles []string, wordFormat string, minLen, maxLen int, firstStyle, secondStyle, thirdStyle string) {

	randomLetter := fetchRandomLetter()

	session := "tri"
	pattern := "triplewgen"

	prompt := fmt.Sprintf(
		"Craft a %s name with styles %s, length %d-%d characters. The first word should have a %s style, the second word a %s style, and the third word a %s style. Connect them with hyphens and make it unusual, add rhythm and a touch of genius. It is imperative that it starts with the following letter(s) '%s'.",
		wordFormat, formatStyles(styles), minLen, maxLen, firstStyle, secondStyle, thirdStyle, randomLetter,
	)

	RunSpinner("Crafting Tiple Ren", 5*time.Second)

	result := callFabricX(session, pattern, prompt)

	firstLine := strings.Split(result, "\n")[0]
	logCreation(firstLine)

	fmt.Println(PrimaryStyle.Render(fmt.Sprintf("\nEvoked Three-Word Ren: %s\n", result)))
}

func GenSymbolic(styles []string, wordFormat string, minLen, maxLen int, theme string) {

	randomLetter := fetchRandomLetter()

	session := "symb"
	pattern := "symbolicgen"

	prompt := fmt.Sprintf(
		"Create a %s with styles %s, length %d-%d characters, and theme: %s. It must start with '%s'.",
		wordFormat, formatStyles(styles), minLen, maxLen, theme, randomLetter,
	)

	RunSpinner("Crafting Symbolic Name", 5*time.Second)

	result := callFabricX(session, pattern, prompt)

	firstLine := strings.Split(result, "\n")[0]
	logCreation(firstLine)

	fmt.Println(PrimaryStyle.Render(fmt.Sprintf("\nEvoked Symbolic Name: %s\n", result)))
}

func GenHaiku(styles []string, wordFormat string, theme string) {

	randomLetter := fetchRandomLetter()

	session := "haiku"
	pattern := "haikugen"

	prompt := fmt.Sprintf(
		"Produce a %s with styles %s, 5-7-5 format, and theme: %s. It must start with '%s'.",
		wordFormat, formatStyles(styles), theme, randomLetter,
	)

	RunSpinner("Crafting Haiku", 8*time.Second)

	result := callFabricX(session, pattern, prompt)

	logHaiku(result)

	fmt.Println(PrimaryStyle.Render(fmt.Sprintf("\nEvoked Haiku:\n\n%s\n", result)))
}

func GenMotto(styles []string, wordFormat string, theme string) {

	randomLetter := fetchRandomLetter()

	session := "motto"
	pattern := "mottogen"

	prompt := fmt.Sprintf(
		"Craft a %s with styles %s,  and theme: %s. It must start with '%s'.",
		wordFormat, formatStyles(styles), theme, randomLetter,
	)

	RunSpinner("Crafting Motto", 8*time.Second)

	result := callFabricX(session, pattern, prompt)

	firstLine := strings.Split(result, "\n")[0]
	logCreation(firstLine)

	fmt.Println(PrimaryStyle.Render(fmt.Sprintf("\nEvoked Motto:\n%s\n", result)))
}

func GenPortmanteau(styles []string, wordFormat string, minLen, maxLen int, theme string) {

	randomLetter := fetchRandomLetter()

	session := "port"
	pattern := "ptmtugen"

	prompt := fmt.Sprintf(
		"Forge a %s name with styles %s, length %d-%d characters, and theme: %s. It must start with '%s'.",
		wordFormat, formatStyles(styles), minLen, maxLen, theme, randomLetter,
	)

	RunSpinner("Crafting Portmanteau", 5*time.Second)

	result := callFabricX(session, pattern, prompt)

	firstLine := strings.Split(result, "\n")[0]
	logCreation(firstLine)

	fmt.Println(PrimaryStyle.Render(fmt.Sprintf("\nEvoked Portmanteau: %s\n", result)))
}

func GenAbstractPhrase(styles []string, wordFormat string, minLen, maxLen int, theme string) {

	randomLetter := fetchRandomLetter()

	session := "asphrase"
	pattern := "abstkphra"

	prompt := fmt.Sprintf(
		"Generate a %s with styles %s, length %d-%d characters, and theme: %s. It must start with '%s'.",
		wordFormat, formatStyles(styles), minLen, maxLen, theme, randomLetter,
	)

	RunSpinner("Crafting Abstract Phrase", 8*time.Second)

	result := callFabricX(session, pattern, prompt)

	firstLine := strings.Split(result, "\n")[0]
	logCreation(firstLine)

	fmt.Println(PrimaryStyle.Render(fmt.Sprintf("\nEvoked Abstract Phrase: %s\n", result)))
}

func GenAlliterative(styles []string, wordFormat string, minLen, maxLen int, theme string) {

	randomLetter := fetchRandomLetter()

	session := "allit"
	pattern := "allitegen"

	prompt := fmt.Sprintf(
		"Generate a %s name with styles %s, length %d-%d characters, and theme: %s. Ensure all words begin with the same letter; innovation is key. It must start with '%s'.",
		wordFormat, formatStyles(styles), minLen, maxLen, theme, randomLetter,
	)

	RunSpinner("Crafting Alliterative Name", 8*time.Second)

	result := callFabricX(session, pattern, prompt)

	firstLine := strings.Split(result, "\n")[0]
	logCreation(firstLine)

	fmt.Println(PrimaryStyle.Render(fmt.Sprintf("\nEvoked Alliterative Name: %s\n", result)))
}

func GenAcronym(styles []string, wordFormat string, minLen, maxLen int, theme string) {

	randomLetter := fetchRandomLetter()

	session := "acronym"
	pattern := "acrogen"

	prompt := fmt.Sprintf(
		"Generate a %s with styles %s, length %d-%d characters, and theme: %s. Ensure it is concise and each letter represents a key word. Imagination is key. It must start with '%s'.",
		wordFormat, formatStyles(styles), minLen, maxLen, theme, randomLetter,
	)

	RunSpinner("Crafting Acronym", 5*time.Second)

	result := callFabricX(session, pattern, prompt)

	firstLine := strings.Split(result, "\n")[0]
	logCreation(firstLine)

	fmt.Println(PrimaryStyle.Render(fmt.Sprintf("\nEvoked Acronym: %s\n", result)))
}

// callFabric functions

func callFabric(prompt string) string {

	cmd := exec.Command("zsh", "-c", fmt.Sprintf(`fabric -p "wordgen" "%s"`, prompt))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error calling Fabric: %s\nOutput: %s\n", err, string(output))
		return "Error: Fabric command failed"
	}

	return strings.TrimSpace(string(output))
}

func callFabricDual(prompt string) string {
	session := "dual"

	cmd := exec.Command("zsh", "-c", fmt.Sprintf(`fabric --session "%s" -p "wordgen" "%s"`, session, prompt))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error calling Fabric: %s\nOutput: %s\n", err, string(output))
		return "Error: Fabric command failed"
	}

	return strings.TrimSpace(string(output))
}

func callFabricX(session string, pattern string, prompt string) string {

	cmd := exec.Command("zsh", "-c", fmt.Sprintf(`fabric --session "%s" -p "%s" "%s"`, session, pattern, prompt))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error calling Fabric: %s\nOutput: %s\n", err, string(output))
		return "Error: Fabric command failed"
	}

	return strings.TrimSpace(string(output))
}
