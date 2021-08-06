package main

import "fmt"


func Draw(point []string){
	var pole string
	var n int16
	n = -1
	for i:=0; i<9; i++{
		n += 1
		if (i % 3 == 0) && (i!=0){
			pole += "|\n"
		}
			pole += "|" + point[n]	
	}

	pole += "|"
	fmt.Println(pole)
}

func main(){
	points := []string{"1","2","3","4","5","6","7","8","9"}
	
	Draw(points)

	var action int16
	turn := 1

	for i:=0; i<9; i++{
		fmt.Scanln(&action)

		if turn % 2 != 0{
			points[action-1] = "X"
			turn++
		}else{
			points[action-1] = "O"
			turn++
		}
		Draw(points)
	}
}


