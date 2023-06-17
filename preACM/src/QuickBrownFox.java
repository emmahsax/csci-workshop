import java.util.ArrayList;
import java.util.Scanner;

public class QuickBrownFox {

  public static void main(String[] args) {
    Scanner input = new Scanner(System.in);

    int caseAmount = input.nextInt();
    input.nextLine();

    for (int i = 0; i<caseAmount; i++){
      String old = input.nextLine();

      ArrayList<Integer> arr = new ArrayList<Integer>();
      String sentence = old.toLowerCase();

      for (int j = 97; j <= 122; j++){
        String c = Character.toString((char)j);
        if (!sentence.contains(c)) {
          arr.add(j);
        }
      }

      if (arr.isEmpty()) {
        System.out.println("pangram");
      } else {
        System.out.print("missing ");
        for (int j=0; j < arr.size(); j++){
          int ch = arr.get(j);
          System.out.print(Character.toString((char) ch));
        }
        System.out.println();
      }
    }
    input.close();
  }
}
