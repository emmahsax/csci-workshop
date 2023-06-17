/**
 * Implement a queue with 2 stacks. Your queue should have an enqueue and a
 * dequeue function and it should be "first in, first out". Optimize for
 * the time cost of m function calls on your queue. These can be any mix
 * of enqueue and dequeue calls. Assume you already have a stack implementation
 * and it gives O(1) time for push() and pop().
 * Gotchas: this is solvable in O(m) runtime for m function calls
 */

import java.util.Stack;

public class QueueWithTwoStacks {
    Stack<Integer> inStack = new Stack<Integer>();
    Stack<Integer> outStack = new Stack<Integer>();

    public static void enqueue(int element) {
        inStack.push(element);
    }

    public static int dequeue() {
        if (outStack.empty()) {
        transferInStackToOutStack();
        }

        return outStack.pop();
    }

    public static void transferInStackToOutStack() {
        while (!inStack.empty()) {
        int element = inStack.pop();
        outStack.push(element);
        }
    }
}
