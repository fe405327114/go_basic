package main
//利用Map
func useMap(str string)string {
	result:=""
	temp:=map[string]byte{}
	for _,v:=range str{
		length:=len(temp)
		//此处不重复的话就新添加key，key唯一
		temp[string(v)]='a'
		if len(temp)!=length{
			result+=string(v)
		}
	}
	return  result
}
