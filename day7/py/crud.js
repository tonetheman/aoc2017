const fs = require('fs');

/*
FOUND THIS SOLUTION ON REDDIT and 
- running it to compare to mine
*/

fs.readFile('..\\part1.input', 'utf8', (err, data) => {
//fs.readFile(__dirname + '..\\part1.input', 'utf8', (err, data) => {
    data = data.trim();
    const disks = {};
    data.split('\n').forEach((line) => {
        const parts = line.split(' -> ');
        const disk = parts[0].split(' ');
        const name = disk[0].trim();
        disks[name] = {
            value: Number(disk[1].substr(1, disk[1].indexOf(')') - 1)),
            children: [],
            total: 0,
        };
        if(parts.length > 1) {
            disks[name].children = parts[1].split(',').map((x) => x.trim());
        }
    });

    const root = getRoot(disks);
    sumWeights(root, disks);
    console.log(balance(root, disks));
});

function getRoot(disks) {
    const keys = new Set(Object.keys(disks));
    for(const key in disks) {
        for(const i in disks[key].children) {
            keys.delete(disks[key].children[i]);
        }
    }

    return keys.values().next().value;
}

function sumWeights(root, tree) {
    tree[root].total = tree[root].value;
    tree[root].children.forEach((c) => {
        tree[root].total += sumWeights(c, tree);
    });

    return tree[root].total;
}

function balance(root, tree, target) {
    const children = {};
    var newTarget;
    tree[root].children.forEach((c) => {
        if(children[tree[c].total] === undefined) {
            children[tree[c].total] = c;
        } else {
            children[tree[c].total] = false;
            newTarget = tree[c].total;
        }
    });
    for(const i in children) {
        if(children[i]) {
            return balance(children[i], tree, newTarget);
        }
    }

    return tree[root].value + target - tree[root].total;
}