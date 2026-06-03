#include "isogram.h"

#define size  ('z' - 'a' + 1)
/*
bool is_isogram(const char phrase[]){
    if ( phrase == NULL)
        return false;
    
    bool char_seen[size] = {false};
    for (char c = *phrase; c!= '\0'; c = *++phrase ){
        if ( c == ' ' || c == '-')
            continue;
    
        unsigned int char_idx = tolower(c) - (int)('a');
        if ( char_seen[char_idx])
            return false;
    
        char_seen[char_idx] = true;
    }
    return true;
}
*/
bool is_isogram(const char *phrase)
{
    if (!phrase)
        return false;
    bool seen[size] = {false};
    for (; *phrase; ++phrase)
    {
        unsigned char c = *phrase;
        if (!isalpha(c))
            continue;
        int idx = tolower(c) - 'a';
        if (seen[idx])
            return false;
        seen[idx] = true;
    }
    return true;
}
