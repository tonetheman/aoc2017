from re import match
from collections import defaultdict

class treeNode():
    parent = None
    weight = 0
    children = []
    name = ""

    def __init__(self):
        self.weight = 0
        self.children = []
        self.name = ""
        self.parent = None

def getWeight(currentNode):
    totalWeight = currentNode.weight

    for nodes in currentNode.children:
        totalWeight += getWeight(nodes)

    return totalWeight

def getBranchesWeight(currentNode):
    branchesWeight = []

    for nodes in currentNode.children:
        branchesWeight.append(getWeight(nodes))

    return branchesWeight

def findWrongProgram(currentNode, dif = 0):
    branchesWeight = getBranchesWeight(currentNode)

    if (len(set(branchesWeight)) <= 1 and len(currentNode.parent.children) > 1):
        return currentNode, dif

    if (dif == 0):
        maxWeightOccurances = branchesWeight.count(max(branchesWeight))
        if (maxWeightOccurances == 1):
            dif = max(branchesWeight) - min(branchesWeight)
        else:
            dif = min(branchesWeight) - max(branchesWeight)
    if (dif > 0):
        branchNum = branchesWeight.index(max(branchesWeight))
    else:
        branchNum = branchesWeight.index(min(branchesWeight))

    branchToCheck = currentNode.children[branchNum]
    return findWrongProgram(branchToCheck, dif)



treeNodes = defaultdict(treeNode)

with open("part1.input") as inputData:

    for row in inputData:
        m = match("^(\w+) \((\d+)\)(?: -> )?(.*)$", row)
        prog = m.group(1)

        treeNodes[prog].weight = int(m.group(2))
        treeNodes[prog].name = prog

        for el in m.group(3).split(", "):
            treeNodes[prog].children.append(treeNodes[el])
            treeNodes[el].parent = treeNodes[prog]

root = treeNodes[prog]

while (root.parent is not None):
    root = root.parent
root.parent = root

part2branch, dif = findWrongProgram(root)

print("Part 1 answer: {}".format(root.name))
if (part2branch == root):
    print("Tree is perfectly balanced")
else:
    print("Part 2 answer: {} <{}>".format(part2branch.weight - dif, part2branch.name))