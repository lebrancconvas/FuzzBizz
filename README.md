# FuzzBizz (Dark FizzBuzz) 

## Overview
FizzBuzz is the stupid game about people try to call the number, but if the number meets some condition, 
They will call it "Fizz" or "Buzz" or "FizzBuzz" instead. 

but it's too boring for me. Everywhere use FizzBuzz this, FizzBuzz that so decide to make "FuzzBizz" the dark side version 
of "FizzBuzz". 

## Traditional FizzBuzz
- If the number can't divide by 3 or 5 -> Return the number
- If the number can divide by 3 (but not divide by 5) -> Return "Fizz"
- If the number can divide by 5 (but not divide by 3) -> Return "Buzz"
- If the number can divide by both 3 and 5 -> Return "FizzBuzz"

### Example of Traditional FizzBuzz 
1 -> 1
2 -> 2
3 -> Fizz 
4 -> 4 
5 -> Buzz 
6 -> Fizz 
7 -> 7
8 -> 8
9 -> Fizz
10 -> Buzz
11 -> 11
12 -> Fizz
13 -> 13 
14 -> 14 
15 -> FizzBuzz 
16 -> 16 
17 -> 17 

## FuzzBizz (Dark FizzBuzz)
- If the number can divide by 3 or 5 -> Return the number
- If the number can't divide by 3 (but can divide by 5) -> Return "Fuzz"
- If the number can't divide by 5 (but can divide by 3) -> Return "Bizz"
- If the number can't divide by both 3 and 5 -> Return "FuzzBizz"

### Example of FuzzBizz (Dark FizzBuzz) 
1 -> FuzzBizz
2 -> FuzzBizz
3 -> Bizz 
4 -> FuzzBizz
5 -> Fuzz
6 -> Bizz
7 -> FuzzBizz
8 -> FuzzBizz
9 -> Bizz
10 -> Fuzz
11 -> FuzzBizz
12 -> Bizz
13 -> FuzzBizz
14 -> FuzzBizz
15 -> 15 
16 -> FuzzBizz 
17 -> FuzzBizz 
