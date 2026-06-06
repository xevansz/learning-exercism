#include "armstrong_numbers.h"
#include<stdlib.h>
#include<math.h>

bool is_armstrong_number(int n){
    int result = n;
    int len=floor(log10(n) + 1);
    while(n){
        result -= pow(n % 10, len);
        n /= 10;
    }
return result == 0;
}