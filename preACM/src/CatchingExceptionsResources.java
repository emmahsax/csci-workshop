// File Name : ExcepTest.java
import java.io.*;
public class ExcepTest{
    public static void main(String args[]){
        try{
            int a[] = new int[2];
            System.out.println("Access element three :" + a[3]);
        }catch(ArrayIndexOutOfBoundsException e){
            System.out.println("Exception thrown  :" + e);
        }
        System.out.println("Out of the block");
    }
}

// PRODUCES:
// Exception thrown  :java.lang.ArrayIndexOutOfBoundsException: 3
// Out of the block

  static class FirstException extends Exception { }
  static class SecondException extends Exception { }

  public void rethrowException(String exceptionName) throws Exception {
        try {
            if (exceptionName.equals("First")) {
                throw new FirstException();
            } else {
                throw new SecondException();
            }
        } catch (Exception e) {
            throw e;
        }
  }

// catching multiple exceptions with one catch block

try{
    System.out.println("printing something");
} catch (IOException|SQLException ex) {
    logger.log(ex);
    throw ex;
}
