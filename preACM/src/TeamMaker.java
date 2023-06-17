import java.util.ArrayList;
import java.util.Collections;
import java.util.Random;
import java.util.Scanner;

public class TeamMaker {
  private static int numPlayers;
  private static int numTeams;
  private static int counter = 0;
  private static String[] allPlayers;
  private static ArrayList<String> players = new ArrayList<String>();
  private static ArrayList<ArrayList<String>> teams = new ArrayList<ArrayList<String>>();

  public static void main(String[] args) {
    Scanner input = new Scanner(System.in);

    System.out.println("Please type in how many players are allowed:");
    numPlayers = Integer.parseInt(input.nextLine());

    System.out.println("Please type in how many teams are allowed:");
    numTeams = Integer.parseInt(input.nextLine());

    System.out.println("Please type in the names of who wishes to play, separated by a semicolon and a space ('; '):");
    allPlayers = input.nextLine().split("; ");

    if (allPlayers.length < numPlayers) { // Amount of people who want to play is less than number of people who can play
      System.out.println("I don't have enough people who want to play.");
      System.exit(0);
    } else if (numPlayers < numTeams) { // Amount of people who can play is less than the number of teams
      System.out.println("The number of players is less than the number of teams.");
      System.exit(0);
    }

    for (counter = 0; counter < numTeams; counter++) {
       teams.add(new ArrayList<String>()); // Make an array list for each team
    }

    shuffleAllPlayers();

    for (counter = 0; counter < numPlayers; counter++) {
      players.add(allPlayers[counter]); // Add only numPlayers amount of people to players array list
    }

    Collections.shuffle(players);

    for (counter = 0; counter < numPlayers; counter++) { // Distribute players to teams
      int result = counter % numTeams;
      teams.get(result).add(players.get(counter));
    }

    for (counter = 0; counter < numTeams; counter++) {
      System.out.println("Team " + (counter+1) + ": " + teams.get(counter)); // Print out each team
    }

    input.close();
  }

  /**
   * Shuffles the strings in allPlayers so that the array is randomized.
   */
  private static void shuffleAllPlayers(){
    Random rand = new Random();

    for (int i = 0; i < allPlayers.length; i++) {
        int randomIndex = rand.nextInt(allPlayers.length);
        String temp = allPlayers[i];

        allPlayers[i] = allPlayers[randomIndex];
        allPlayers[randomIndex] = temp;
    }
  }
}
