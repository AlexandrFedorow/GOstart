package main

import "fmt"


func Draw(point []string, part []string){
	var pole1 string
	var pole2 string

	var n int16
	n = -1
	for i:=0; i<9; i++{
		n += 1
		if (i % 3 == 0) && (i!=0){
			pole1 += "|\n"
		}
			pole1 += "|" + point[n]
	}

	n = -1

	for i:=0; i<9; i++{
		n += 1
		if (i % 3 == 0) && (i!=0){
			pole2 += "|\n"
		}
			pole2 += "|" + part[n]
	}

	pole1 += "|"
	pole2 += "|"

	fmt.Println(pole1)
	fmt.Println("\n")
	fmt.Println(pole2)
}

func main(){
	points := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	parts := []string{"-", "-", "-", "-", "-", "-", "-", "-", "-"}
	Draw(points, parts)

	var action int16
	turn := 1

	for i:=0; i<9; i++{
		fmt.Scanln(&action)

		if action == 0{
			break
		}

		if (action > 0) && (action < 10){
			if turn % 2 != 0{
				parts[action-1] = "X"
				turn++
			}else{
				parts[action-1] = "O"
				turn++
			}
		}else{
			fmt.Println("huita")
		}
		Draw(points, parts)
	}
}
