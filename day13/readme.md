Some thoughts

PART2
----------------------------
calc_hit NOT CALLED
p0
delay=1
person=-1
[S 0] 
[S 0 0]

calc_hit(period=2,step=0,delay=1)
p1
delay=1
person=0
*[(0) S]
[0 S 0]

calc_hit(period=3,step=1,delay=1)
p2
delay=1
person=1
[S 0]
*[(0) 0 S]




PART1
-------------------------

Maybe I was too quick to dismiss the math
- Can I use this?

2 slots -> pico%2==0 caught = true
3 slots -> pico%4==0 caught = true
4 slots -> pico%6==0 caught = true


s0
pico0 = 0
pico1 = 1
pico2 = 2
pico3 = 1
pico4 = 0
pico5 = 1
pico6 = 2
pico7 = 1

pico%4== 0 -> caught
pico

pico = 0
s1 = 0

pico = 1
s2 = 1

pico = 2
s1 = 0

so formula for s1 is if pico%2==0 caught=true

s4
p0 = 0
p1 = 1
p2 = 2
p3 = 3
p4 = 2
p5 = 1
p6 = 0
p7 = 1
p8= 2
p9=3
p10=2
p11=1
p12=0
pico%6==0 caught = true