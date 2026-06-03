#include "armstrong_numbers.h"
#include<math.h>
bool is_armstrong_number(int n){
    int result = 0;
    int len=(log10(n) + 1);
    for ( result = n; n!=0; n /= 10){
         result -= pow( n % 10, len);
    }
return result == 0;
}