import java.util.Scanner;

public class CircuitCounting {

	public static void main(String[] args) {
		Scanner input = new Scanner(System.in);
		
		int caseAmount = input.nextInt();
		int[] x = new int[caseAmount];
		int[] y = new int[caseAmount];
		
		for (int i = 0; i < caseAmount; i++){
			x[i] = input.nextInt();
			y[i] = input.nextInt();
		}
		
		int[] values = {0, 1, 2, 3, 4};
		
		breakOut();
		
		input.close();
	}
	
	private static void breakOut(int startIndex, int endIndex, int startIndex2, int endIndex2) {
		
	}

}
