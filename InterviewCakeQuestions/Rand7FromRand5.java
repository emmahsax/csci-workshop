/**
 * Given a function rand5() that generates an integer from 1 to 5 with equal probability, write
 * a function rand7() that generates an integer from 1 to 7 with equal  probability.
 */

// Use rand5() twice to get the number of outcomes greater than 5. Then use fancy math (or the
// array commented out) to generate a number that is between 1 and 7. We use the last four values
// as "extraneous" because it makes the probabilities of each number uneven.
public class Rand7FromRand5 {
    public static int rand7() {
        // int[][] intArray = new int[][] {
        //   {1, 2, 3, 4, 5},
        //   {6, 7, 1, 2, 3},
        //   {4, 5, 6, 7, 1},
        //   {2, 3, 4, 5, 6},
        //   {7, -1, -1, -1, -1}
        // };

        while (true) {
        int rand1 = rand5() - 1;
        int rand2 = rand5() - 1;

        // this math is simulating the nested array commented out above
        int potentialOutcome = ((rand1 * 5) + rand2) + 1; // + 1 to not be "indexed at zero"

        if (potentialOutcome > 21) {
            continue;
        }

        return potentialOutcome % 7 + 1;
        }
    }
}
