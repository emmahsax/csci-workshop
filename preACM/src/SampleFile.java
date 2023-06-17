import java.util.Arrays;
import java.util.Scanner;

public class SampleFile {
  public static void main(String[] args) {
    String inputFile = “C://Users/DKC3/Desktop/stuff/testInput.txt”;
    String outputFile = “C://Users/DKC3/Desktop/stuff/testOutput.txt”;
    try {
      BufferedReader in = new BufferedReader(new FileReader(inputFile));
      PrintWriter out = new PrintWriter(new FileWriter(outputFile));

      while(in.ready()) {
        try {
          String text = in.ReadLine();

          // int x = Integer.parseInt(text);
          // double y = Double.parseDouble(text);
          // long z = Long.parseLong(text);
          // String[] inputArray = text.split(" ");
          // do something cool with the text

          out.println(text);
        } catch (Throwable e) {
          // System.out.println(e);
        }
      }

      in.close();
      out.close();
    } catch(Throwable e) {
      // System.out.println(e);
    }
  }
}
