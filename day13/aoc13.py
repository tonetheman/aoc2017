
DOWN = 1
UP = 2

class Layer:
    def reset(self):
        self.direction = DOWN
        self.data = []
        for i in range(self.depth):
            self.data.append(0)
        if self.depth>0:
            self.data[0] = 9
        self.scanner_pos = 0        
    def __init__(self,layer_number,depth):
        self.layer_number = layer_number
        self.depth = depth
        self.direction = DOWN
        self.data = []
        for i in range(self.depth):
            self.data.append(0)
        if self.depth>0:
            self.data[0] = 9
        self.scanner_pos = 0
        #if layer_number==0:
        #    print("dbg:Layer ctor",self.scanner_pos)
    def cost(self):
        return self.depth * self.layer_number
    def step(self):
        #if self.layer_number==0:
        #    print("dbg:Layer step called",self.depth,self.direction)
        if self.depth==0:
            return
        if self.direction == DOWN:
            if self.scanner_pos == self.depth-1:
                self.direction = UP
                self.data[self.scanner_pos] = 0
                self.scanner_pos = self.depth-2
                self.data[self.scanner_pos] = 9
            else:
                self.data[self.scanner_pos] = 0
                self.scanner_pos = self.scanner_pos + 1
                self.data[self.scanner_pos] = 9
        elif self.direction == UP:
            #if self.layer_number==0:
            #    print("dbg:Layer dir is up",self.scanner_pos)
            if self.scanner_pos == 0:
                #print("dbg:Layer dir isup and scanner pos is 0")
                self.direction = DOWN
                self.data[self.scanner_pos] = 0
                self.scanner_pos = 1
                self.data[self.scanner_pos] = 9
            else:
                self.data[self.scanner_pos] = 0
                self.scanner_pos = self.scanner_pos-1
                self.data[self.scanner_pos] = 9
                
                
            
    def __repr__(self):
        return str(self.data)
    
layers = [
    Layer(0,3),
    Layer(1,2),
    Layer(2,0),
    Layer(3,0),
    Layer(4,4),
    Layer(5,0),
    Layer(6,4)
    ]

def tick():
    count = 0
    for layer in layers:
        layer.step()
        #if count == 0:
        #    print("dbglayer0",layer)
        count = count + 1

def test_case(delay):
    for l in layers:
        l.reset()
    count = 0
    person = -1 # the person is at layer -1 (not on the board)
    total_cost = 0
    max_loop = delay+ 7
    got_caught = False
    for i in range(max_loop):
        print "start of tick",count
        # print layers

        if delay==0:
            person = person + 1
            #print "person moves into",person,layers[person]
            if len(layers[person].data)>0:
                if layers[person].data[0] ==9:
                    #print "COST here"
                    got_caught = True
                    total_cost = total_cost + layers[person].cost()

        #print("tick happens")
        tick() # this moves the scanners
        #print("-----")
        #print()

        #mark down the dely
        if delay!=0:
            print "not putting person in new delay is",delay,delay-1
            delay = delay -1

        count = count + 1
        if count == max_loop:
            print "breaking for count"
            break
        

    print "total cost is",total_cost
    print ""
    return (got_caught,total_cost)

def read_input_file(filename):
    data = open(filename,"r").readlines()
    input_layers = []
    for i in range(len(data)):
        data[i] = data[i].strip()
        d= data[i].split(":")
        input_layers.append(Layer(int(d[0]), int(d[1])))
    return input_layers

def part1():
    input_layers = read_input_file("part1.input")
    # find max layer number in the input
    max_layer_num = 0
    for l in input_layers:
        if l.layer_number > max_layer_num:
            max_layer_num = l.layer_number
    print("max layer number is",max_layer_num)
    # now get the layers setup like the test case
    layers =[]
    # routine that will find a layer if read from input
    def find_input_layer(i):
        for l in input_layers:
            if l.layer_number==i:
                return l
        # if you do not find one return None
        return None
    for i in range(max_layer_num+1):
        l = find_input_layer(i)
        if l is  None:
            layers.append(Layer(i,0))
        else:
            layers.append(l)
    # now layers looks like the test case
    count = 0
    person = -1 # the person is at layer -1 (not on the board)
    total_cost = 0
    max_loop = delay+ 89
    for i in range(89):
        print "start of tick"
        print layers
        person = person + 1
        print("person moves into",person)
        if len(layers[person].data)>0:
            if layers[person].data[0] ==9:
                print "COST here"
                total_cost = total_cost + layers[person].cost()
        print("tick happens")
        tick() # this moves the scanners
        print("-----")
        print()
        count = count + 1
        if count == 89:
            break
        

    print("total cost is",total_cost)

def part2(delay,dbg):
    input_layers = read_input_file("part1.input")
    # find max layer number in the input
    max_layer_num = 0
    for l in input_layers:
        if l.layer_number > max_layer_num:
            max_layer_num = l.layer_number
    if dbg:
        print("max layer number is",max_layer_num)
    # now get the layers setup like the test case
    layers =[]
    # routine that will find a layer if read from input
    def find_input_layer(i):
        for l in input_layers:
            if l.layer_number==i:
                return l
        # if you do not find one return None
        return None
    for i in range(max_layer_num+1):
        l = find_input_layer(i)
        if l is  None:
            layers.append(Layer(i,0))
        else:
            layers.append(l)
    # now layers looks like the test case
    count = 0
    person = -1 # the person is at layer -1 (not on the board)
    total_cost = 0
    got_caught = False
    max_loop = delay+ 89
    for i in range(max_loop):
        if dbg:
            print("start of tick")
            print layers

        if delay==0:
            person = person + 1
            if dbg:
                print("person moves into",person,layers[person])
            if len(layers[person].data)>0:
                if layers[person].data[0] ==9:
                    if dbg:
                        print("COST here")
                    got_caught = True
                    total_cost = total_cost + layers[person].cost()

        if dbg:
            print("person moves into",person)
        
        if dbg:
            print("tick happens")
        
        tick() # this moves the scanners
        
        if dbg:
            print("-----")
            print()

        #mark down the dely
        if delay!=0:
            if dbg:
                print("not putting person in new delay is",delay,delay-1)
            delay = delay -1
        
        count = count + 1
        if count == max_loop:
            break
        
    if dbg:
        print("total cost is",total_cost)
    
    return (got_caught,total_cost)

def testcase_part2():
    for i in range(11):
        caught,cost = test_case(i)
        if not caught:
            print "GOT IT",i
            break

def real_part2():
    for i in range(30000,40000):
        caught,cost = part2(i,False)
        if not caught:
            print "GOT IT",i
            break

# testcase_part2()
real_part2()
