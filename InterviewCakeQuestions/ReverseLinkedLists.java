/**
 * Given a singly linked list, so where every node has just a value and a pointer to another node,
 * write a function that will reverse this linked list completely in place, and will return the new
 * head of the reversed linked list.
 * Gotchas: the order in which you reconnect pieces is important
 *          watch out for your edge cases (when there are 0, 1, or 2 nodes in the list)
 *          this is solvable in O(n) time and O(1) space
 */

public class ReverseLinkedLists {
    public static Node reverseLinkedList(Node head) {
        Node previousNode = null;
        Node currentNode = head;
        Node nextNode = null;

        while (currentNode != null) {
            nextNode = currentNode.next;
            currentNode.next = previousNode;
            previousNode = currentNode;
            currentNode = nextNode;
        }

        return previousNode;
    }

    public static class Node {
        public int value;
        public Node next;

        public Node(int value) {
        this.value = value;
        }
    }
}
