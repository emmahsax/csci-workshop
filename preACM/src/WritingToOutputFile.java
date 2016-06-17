import java.io.File;
import java.io.FileNotFoundException;
import java.io.PrintWriter;
import java.util.Scanner;


public class WritingToOutputFile {
        public static void main(String[] args) {
                String inputFileName = args[0];
                String outputFileName = args[1];
                
                String inputText = "";
                Scanner in;
                try{
                        in = new Scanner(new File(inputFileName));
                        while(in.hasNextLine()) {
                                inputText = in.nextLine();
                        }
                        in.close();
                } catch(FileNotFoundException e){
                        e.printStackTrace();
                }
                
                try{ // below makes output file if it doesn't exist.
                        PrintWriter out = new PrintWriter(outputFileName);
                        for (int i = 0; i < 5; i++){
                                out.println(inputText);
                        }
                        out.close();
                } catch(FileNotFoundException e) {
                        e.printStackTrace();
                }
                
        }

}

