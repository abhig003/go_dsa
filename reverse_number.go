//leetcode 7 reverse number problem


func reverse(x int) int {

max := math.MaxInt32
min := math.MinInt32
var ans int=0
for x!=0 {
  
   digit:=x%10;
   if ans*10 >=max || ans*10 <=min {
    return 0;
   }
   ans=ans*10+digit
   x=x/10;
}
return ans;
}