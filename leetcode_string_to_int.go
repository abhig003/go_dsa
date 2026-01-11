//problem string to int


func myAtoi(s string) int {

    int_min:= math.MinInt32
    int_max:=math.MaxInt32
    
    length:=len(s);
    i:=0
    for i<length && s[i]==' '{
        i=i+1
    }
    var pos_sign_digit bool=true// consider as positve number
    if i<length &&(s[i]=='+' || s[i]=='-'){// 
        if s[i]=='-'{ // if the number is negative then marks as neagtive
         pos_sign_digit=false
        }
        i=i+1;
    }
    //till now here we have skipped the space and indentify the the sign of number
    ans:=0
    for i<length && s[i]>='0' && s[i]<='9'{

        digit:=int(s[i]-'0');
        if ans > (int_max-digit)/10{

            if(!pos_sign_digit){
                return int_min
            }
            return int_max;
        }
        ans = ans*10 + digit
        i=i+1;

    // we have need to check the overflow
 

    }
    if !pos_sign_digit {
        return -ans
    }
return ans;





}