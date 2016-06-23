/** Write a function to check that a binary tree is a valid binary tree.
 * Here's a sample binary search tree:
 */

public static class BinaryTreeNode {
  public int value;
  public BinaryTreeNode left;
  public BinaryTreeNode right;
  public int upperBound;
  public int lowerBound;

  public BinaryTreeNode(int value) {
    this.value = value;
  }

  public BinaryTreeNode insertLeft(int leftValue) {
    this.left = new BinaryTreeNode(leftValue);
    this.left.upperBound = this.value;
    this.left.lowerBound = this.lowerBound;
    return this.left;
  }

  public BinaryTreeNode insertRight(int rightValue) {
    this.right = new BinaryTreeNode(rightValue);
    this.right.lowerBound = this.value;
    this.right.upperBound = this.upperBound;
    return this.right;
  }
}

/**
 * GOTCHAS: don't forget that even though a node may be in the correct placement according to its parent, it also must be in the correct place according to ALL of its ancestors
 *          this is solvable in O(n) time and additional space (additional space is O(lg n) if tree is balanced)
 */

public boolean validBinaryTree(BinaryTreeNode headNode) {
  Stack<BinaryTreeNode> nodeStack = new Stack<BinaryTreeNode>(); // stack of nodes that we're "counting" as we go
  nodeStack.push(headNode);

  while (!nodeStack.empty()) {
    BinaryTreeNode node = nodeStack.pop;

    if ((node.value < node.lowerBound) || (node.value > node.upperBound)) { // node doesn't fit within its set bounds
      return false;
    }

    if (node.left != null) {
      nodeStack.push(node.left);
    }

    if (node.right != null) {
      nodeStack.push(node.right);
    }
  }

  return true;
}
