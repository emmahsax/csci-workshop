/**
 * Given a function rand(7) that generates an integer from 1 to 7 with equal probability, write
 * a function rand(5) that generates an integer from 1 to 5 with equal  probability.
 */

// This works because we make no gurantees about run-time or about what the probability
// of each number should be. So with this solution, it does that the chance of never returning,
// and the probability that the answer will be any given number is 1/7.
public class Rand5FromRand7 {
  public static int rand5() {
    int ans = rand7();

    while (ans > 5) {
      ans = rand7();
    }

    return ans;
  }
}
