/**
 * Write a function that, given a sentence with the position of a single opening parenthesis, finds the corresponding
 * closing parenthesis.
 * Gotchas: this can be done in O(n) time and O(1) additional space
 */

public class ParenthesisMatching {
    public static int match(String sentence, int index) {
        int parenCount = 0;

        for (int i = index + 1; i < sentence.length(); i++) {
        if (sentence.charAt(i) == '(') {
            parenCount++;
        } else if (sentence.charAt(i) == ')') {
            if (parenCount == 0) {
            return i;
            } else {
            parenCount--;
            }
        }
        }

        throw new IllegalArgumentException("No closing parenthesis");
    }
}
