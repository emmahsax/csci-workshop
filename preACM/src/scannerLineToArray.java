import java.util.Arrays;
import java.util.Scanner;

public class scannerLineToArray {

		public static void main(String[] args) {
			Scanner input = new Scanner(System.in);

			String sentence = input.nextLine();
			String[] wordArray = sentence.split(" ");

			System.out.println(Arrays.toString(wordArray));
			input.close();
		}
	}

/** HOW THIS WORKS:
 * This will turn the string into an array.
 *
 * [This, will, turn, the, string, into, an, array.]
 */

/** if using with ints, then do
 * int foo = Integer.parseInt(array[i]);
 */
