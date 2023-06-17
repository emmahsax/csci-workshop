import java.util.Scanner;

public class SimonSays {

	public static void main(String[] args) {
		Scanner input = new Scanner(System.in);

		int caseAmount = input.nextInt();
		input.nextLine();

		String compare = "Simon says";

		for (int i = 0; i < caseAmount; i++){
			String sentence = input.nextLine();
			if (sentence.length() > 10) {
				if (sentence.substring(0, 10).equals(compare)){
					System.out.println(sentence.substring(11));
				}
			}
		}

		input.close();
	}
}
