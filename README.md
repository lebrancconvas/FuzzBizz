# FuzzBizz (Dark FuzzBizz) 

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

## FuzzBizz
- If the number can divide by 3 or 5 -> Return the number
- If the number can't divide by 3 (but can divide by 5) -> Return "Fuzz"
- If the number can't divide by 5 (but can divide by 3) -> Return "Bizz"
- If the number can't divide by both 3 and 5 -> Return "FuzzBizz"  
