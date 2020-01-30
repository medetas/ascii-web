package main

import(
	"os"
	"fmt"
	"log"
	"bufio"
	//"github.com/01-edu/z01"
	"strings"
	//student "../go-reloaded"
)

func Ascii(args []string) string{
	/*if len(args) == 1{
		fmt.Println()
		return 
	}	
	*/
	outputFile := "output.txt"
    
	filename := "standard.txt"
	var str string
	for _, flag := range args{        //checking flags for fonts and colors
		if strings.EqualFold(flag, "shadow"){
			filename = "shadow.txt"
		}else if strings.EqualFold(flag, "thinkertoy"){
			filename = "thinkertoy.txt"
		}else if strings.EqualFold(flag, "standard"){
			filename = "standard.txt"
		}else if len(flag) > 9 && strings.EqualFold(flag[:9], "--output="){
			outputFile = flag[9:]
		}else if len(flag) == 9 && strings.EqualFold(flag[:9], "--output="){
			continue
		}else if str == ""{
			str = flag
		}else{
			str = str +" " + flag		
		}
	}


	//Output.txt Needed for the next project
	output, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()
	outputString := ""

	file, err := os.Open(filename)     //banner file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	var art string
	data := make(map[rune]string)
	var currChar rune
	currChar = ' '

	for scanner.Scan() {   // internally, it advances token based on '\n'
		counter++
		if counter == 9{
			data[currChar] = art
			art = ""
			currChar++
			counter = 0
		}else if counter == 1{
			art = scanner.Text()
		}else{
			art = art + "\n" + scanner.Text()
		}
	}

	if str == "\\n"{
		for i:=0; i<8; i++{
			fmt.Println()
		}
		os.Exit(0)
	}
	
	//We need to tokenize the main string by '\n', and process tokens one by one
	strs := strings.Split(str, "\\n")

	for _, singleString := range strs{                    //For each token
		for i := 1; i < 9; i++{                       //Printing the output line by line
			for _, val := range singleString{
				PrintLine(data[val], i, &outputString)
			}
			outputString = outputString + "\n"	
		} 
	}
	_, err = output.WriteString(outputString)       //For ascii-art-web
	if err != nil{
		log.Fatal(err)	
	}
	return outputString	
}

//Prints particular line of the given banner
func PrintLine(str string, givenLine int, outputString *string){
	lineCounter := 1
	for _, char := range str{
		if char == '\n'{
			lineCounter++
			continue
		}
		if lineCounter == givenLine {
			*outputString = *outputString + string(char)
		}
	}
}

