package main
// Imports and packages and the likes 
import (
    "fmt"
    "html/template"
    "net/http"
    "math/rand"
    "regexp"

)
// function that returns the bots responses if the user input matches string, got these from one of the labs we done 
func ElizaResponse(input string) string {
    // If the bot picks up father, return his reply 
    if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
    
        return "daddy issues dont really intrest me that much"
    }

    if matched, _ := regexp.MatchString(`(?i).*\bhello\b.*`, input); matched {
    
        return "Howdy howdy hey! "
    }

    // Bot picks up mother 
    if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, input); matched {
		//return the string below
		return "Ohhhh you have a mom? Tell me more!"
	}
	
	// Bot picks up school 
	if matched, _ := regexp.MatchString(`(?i).*\bschool\b.*`, input); matched {
		//return the string below
		return "Blah! I hate school! Talk about anything else, please"
	}

    // Bot picks up despair 
    if matched, _ := regexp.MatchString(`(?i).*\bdespair\b.*`, input); matched {
		//return the string below
		return "Hmmmm...now your talking my language!"
	}
    // Bot picks up junko 
    if matched, _ := regexp.MatchString(`(?i).*\bjunko\b.*`, input); matched {
		//return the string below
		return "Eh eh eh I have no idea who you're talking about!"
	}
    // Bot picks up headmaster 
    if matched, _ := regexp.MatchString(`(?i).*\bheadmaster\b.*`, input); matched {
		//return the string below
		return "Im the only headmaster around here, buddy! I've owned 53 schools!"
	}
    // Bot picks up sing 
    if matched, _ := regexp.MatchString(`(?i).*\bsing\b.*`, input); matched {
		//return the string below
		return "Do I look like some kind of show bear to you? "
	}
    // Bot picks up dance 
    if matched, _ := regexp.MatchString(`(?i).*\bdance\b.*`, input); matched {
		//return the string below
		return "Sorry pal, but do you see anyway I can move my non existent body? I ain't no .gif!"
	}

    if matched, _ := regexp.MatchString(`(?i).*\rawr\b.*`, input); matched {
		//return the string below
		return "GRRRRRRR RAAAAAAAWWWWRRRRRR!"
	}
    // Bot picks up hope 
    if matched, _ := regexp.MatchString(`(?i).*\bhope\b.*`, input); matched {
		//return the string below
		return "My code tells me I cant make threats, but dont talk about that...please"
	}



    
    // captures can be used to take the user input after the captured sentance and then return it 
    // Capture I am 
    // square boxes are used for case sensitivity in reg exp 
    re := regexp.MustCompile(`(?i)[iI] am ([^.?!]*)[.?!]?`)
    
    if matched := re.MatchString(input); matched {
    
        return re.ReplaceAllString(input, "How do you know you are $1?")
    }    
    // Capture You are 
    re = regexp.MustCompile(`(?i)[yY]ou are ([^.?!]*)[.?!]?`)
    
    if matched := re.MatchString(input); matched {
    
        return re.ReplaceAllString(input, "Hmpfh! How dare you say I'm $1! I'm the best bear!")
    }    
    // Capture shut up
    re = regexp.MustCompile(`(?i)[sS]hut up ([^.?!]*)[.?!]?`)
    
    if matched := re.MatchString(input); matched {
    
        return re.ReplaceAllString(input, "phuhuh... why dont YOU shut up $1?")
    }   
    // Capture 
    re = regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)
    
    if matched := re.MatchString(input); matched {
    
        return re.ReplaceAllString(input, "How do you know you are $1?")
    }     
    // Capture Are you 
    re = regexp.MustCompile(`(?i)[aA]re you ([^.?!]*)[.?!]?`)
    
    if matched := re.MatchString(input); matched {
    
        return re.ReplaceAllString(input, "I dont know tough guy, am I $1?")
    }  
    // random strings to be used if the bot cant find anything that matches the regex
    answers := []string{
    
        "I'm not sure what you are trying to say. Could you explain it to me?",
		"How does that make you feel?",
		 "Why do you say that?",
		 "Not sure what you mean...but I can vibe to it!",
		 "Phuhuhuhu! You're very funny!",
		 "Thats a topic I dont even wanna touch!",
		 "hmmmmm...explain in detail?",
		 "Do you kiss your mother with that mouth?!",
		 "Sorry pal, but I have no idea what you said, heres a random response!",
         "Intresting! But you should check out my muscles!",
         "Yawwwnnnnnn, You're so dull!",
         "Phuhuhu! You're my kinda guy!",
         "Spending your day on a danganronpa chatbot? God you must be boooriinggg",
         "Fourth wall breaks are cheap, but this is the final random response!",

    }
    
    return answers[rand.Intn(len(answers))]
    
}
// main function 
func main() {
   // templates
    tmpl1 := template.Must(template.ParseFiles("eliza.html"))
    tmpl2 := template.Must(template.ParseFiles("echomessage.html"))

   
    // handles function with http 
    http.HandleFunc("/eliza.html", func(w http.ResponseWriter, r *http.Request) {
            
            tmpl1.Execute(w, nil)
            
    })
    // used to reply the bots response to page 
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
	
} // End
