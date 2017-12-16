
inf = open("part1.input","r")
data = inf.readlines()
import string
data = map(str.strip,data)
inf.close()

# print data

class ChildNode:
    def __init__(self, name, list_index):
        self.name = name
        self.list_index = list_index
    def __repr__(self):
        return self.name + str(self.list_index)

class Node:
    def __init__(self, name, weight, list_index):
        self.name = name
        self.children = []
        self.weight = weight
        self.list_index = list_index
        self.hasChildren = False
    def __repr__(self):
        return self.name + " - (" + str(self.weight) + ") " + str(self.children) + " list_index: " + str(self.list_index)
        
all_nodes = {}
list_nodes = []
list_index = 0

def parse_data_line(d):
    global list_index
    d = d.split()
    dlen = len(d)
    name = d[0]
    weight = d[1].split("(")[1]
    weight = int(weight.split(")")[0])
    #print d
    # print "name",name,"weight",weight
    n = Node(name,weight,list_index)
    all_nodes[name] = n
    list_index = list_index + 1

    if dlen!=2:
        n.hasChildren = True
        # we have children to deal with
        for c in  d[3:]:
            tmpname = c
            if c[-1] == ",":
                c = c[0:-1]
            n.children.append(ChildNode(c,-1))    

def lookup_all_children(n):
    print(n.name)
    for c in n.children:
        print("\t",all_nodes[c.name])

for d in data:
    parse_data_line(d)

def checknode0():
    n = all_nodes["gbyvdfh"]
    lookup_all_children(n)

def fill_out_child_list_indexes():
    for node_name in all_nodes.keys():
        n = all_nodes[node_name]
        # print "working on node",n
        for c in n.children:
            nc = all_nodes[c.name]
            c.list_index = nc.list_index

def sum_children(a):
    # print "sum_children is called",a
    sum = 0
    for c in a:
        if all_nodes[c.name].hasChildren:
            sum = sum + sum_child_branch(c.name)
        else:
            sum = sum + all_nodes[c.name].weight
    return sum

def sum_child_branch(child_name):
    # print "sum_child_branch called",child_name
    # print "myweight", all_nodes[child_name].weight
    return (all_nodes[child_name].weight +
        sum_children(all_nodes[child_name].children) )

def with_test_input():
    fill_out_child_list_indexes()
    root_name = "tknk"
    print( all_nodes[root_name])
    for c in all_nodes[root_name].children:
        w0 = sum_child_branch(c.name)
        print("w0 is",c.name, w0)

fill_out_child_list_indexes()
root_name = "qibuqqg"
print(all_nodes[root_name])
for c in all_nodes[root_name].children:
        w0 = sum_child_branch(c.name)
        print("w0 is",c.name, w0)

print("wrong program at root", all_nodes["dwggjb"])

level1_name = "dwggjb"
for c in all_nodes[level1_name].children:
    w0 = sum_child_branch(c.name)
    print("w0 is",c.name,w0)

print("wrong program at level1", all_nodes["wknuyhc"])

level2 = "wknuyhc"
for c in all_nodes[level2].children:
    w0 = sum_child_branch(c.name)
    print("w0 is",c.name,w0)

print("wrong program at level2", all_nodes["egbzge"])
level3 = "egbzge"
for c in all_nodes[level3].children:
    w0 = sum_child_branch(c.name)
    print("w0 is",c.name,w0)
