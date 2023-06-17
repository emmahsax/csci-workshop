/**
 * Use the built-in stack method to create a new class MaxStack with a function
 * getMax() that returns the largest element in the stack. getMax() should not
 * remove the item.
 * Gotchas: this is solvable in O(1) time for push(), pop(), and getMax()
 */

import java.util.Stack;

public class MaxStack {
    Stack<Integer> stack = new Stack<Integer>();
    Stack<Integer> maxsStack = new Stack<Integer>();

    public int push(int element) {
        stack.push(element);
        if (maxsStack.empty() || element >= (int) maxsStack.peek()) {
        maxsStack.push(element);
        }
        return element;
    }

    public int pop() {
        int element = (int) stack.pop();
        if (element == (int) maxsStack.peek()) {
        maxsStack.pop();
        }
        return element
    }

    public int getMax() {
        return (int) maxsStack.peek();
    }
}
