func groupAnagrams(strs []string) [][]string {
	hashSet:=make(map[string]*[]string)
	for _,e:=range strs{
		s:=[]int32(e)
		sort.SliceStable(s,func(i,j int)bool{
			return s[i]<s[j]
		})
		sortedS:=[]string{}
		for _,i:=range s{
			sortedS=append(sortedS,strconv.Itoa(int(i)))
		}
		newS:=strings.Join(sortedS,"+")
		if hashSet[newS]==nil{
			newLine:=make([]string,0)
			newLine=append(newLine,e)
			hashSet[newS]=&newLine
		}else{
			*hashSet[newS]=append(*hashSet[newS],e)
		}
	}
	result:=make([][]string,0)
	for _,i:=range hashSet{
		result=append(result, *i)
	}
	return result
}
