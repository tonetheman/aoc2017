# tonys advent of code 2017
-------------------------

## Day 2
Also in go. Interesting thing here is IO for me. I am still new enough to go and things like dealing with files are still not natural.

I had to look up the Scanner interface.

The other thing that threw me about the problem was that my input was different for the test vs the problem. The test was copy/pasted from the page and was space delimited. The problem input was also copy/pasted but was tab delimited.

The strings.Split function messed up depending on the input. So I changed that manually for testing vs the final problem.


## Day 1
First day. Decided to use go. The solution is straight forward. No real oddness here. The only thing that was interesting was when I solved the first part I was using one technique to do it: mainly +1 to get the next item and then an extra if to handle the wrap around case.

Once I made it to part2 I rethought part1 to include a mod so there was no if statement.

I later added a functional-ish looking python and javascript version just to see if I could do it. The python version runs out of stack space. The JS one does what it should but is gross to read.
