
def part1():
    def read_input_file(filename):
        data = open(filename).readlines()
        input_layers = []
        for i in range(len(data)):
            data[i] = data[i].strip()
            d= data[i].split(":")
            input_layers.append((int(d[0]), int(d[1])))
        return input_layers

    def find_layer(layernum,data):
        for a,b in data:
            if a==layernum:
                return (a,b)
        return None

    def calc_hit(period,step):
        m = (period-1)*2
        print("DBG: m",m)
        return step%m==0

    #data = read_input_file("test.input")
    data = read_input_file("part1.input")
    cost = 0
    #for i in range(6+1):
    for i in range(88+1):
        d = find_layer(i,data)
        if d is not None:
            if calc_hit(d[1],i):
                cost = cost + (i*d[1])
            print(i,d,calc_hit(d[1],i))
        else:
            print(i,d)
    print("final cost",cost)

################################
# part2

def read_input_file(filename):
    data = open(filename).readlines()
    input_layers = []
    for i in range(len(data)):
        data[i] = data[i].strip()
        d= data[i].split(":")
        input_layers.append((int(d[0]), int(d[1])))
    return input_layers

def find_layer(layernum,data):
    for a,b in data:
        if a==layernum:
            return (a,b)
    return None

def calc_hit(period,step,delay):
    m = (period-1)*2
    return (step)%m==0

def part2not():
	data = read_input_file("test.input")
	#data = read_input_file("part1.input")
	cost = 0
	delay = 0
	print("delay is",delay)
	for i in range(6+1):
	#for i in range(88+1):
	    d = find_layer(i,data)
	    if d is not None:
		if calc_hit(d[1],i,delay):
		    cost = cost + (i*d[1])
		print(i,d,calc_hit(d[1],i,delay))
	    else:
		print(i,d)
	print("final cost",cost)

def tonypart2():
	def calc(period,step):
		m = (period-1)*2
		return step%m == 0

	#data = read_input_file("test.input")
	data = read_input_file("part1.input")

	# i is the amount of delay
	for i in range(3000000,4000000):
		if i%1000==0:
			print(i,"*")
		cost = False
		# print("delay is",i,"start without cost",cost)
		
		# j here is taking the step across the scanners
		#for j in range(6+1):
		for j in range(88+1):
			d = find_layer(j,data)
			if d is not None:
				# print "layer",d, calc(d[1], i+j)
				if calc(d[1],i+j):
					cost = True
					break
		# print("delay end",i,"cost is",cost)

		if cost == False:
			print "FOUND ONE WITHOUT COST",i
			return			

tonypart2()


