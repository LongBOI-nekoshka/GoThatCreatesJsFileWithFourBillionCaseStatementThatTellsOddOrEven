package main

import (
    // "bufio"
    // "fmt"
	"strconv"
    "os"
)


func main() {

	os.Mkdir("jsFile",0777)
	os.Remove("jsFile/fourBillionCase.js");
	f,err:= os.Create("jsFile/fourBillionCase.js");
	if(err != nil) {
		panic(err)
	}else {
		go zeroToEightHunderedM(f)
		go eightHunderedMToOnePointSix(f)
		go onePointSixToTwoPointFour(f)
		go twoPointFourToThreePointTwo(f)
		threePointTwoToFourB(f)
		// f.WriteString("export const oddOrEven = (n) => {")
		// f.WriteString("switch(n){")
		// for i := 0; i < 4000000000; i++ {
		// 	f.WriteString("case "+ strconv.Itoa(i) +":")
		// 	var switcher = i % 2
		// 	switch switcher{
		// 		case 0: 
		// 			f.WriteString("return" +  `"Even";`)
		// 		case 1:
		// 			f.WriteString("return" +  `"Odd";`)
		// 	}
		// }
		// f.WriteString("}")
		// f.WriteString("}")
	}

}

func zeroToEightHunderedM(f *os.File) {
	for i := 0; i <= 800000000; i++ {
		f.WriteString("case "+ strconv.Itoa(i) +":")
		var switcher = i % 2
		switch switcher{
			case 0: 
				f.WriteString("return" +  `"Even";`)
			case 1:
				f.WriteString("return" +  `"Odd";`)
		}
	}
}

func eightHunderedMToOnePointSix(f *os.File) {
	for i := 800000001; i <= 1600000000; i++ {
		f.WriteString("case "+ strconv.Itoa(i) +":")
		var switcher = i % 2
		switch switcher{
			case 0: 
				f.WriteString("return" +  `"Even";`)
			case 1:
				f.WriteString("return" +  `"Odd";`)
		}
	}
}

func onePointSixToTwoPointFour(f *os.File) {
	for i := 1600000001; i <= 2400000000; i++ {
		f.WriteString("case "+ strconv.Itoa(i) +":")
		var switcher = i % 2
		switch switcher{
			case 0: 
				f.WriteString("return" +  `"Even";`)
			case 1:
				f.WriteString("return" +  `"Odd";`)
		}
	}
}

func twoPointFourToThreePointTwo(f *os.File) {
	for i := 2400000001; i <= 3200000000; i++ {
		f.WriteString("case "+ strconv.Itoa(i) +":")
		var switcher = i % 2
		switch switcher{
			case 0: 
				f.WriteString("return" +  `"Even";`)
			case 1:
				f.WriteString("return" +  `"Odd";`)
		}
	}
}
func threePointTwoToFourB(f *os.File) {
	for i := 3200000001; i <= 4000000000; i++ {
		f.WriteString("case "+ strconv.Itoa(i) +":")
		var switcher = i % 2
		switch switcher{
			case 0: 
				f.WriteString("return" +  `"Even";`)
			case 1:
				f.WriteString("return" +  `"Odd";`)
		}
	}
	f.WriteString("}")
	f.WriteString("}")
}