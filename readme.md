# tonys advent of code 2017
-------------------------

## Day 3 (revisited)
After some thought today and talking to some other folks, I was able to finish day 3.

A lot of it was math for part 1. I came up with formulas for the bottom right corner for each layer of the spiral. Once I got that corner it was clear this value bounded the target. Meaning if you ever find a square bigger than your target then you have found the layer the target is on.

Once I did that I came up with another set of formula for each mid point on the 4 sides of each layer. Not as hard as it sounds. Then it is a simple determination of which side the target is on and which side it is closer to.

For the second part of the problem I just built the spiral. It seemed easier. I used the recursive definition and got the number on the first try.

## Day 4
Much easier today.

Used code from day 2 to read in a file and a text scanner also from day 2.

The only golang thing I did different was I tried to use a range for walking the map and some of the for-loops. I messed up the syntax a few times but it makes sense to me. I need to keep using it.

## Day 3
Gross gross day.

I could not solve this one. So early in the cycle and I cannot solve one. Hopefully this does not mean the others will be harder.

I worked out some of the math on a spreadsheet for this problem. I can see the right bottom corner is related to a square root. I think I have worked out which layer the target number is on.

Next I have to determine how far off 0 it is. Sure there is some math I can do there too. But it has escaped me so far.

Horrible for a problem to be so hard so early.

## Day 2
Also in go. Interesting thing here is IO for me. I am still new enough to go and things like dealing with files are still not natural.

I had to look up the Scanner interface.

The other thing that threw me about the problem was that my input was different for the test vs the problem. The test was copy/pasted from the page and was space delimited. The problem input was also copy/pasted but was tab delimited.

The strings.Split function messed up depending on the input. So I changed that manually for testing vs the final problem.


## Day 1
First day. Decided to use go. The solution is straight forward. No real oddness here. The only thing that was interesting was when I solved the first part I was using one technique to do it: mainly +1 to get the next item and then an extra if to handle the wrap around case.

Once I made it to part2 I rethought part1 to include a mod so there was no if statement.

I later added a functional-ish looking python and javascript version just to see if I could do it. The python version runs out of stack space. The JS one does what it should but is gross to read.
