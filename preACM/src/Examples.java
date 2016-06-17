import java.util.Arrays;
import java.util.Scanner;

public class Examples {

	public static void main(String[] args) {
///////////////////////Scanner In/////////////////////////////////
		Scanner sc = new Scanner(System.in);
		
		//System.out.println(sc.nextInt());
		//System.out.println(sc.nextLine());
		
		//String[] scVal = sc.nextLine().split(" ");
		//System.out.println(Arrays.toString(scVal));
		
		
///////////////////////Converting Bases/////////////////////////////////

		//Convert to base 2
		String a = Integer.toString(6, 2); //returns "110"
		//Convert from base 2
		int b = Integer.parseInt("110",2); //returns 6
		
		//Convert to base 3
		String c = Integer.toString(6, 3); //returns "20"
		//Convert from base 3
		int d = Integer.parseInt("20",3);  //returns 6
		
		int e = Integer.parseInt(Integer.toString(25,5));   //returns 100
		int f = Integer.parseInt(Integer.toString(100), 5); //returns 25
		
		//BigInteger
		BigInteger aa = new BigInteger("1993");//returns 1993
		BigInteger bb = new BigInteger("11", 2); //returns 3
		String cc = b.toString(2); //returns "11"
		BigInteger dd = a.add(BigInteger.ONE); //returns 1994
				
/*
		Examples from http://stackoverflow.com/questions/13654671/
		what-is-the-method-in-the-api-for-converting-between-bases:
		parseInt("0", 10)           // returns 0
		parseInt("473", 10)         // returns 473
		parseInt("+42", 10)         // returns 42
		parseInt("-0", 10)          // returns 0
		parseInt("-FF", 16)         // returns -255
		parseInt("1100110", 2)      // returns 102
		parseInt("2147483647", 10)  // returns 2147483647
		parseInt("-2147483648", 10) // returns -2147483648
		parseInt("2147483648", 10)  // throws NumberFormatException
	(You can't write a number in base 8 that has 9s in it...):
		parseInt("99", 8)           // throws NumberFormatException 
		parseInt("Kona", 10)        // throws NumberFormatException
		parseInt("Kona", 27)        // returns 411787
*/
		
///////////////////////String to char/////////////////////////////////
int g = ((int)'a'); //97
int h = ((int)'z'); //122
char i = ((char) 97); // 'a'
String j = ((char) 122) + ""; // "z"
String k = String.valueOf((char) 122); // "z"

///////////////////////regular expressions/////////////////////////////////

//See Class Pattern in the java API

/* Order of operations
	1     	Literal escape    \x
	2     	Grouping 	     [...]
	3     	Range 	         a-z
	4     	Union 	         [a-e][i-u]
	5     	Intersection 	 [a-z&&[aeiou]]
*/
System.out.println("test string".contains("z")); //false, since no was "Z" found

System.out.println("aaeiou".matches("[aeiou]*")); // true, is 0 or more vowel

System.out.println("d".matches("[a-f&&[^a-c]&&[^e]]")); // true, matches only d and f

String test = "https://harre096.gitub.io";
System.out.println(test.matches("^(http|https|ftp)://.*$")); //true

System.out.println("stop first".replaceAll("st", "pf")); //pfop firpf


	}

}
