# tonys advent of code 2017
-------------------------

## Day 12 finished
I wrote a lot and decided I had forgot my plan and deleted most of it.
Once I did that the problem was not bad. I did have to come up with my own stack implementation. Still gross.

Second part of the problem was a generalization of the first part.

## Day 12
Problem seems easy. Go does not have a built in stack? I found an issue that bothers me about go. I could rely on some rando from the internet? Maybe that is the intention. OR I could on myself. I trust neither the random nor me to write the stack correctly.


## Day 11
Fairly easy today.

Mainly from other experience I have had. I have written games (or failed prototypes) with hex coordinates before. So the resource was familar to me as was the subject.

Nothing to crazy.

## Day 10
The first part made sense and I did it fairly quickly.

The second part of the problem was written really badly. The first example was completely unclear where the numbers were coming from. The changes for the algorithm were easy (just moved two variables to globals gross).

Once I finally decided I understood the wording the problem was not hard.

## Day 9 (part2)
Not hard at all since I made some good choices on part 1.

Only had to add a counter into the place where I was counting garbage characters.

## Day 9 (part1)
First part of day 9 was more involved than I had originally thought. I started to try out antlr for go but decided it would take me longer to figure out antlr than just write it myself.

The first version had the parsing all over the place. I took some time and split the parsing of the garbage from the parsing of the groups. Once I realized (aka read the problem) that I was going to need to keep up with the groups I added a stack. Got the answer on the first try once all the tests passed.

## Day 8
This type of puzzle is easy for me. The main mistake I made was I did not know all the cases until I ran the puzzle input. The code I wrote has some duplication in it so that could be fixed.

Part 2 was just a variation on part1 with a running max value. Also not a problem.

Mainly learned more about switch/case in go.

## Day 7
Part1 was not bad. Spent a lot of time parsing input.

Found some weirdness with golang. When I was accessing a value inside a struct that was inserted into a map, I could not change the value. I had to copy the struct change the value then reinsert into the map. Might be how maps/structs work in go but it was not intuitive at all.

The algorithm I went with came from me staring at the input. I removed anyone who was listed as a child from the running and one node popped out as the answer. It worked.

## Day 6
More code than I thought I would need. Meh.

Had to look up map syntax for go again. The way to test for existing items in the map does not make a lot of sense to me so I have to look up the syntax everytime.

The way I solved part1 made part2 super simple.

## Day 5
Cool problem. Reminds me of the machine from last year AOC. Just jmp instructions in this case. I again stole code from Day 2 for reading files and scanning in golang. I feel like it will take me a while to really remember how to read a file in golang line by line.

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
