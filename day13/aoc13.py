
DOWN = 1
UP = 2

class Layer:
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
    def cost(self):
        return self.depth * self.layer_number
    def step(self):
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
            if self.scanner_pos == 0:
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
    for layer in layers:
        layer.step()
        print(layer)

count = 0
person = -1 # the person is at layer -1 (not on the board)
total_cost = 0
for i in range(7):
    print "start of tick"
    print layers
    person = person + 1
    print "person moves into",person
    if len(layers[person].data)>0:
        if layers[person].data[0] ==9:
            print "COST here"
            total_cost = total_cost + layers[person].cost()
    print "tick happens"
    tick() # this moves the scanners
    print("-----")
    print()
    count = count + 1
    if count == 7:
        break
    

print "total cost is",total_cost