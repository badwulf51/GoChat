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

	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
    
        return "daddy issues dont really intrest me that much"
    }


    if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, input); matched {
		//return the string below
		return "Ohhhh you have a mom? Tell me more!"
	}
	
	
	if matched, _ := regexp.MatchString(`(?i).*\bschool\b.*`, input); matched {
		//return the string below
		return "Blah! I hate school! Talk about anything else, please"
	}


    if matched, _ := regexp.MatchString(`(?i).*\bdespair\b.*`, input); matched {
		//return the string below
		return "Hmmmm...now your talking my language!"
	}

    if matched, _ := regexp.MatchString(`(?i).*\bjunko\b.*`, input); matched {
		//return the string below
		return "Eh eh eh I have no idea who you're talking about!"
	}

    // Capture I am 
    
    re := regexp.MustCompile(`(?i)[iI] am ([^.?!]*)[.?!]?`)
    
    if matched := re.MatchString(input); matched {
    
        return re.ReplaceAllString(input, "How do you know you are $1?")
    }    

    re = regexp.MustCompile(`(?i)[yY]ou are ([^.?!]*)[.?!]?`)
    
    if matched := re.MatchString(input); matched {
    
        return re.ReplaceAllString(input, "Hmpfh! How dare you say I'm $1! I'm the best bear!")
    }    
    
    re = regexp.MustCompile(`(?i)[sS]hut up ([^.?!]*)[.?!]?`)
    
    if matched := re.MatchString(input); matched {
    
        return re.ReplaceAllString(input, "phuhuh... why dont YOU shut up $1?")
    }   

    re = regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)
    
    if matched := re.MatchString(input); matched {
    
        return re.ReplaceAllString(input, "How do you know you are $1?")
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
