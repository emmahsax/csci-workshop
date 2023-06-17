/**
 * Write an efficient function that checks whether any permutation of an input string is a palindrome. For example, "civic" should return
 * true, "ivicc" should return true, "civil" should return false, and "livci" should return false.
 * Gotchas: this is computable in O(n) time
 */


// Our solution here is to keep track of all of the characters that have appeared an odd amount of times.
// We don't care what orders the letters are in, so as long as only one letter or less appears an odd amount of times
// (to be the center of a palindrome), we don't care about the other letters.
import java.util.Set;
import java.util.HashSet;

public class PermutationPalindrome {
    public boolean hasPalindromePermutation(String input) {
        Set<Character> unpairedCharacters = new HashSet<Character>();

        for (char letter : input.toCharArray()) {
            if (unpairedCharacters.contains(letter)) {
                unpairedCharacters.remove(letter);
            } else {
                unpairedCharacters.add(letter);
            }
        }

        return unpairedCharacters.size() <= 1;
    }
}
