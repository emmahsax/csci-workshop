/**
 * Given a singly linked list, so where every node has just a value and a pointer to another node,
 * we want to determine if there are any cycles, so if any nodes point back to another node in the
 * linked list, creating a cycle with no ending node. Write a function that will take in a single node
 * that begins a linked list (or may not begin a linked list) and will return true if there is a cycle
 * and false if there is not.
 * Gotchas: a cycle may exist in the middle of a list, or betweeen the beginning and the end of a list
 *          this is solvable in O(n) time and O(1) space
 */

public class SingleLinkedListContainCycle {
  public static bool containsCycle(LinkedListNode firstNode) {
    int slowRunner = firstNode;
    int fastRunner = firstNode;

    while (slowRunner.next != null && fastRunner.next.next != null) {
      slowRunner = firstNode.next;
      fastRunner = firstNode.next.next;

      if (slowRunner == fastRunner) {
        return true;
      }
    }

    return false;
  }

  public static class LinkedListNode {
    public int value;
    public LinkedListNode next;

    public LinkedListNode(int value) {
      this.value = value;
    }
  }
}
