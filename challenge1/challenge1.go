/*
https://adventofcode.com/2025/day/1

Need a password
Theres a safe with a dial that goes from 0-99
Puzzle input = a sequence of rotations
Each 'rotation' has both a distance and direction
A rotation to the left is a 'negative' value in the context of 0 through 99, and a rotation to the right is positive
- I.e 25 + L5 = 20; 25 + R1 = 26
Since the dial is a circle, a rotation to the left going past 0 sets it to 99, and vice-versa
The dial starts at position 50
The password (flag) is the number of times the dial is left pointing at 0 after any rotation in the sequence
- I.e, how many times is the 'dial' value 0?

TODO:
> Save input data as a file so it can be ingested to the program (data is very large)
> Declare some 'dial' integer variable initialized to 50
> Declare some 'dial_zero_freq' integer variable initialized to 0, used to track globally the # of times the dial is at 0
> Declare some function 'adjust_dial' to handle modifying the value of dial.
* Value can only be from 0 to 99. Handle logic for going past 0 or 99 like a dial. Params are direction and magnitude
* After each time the dial is modified, check if the dial is at 0. If it is, increment 'dial_zero_freq'
> Open the file to begin reading data from it (presumably, not sure how Golang handles this)
> Iterate through the input data, passing each line as an input to 'adjust_dial'. Increment 'dial_zero_freq' as needed
> Close the file
> Return 'dial_zero_freq'
*/

package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
