import java.util.Scanner;

public class Message {

  public static void main(String[] args) {
    Scanner input = new Scanner(System.in);

    int caseAmount = input.nextInt();
    input.nextLine();

    for (int i = 0; i < caseAmount; i++){
      String message = input.nextLine();
      int SQ = 0;
      int counter = 0;

      while(SQ == 0) {
        if ((counter * counter) >= message.length()) {
          SQ = counter;
        }
        counter++;
      }

      for (int j = SQ; j > 0; j--) {
        for (int k = 0; k < SQ; k++) {
          int charIndex = SQ * (SQ - k) - j;
          if (charIndex < message.length()) {
            System.out.print(message.charAt(charIndex));
          }
        }
      }
      int charIndex = (SQ * SQ)- 0;
      if (charIndex < message.length()) {
        System.out.print(message.charAt(charIndex));
      }
      System.out.println();
    }

    input.close();
  }
}
