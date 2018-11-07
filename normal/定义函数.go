package main

func test03(a int,b int)(sum int){
	sum=a+b
return sum
}

type nihao func(a int,b int)(sum int){
	var wode nihao
	var s int
	wode=test03

	s=wode(10,10)
	fmt.Println(s)

}
