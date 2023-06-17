import java.util.Scanner;

public class CuttingBrownies {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        int caseAmount = input.nextInt();
        input.nextLine();

        for (int i = 0; i < caseAmount; i++){
        int n1 = input.nextInt();
        int n2 = input.nextInt();
        String person = input.next();

        if (person.equals("Harry")){
            if ((n2 - 1) > n1) {
                System.out.println("Harry can win");
            } else if (n2 == 2 && n1 == 1) {
                System.out.println("Harry can win");
            } else {
                System.out.println("Harry cannot win");
            }
        } else if (person.equals("Vicky")) {
            if ((n1 - 1) > n2) {
                System.out.println("Vicky can win");
            } else if (n1 == 2 && n2 == 1) {
                System.out.println("Vicky can win");
            } else {
                System.out.println("Vicky cannot win");
            }
        }
        }

        input.close();
    }
}
