package main


//imports
import (
	"math/rand"
	"time"
	"regexp"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func ElizaResponse(inputStr string) string{

	input :=inputStr
	
//back ticks instead of quotes to make sure it doesnt leave the characters first
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		//return the string below
		return "Why dont you tell me more about your father?"
	}
	
	if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, input); matched {
		//return the string below
		return "Why dont you tell me more about your mother?"
	}
	
	
	if matched, _ := regexp.MatchString(`(?i).*\bschool\b.*`, input); matched {
		//return the string below
		return "Blah! I hate school! Talk about anything else, please"
	}

//capture I am
	re := regexp.MustCompile("[iI] am ([^.!?]*)[.!?]?")
	
	if re.MatchString(input){
		return re.ReplaceAllString(input, "How do you know you are $1?") 

	}
	
// Capture He was 
	re = regexp.MustCompile("[hH]e was ([^.!?]*)[.!?]?")
	
	if re.MatchString(input){
		return re.ReplaceAllString(input, "Are you SURE he was $1?") 

	}
	
	// Capture She was
	// xir is temporary, for some reason i cant get it to capture she just for the time being 
	re = regexp.MustCompile("[xX]ir was ([^.!?]*)[.!?]?")
	
	if re.MatchString(input){
		return re.ReplaceAllString(input, "Are you SURE xir was $1?") 

	}

	

	answers := []string{
		"I'm not sure what you are trying to say. Could you explain it to me?",
		"How does that make you feel?",
		 "Why do you say that?",
		 "Not sure what you mean...but I can vibe to it!",
		 "Phuhuhuhu! You're very funny!",
		 "Thats a topic I dont even wanna touch!",
		 "hmmmmm...explain in detail?",
		 "Do you kiss your mother with that mouth?!",
		 "Gee thats something you dont want your father to hear!",
		 
		 
	}
	//returning a single string response
	return answers[rand.Intn(len(answers))]

}
//==================================================================================================================
func Reflect(input string) string {
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)
	
	// List the reflections.
	reflections := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `me`},
		{`my`, `your`},
		{`your`, `my`},
	}
	
	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}
	
	// Put the tokens back together.
	return strings.Join(tokens, ``)
}
// =========================================================================================================

func main(){
	rand.Seed(time.Now().UTC().UnixNano())//get a random number
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please message me")
	i := 1
	for i <= 5 {
		i++
	input, _ := reader.ReadString('\n')
	fmt.Scanf("%s", &input)
	fmt.Println(ElizaResponse(input))
	}
	

}
