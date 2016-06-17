import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Scanner;

public class SafePassage {
	
	private static ArrayList<Integer> notSafe = new ArrayList<Integer>();
	private static ArrayList<Integer> safe = new ArrayList<Integer>();
	private static int total = 0;
	private static int secondFastest = 0;
	private static int fastest = 0;
	private static boolean cloakAtDorm = false;

	public static void main(String[] args) {
		Scanner input = new Scanner(System.in);
		
		int caseAmount = input.nextInt();
		
		int[] arr = new int[caseAmount];
		
		for (int i = 0 ; i < caseAmount; i++){
			arr[i] = input.nextInt();
		}
		
		Arrays.sort(arr);
		ArrayList<Integer> sortedArr = new ArrayList<Integer>();
		
		for (int i = 0; i < arr.length; i++) {
			sortedArr.add(arr[i]);
			notSafe.add(arr[i]);
		}
		secondFastest = sortedArr.get(1);
		fastest = sortedArr.get(0);
		
		
		while (safe.size() != caseAmount) {
//			System.out.println("Not safe: " + notSafe);
//			System.out.println("Safe: " + safe);
			if (notSafe.size() == 3 && cloakAtDorm == false && notSafe.contains(secondFastest) && notSafe.contains(fastest)) {
				caseThreeLeft();
				Collections.sort(safe);
				Collections.sort(notSafe);
			}else if (cloakAtDorm == false && notSafe.contains(secondFastest)) {
				moveFastest();
				Collections.sort(safe);
				Collections.sort(notSafe);
			} else if (cloakAtDorm == false && safe.contains(secondFastest)) {
				moveSlowest();
				Collections.sort(safe);
				Collections.sort(notSafe);
			} else if (cloakAtDorm == true) {
				sendFastestBack();
				Collections.sort(safe);
				Collections.sort(notSafe);
			}
//			System.out.println(total);
//			System.out.println("Cloak: " + cloakAtDorm);
		}
		
		System.out.println(total);

		input.close();
	}

	//call when  2nd fastest is NOT SAFE and cloak is false
	private static void moveFastest(){
		safe.add(notSafe.remove(1));
		safe.add(notSafe.remove(0));
		total += secondFastest;
		cloakAtDorm = true;
	}
	
	//call when 2nd fastest is safe and cloak is false
	private static void moveSlowest() {
		int slowGuySpeed = notSafe.remove(notSafe.size()-1);
		safe.add(slowGuySpeed);
		safe.add(notSafe.remove(notSafe.size()-1));
		total += slowGuySpeed;
		cloakAtDorm = true;
	}
	
	//call when cloak is true
	private static void sendFastestBack() {
		int fastest = safe.remove(0);
		notSafe.add(fastest);
		total += fastest;
		cloakAtDorm = false;
	}
	
	//case three people left, send fastest and slowest
	private static void caseThreeLeft(){
		int slowGuySpeed = notSafe.remove(notSafe.size()-1);
		safe.add(slowGuySpeed);
		safe.add(notSafe.remove(0));
		total += slowGuySpeed;
		cloakAtDorm = true;
	}
}
