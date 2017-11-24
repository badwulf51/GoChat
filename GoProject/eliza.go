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


// main function 
func main() {
   
    tmpl1 := template.Must(template.ParseFiles("eliza.html"))
    tmpl2 := template.Must(template.ParseFiles("echomessage.html"))

   
    
    http.HandleFunc("/eliza.html", func(w http.ResponseWriter, r *http.Request) {
            
            tmpl1.Execute(w, nil)
            
    })
    // used to reply the bots response to paghe 
    http.HandleFunc("/echomessage.html", func(w http.ResponseWriter, r *http.Request) {
            
        r.ParseForm()
        usermessage := r.Form.Get("message")
        fmt.Println(usermessage)
        
        response := ElizaResponse(usermessage)
        fmt.Println(response)
        
        tmpl2.Execute(w, struct {Message string}{response})
    })
    // web server (127.0.0.1:8080/eliza.html)
    http.ListenAndServe(":8080", nil)
}
