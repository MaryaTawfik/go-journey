package main
import (
	"fmt"
	"os"
)

func averageCalculator(Grades []float32 , )  {
	 var total float32

	for i := 0; i < len(Grades); i++ {
		total += Grades[i]
	}
	 var ave float32 = total/float32(len(Grades))
	fmt.Printf("total = %f\n",total)
	fmt.Printf("average= %f\n",ave)

}


func main() {
	name := ""
	var numberOfSubjects int

	fmt.Println("Enter your name")
	fmt.Scan(&name)

	fmt.Printf("Hey %s could you Enter number of the subject that you have taken\n",name)
	fmt.Scan(&numberOfSubjects)
	if numberOfSubjects <=0 {
		fmt.Println("invalid input")
		os.Exit(0)
	}
	studentGradeInfo := map[string]float32{}
	
	for i := 0; i < numberOfSubjects; i++ {
		subName :=""
		var score float32
		 fmt.Printf("enter the %d subject name\n",i+1)
		 fmt.Scan(&subName)
		//  fmt.Printf("tell me the scor that you hve in %s subName")
		//  fmt.Scan(&score)

	_, exists := studentGradeInfo[subName]
    for exists {
        fmt.Println("You already entered a value for that subject. Please enter a different name:")
        fmt.Scanln(&subName)
        _, exists = studentGradeInfo[subName] // Re-check the map with the new name
    }
	fmt.Printf("Tell me the score that you have in %s: ", subName)
    fmt.Scan(&score)

		 for score < 1 || score > 100 {
			fmt.Println("The score you entered is not valid. Enter a valid value in the range 1-100:")
			fmt.Scan(&score)

			
		 }
		
		 
		studentGradeInfo[subName]=score
		 
		 

		
	}
	var grades []float32
	fmt.Println("final grade report")
	for key,val := range studentGradeInfo {
		fmt.Printf("%s : %f\n",key,val)
		grades=append(grades, val)
	}

	averageCalculator(grades)
}