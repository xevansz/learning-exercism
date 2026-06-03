#include "armstrong_numbers.h"
#include<stdlib.h>
#include<math.h>

bool is_armstrong_number(int n){
    int result = n;
    int len=(log10(n) + 1);
    while(n){
        result -= pow(n % 10, len);
        n /= 10;
    }
return result == 0;
}

/*
bool is_armstrong_number(int n){
    int result = 0;
    int remainder=0;
    while(n != 0){
        remainder = n % 10;
        result += remainder * remainder * remainder;
         n = n/10;
    }
return result == n;
}
*/