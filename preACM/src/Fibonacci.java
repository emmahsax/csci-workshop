public class Fibonacci {
    public static HashMap<Integer,Integer> prevCalculatedFibs = new HashMap<Integer,Integer>();

    public static void main (String[] args) {
        // Some examples of using the method fib()
        System.out.println("fib(1): " + fib(1)); // should print 1
        System.out.println("fib(2): " + fib(2)); // should print 1
        System.out.println("fib(5): " + fib(5)); // should print 5
        System.out.println("fib(6): " + fib(6)); // should print 8
    }

    public static int fib(int input) {
        if (input == 1 || input == 2) { // these are our base cases
            return 1;
        } else {
            int ans = prevCalculatedFibs.get(input);
            if (ans != null) { // we have seen this input already
                return ans;
            } else { // we haven't seen this input already
                int calculated_ans = fib(input - 1) + fib(input - 2);
                prevCalculatedFibs.put(input, calculated_ans);
                return calculated_ans;
            }
        }
    }
}
