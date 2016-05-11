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
