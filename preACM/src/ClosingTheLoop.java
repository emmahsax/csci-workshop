import java.util.ArrayList;
import java.util.Collections;

public class ClosingTheLoop {

	public static void main(String[] args) {
		Kattio input = new Kattio(System.in);
		
		int caseAmount = input.getInt();

		for (int k = 0; k < caseAmount; k++) {
			ArrayList<Integer> blue = new ArrayList<Integer>();
			ArrayList<Integer> red = new ArrayList<Integer>();
			int stringCount = input.getInt();
			
			for(int i = 0; i < stringCount; i++) {
				String token = input.getWord();
				int n = new Integer(token.substring(0, (token.length() - 1)));
				if (token.substring(token.length()-1).equals("R")) {
					red.add(n - 1);
				} else {
					blue.add(n - 1);
				}
			}
			
			Collections.sort(blue);
			Collections.reverse(blue);
			Collections.sort(red);
			Collections.reverse(red);
			
			int sum = 0;
			int smallerSize = Math.min(blue.size(), red.size());
			for (int i = 0; i < smallerSize; i++) {
				sum += blue.get(i) + red.get(i);
			}
			
			System.out.println("Case #" + (k+1) + ": " + sum);
		}
		
		input.close();
	}
}
