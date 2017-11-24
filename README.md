# GoChat

N.B. VIEW IN RAW PASTE DATA FORM PLEASE.

For a class project I have been tasked with making a web app chat bot with go. For my research into this project I looked up many different ways to actually code the chat bot itself as in my own opinion it is the most challenging part of the project. The best method I found was using regular expression. As seen in many chat bots online, devs tend to give them a personality of a popular pop culture character, due to this I have decided to model my bots personality off "Monokuma" from the game series "Danganronpa". Having a personality like this brings a certain uniqueness to the bot and will make its responses less boring. I have used regex to capture various responses and keywords entered by the user, once the user enters one of these key words the bot will output a response that fits the word or sentence entered. If the user enters a string that is not captured by the bot, the bot will output a random response. Some of the random responses that the bot outputs I have called "Conversation starters", for example, if the user enters a string that is not captured by the bot, the bot will reply with a random string that says "Gee do you kiss your mother with that mouth!?"Which in turn tempts the user to make his/her next string feature the word "mother" which is captured by the bot already. Here is a list of phrases that the bot will pick up upon once entered. 

General words: 
Hello 
Father 
Mother 
School 
Despair 
Junko 
Headmaster 
Sing 
Dance
Rawr
Hope 

Captures: 
I am 
You are 
Shut up 
Are you

Random strings can be found inside the go code. 

How to run it: 
1.) Copy the contents of the "go project folder" 
2.) Open up cmd window. 
3.) Cd to the go project directory 
4.) type "go build" while inside the folder to create an exe file which can be run. 
5.) Alternatively, type "go run eliza.go" and navigate to "http://127.0.0.1:8080/eliza.html" while the file is running. 

How to use: 
1.) Simply type various phrases into the text box and the Monokuma robot will respond to you!


