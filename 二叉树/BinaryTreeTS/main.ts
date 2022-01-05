class TreeNode {
    key: number;
    leftChild: TreeNode | any;
    rightChild: TreeNode | any;
    constructor(key: number) {
        this.key = key;
        this.leftChild = null;
        this.rightChild = null;
    }
}

class BinaryTree {
    root: TreeNode | any;
    nodes: number[];
    result: number[];
    constructor(data: number[]) {
        this.root = null;
        this.nodes = data;
        this.result = [];
    }
    insertNode = (node: TreeNode, newNode: TreeNode) => {
        if (newNode.key < node.key) {
            if (node.leftChild === null) {
                node.leftChild = newNode;
            } else {
                this.insertNode(node.leftChild, newNode);
            }
        } else {
            if (node.rightChild === null) {
                node.rightChild = newNode;
            } else {
                this.insertNode(node.rightChild, newNode);
            }
        }
    }
    insert = (key: number) => {
        const newNode = new TreeNode(key);
        if (this.root === null) {
            this.root = newNode;
        } else {
            this.insertNode(this.root, newNode);
        }
    }
    inOrderTraverseNode = (node: TreeNode | any, callback: Function) => {
        if (node !== null) {
            this.inOrderTraverseNode(node.leftChild, callback);
            callback(node.key);
            this.inOrderTraverseNode(node.rightChild, callback);
        }
    }
    inOrderTraverse = (callback: Function) => {
        this.inOrderTraverseNode(this.root, callback);
    }
    callback = (key: number) => {
        this.result.push(key);
    }

    run = () => {
        this.nodes.forEach(key => {
            this.insert(key);
        })
        // console.log(this.root);
        this.inOrderTraverse(this.callback)
        console.log(this.result.join("->"))
    }
}

function main() {
    const nodes: number[] = [8, 3, 10, 1, 6, 6, 14, 4, 7, 13]
    const binaryTree = new BinaryTree(nodes);
    binaryTree.run();
}
main();