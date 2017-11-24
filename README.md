# GoChat

For a class project I have been tasked with making a web app chat bot with go. For my reasearch into this project I looked up many different ways to actually code the chat bot itself as in my own opinion it is the most challenging part of the project. The best method I found was using regullar expression. As seen in many chat bots online, devs tend to give them a personality of a popular pop culture character, due to this I have decided to model my bots personality off "Monokuma" from the game series "Danganronpa". Having a personailty like this brings a certain uniquiness to the bot and will make its responses less boring. I have used regex to capture various responses and keywords enterted by the user, once the user enters one of these key words the bot will output a response that fits the word or sentace entered. If the user enters a string that is not captured by the bot, the bot will output a random response. Some of the random responses that the bot outputs I have called "Converstation starters", for example, if the user enters a string that is not captured by the bot, the bot will reply with a random string that says "Gee do you kiss your mother with that mouth!?"which in turn tempts the user to make his/her next string feture the word "mother" which is captured by the bot already. Here is a list of phrases that the bot will pick up upon once entered. 

Generak words: 
Hello 
Father 
Mother 
School 
Despair 
Junko 
Headmaster 
Sing 
dance
rawr
hope 

Captures: 
I am 
You are 
shut up 
are you

Random strings can be found inside the go code. 

How to run it: 
1.) copy the contents of the "go project folder" 
2.) open up cmd window. 
3.) cd to the go project directory 
4.) type "go build" while inside the folder to create an exe file which can be run. 
5.) alternativly, type "go run eliza.go" and navigate to "http://127.0.0.1:8080/eliza.html" while the file is running. 

How to use: 
1.) simply type various phrases into the text box and the monokuma robot will respond to you! 

