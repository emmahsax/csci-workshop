
public class ADifferentProblem {
	public static void main(String[] args) {
		Kattio nums = new Kattio(System.in);
		
		while (nums.hasMoreTokens()) {
			long num1 = nums.getLong();
			long num2 = nums.getLong();
			System.out.println(Math.abs(num1 - num2));
		}
		
		nums.close();
	}

}
