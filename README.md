This program solves the following puzzle: Cover a rectangular area with tiles of the form 

	X X X X
	  X

that can be rotated and flipped over. You can run it with two numerical arguments

	./pento 5 10

to try to fill a rectangular area 5 x 10, and it will print all solutions, or nothing at all if no solution can be found.

Note that (5,10) is the smallest rectangle where solution exists, and in general it seems that one 
of the sides should be divisible by 5 and the other by 10. (That one side should be divisible by 5 follows
from the fact that the area of the rectangle should be divisible by 5.)