package BinaryTreeJava;

class TreeNode {
    private int key;
    private TreeNode leftChild, rightChild;

    public TreeNode(int nodeKey) {
        this.key = nodeKey;
        this.leftChild = null;
        this.rightChild = null;
    }

    public int getKey() {
        return key;
    }

    public void setLeftChild(TreeNode node) {
        this.leftChild = node;
    }

    public void setRightChild(TreeNode node) {
        this.rightChild = node;
    }

    public TreeNode getLeftChild() {
        return this.leftChild;
    }

    public TreeNode getRightChild() {
        return this.rightChild;
    }
}

public class BinaryTree {
    public static TreeNode root = null;

    public static TreeNode newTreeNode(int key) {
        TreeNode node = new TreeNode(key);
        return node;
    }

    public static void insertTreeNode(TreeNode node, TreeNode newNode) {
        if (newNode.getKey() < node.getKey()) {
            if (node.getLeftChild() != null) {
                insertTreeNode(node.getLeftChild(), newNode);
            } else {
                node.setLeftChild(newNode);
            }
        } else {
            if (newNode.getRightChild() != null) {
                insertTreeNode(node.getRightChild(), newNode);
            } else {
                node.setRightChild(newNode);
            }
        }
    }

    public static void createBinaryTree(int[] nodes) {
        for (int key : nodes) {
            TreeNode newNode = new TreeNode(key);
            if (root == null) {
                root = newNode;
            } else {
                insertTreeNode(root, newNode);
            }
        }
    }

    public static void echoNode(TreeNode node){
        System.out.println("树节点key:" + node.getKey());
        System.out.println("树节点leftChild:" + node.getLeftChild());
        System.out.println("树节点rightChild:" + node.getRightChild());
    }

    public static void callback(int key) {
        System.out.println(key);
    }

    public static void inOrderBinaryTree(TreeNode node) {
        if (node != null) {
            inOrderBinaryTree(node.getLeftChild());
            // callback(node.getKey());
            echoNode(node);
            inOrderBinaryTree(node.getRightChild());
        }
    }

    public static void main(String[] args) {
        int[] nodes = { 8, 3, 10, 1, 6, 6, 14, 4, 7, 13 };
        System.out.println("构建二叉树");
        createBinaryTree(nodes);
        System.out.println("遍历二叉树");
        inOrderBinaryTree(root);
    }
}

/**
 *                    [8]
 *               [3]      [10]
 *            [1]   [6]       [14]
 *                [4] [6]   [13]
 *                      [7] 
 * 
 */