
DOWN = 1
UP = 2

class Layer:
    def __init__(self,depth):
        self.depth = depth
        self.direction = DOWN
        self.data = []
        for i in range(self.depth):
            self.data.append(0)
        if self.depth>0:
            self.data[0] = 9
        self.scanner_pos = 0
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
    Layer(3),
    Layer(2),
    Layer(0),
    Layer(0),
    Layer(4),
    Layer(0),
    Layer(4)
    ]

def tick():
    for layer in layers:
        layer.step()
        print(layer)

count = 0
for i in range(100):
    tick()
    print("-----")
    count = count + 1
    if count == 3:
        break
    

