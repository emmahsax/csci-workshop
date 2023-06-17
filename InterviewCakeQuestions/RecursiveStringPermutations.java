/**
 * Write a recursive function that generates all permutations of a given string. Return the permutations as a set.
 * Your function can have loops. It just also must be recursive.
 */

public class RecursiveStringPermutations {
    public static Set<String> getPermutations(String input) {
        Set<String> permutationsSet = new HashSet<String>();

        if (input.length() <= 1) {
            return new HashSet<String>(Arrays.asList(input));
        }

        String inputExceptLastLetter = input.substring(0, inputString.length() - 1);
        char lastLetter = input.charAt(input.length() - 1);

        Set<String> permutationsExceptLastLetter = getPermutations(inputExceptLastLetter);

        for (String permutationSoFar : permutationsExceptLastLetter) {
            for (int i = 0; i <= permutationSoFar.length(); i++) {
                String permutation = permutationsSoFar.substring(0, i) + lastLetter + permutationsSoFar.substring(i);
                permutationsSet.add(permutation);
            }
        }

        return permutationsSet;
    }
}
