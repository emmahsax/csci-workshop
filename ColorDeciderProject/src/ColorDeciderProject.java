import java.util.ArrayList;
import java.util.Random;

public class ColorDeciderProject {
	private static ArrayList<Dancer> dancers = new ArrayList<Dancer>();

	public static void main(String[] args) {
		dancers.add(new Dancer("Emma")); // 0
		dancers.add(new Dancer("Halie")); // 1
		dancers.add(new Dancer("Hannah")); // 2
		dancers.add(new Dancer("EJ")); // 3
		dancers.add(new Dancer("Erin")); // 4
		dancers.add(new Dancer("Rachel")); // 5
		dancers.add(new Dancer("Mariah")); // 6
		dancers.add(new Dancer("Katie")); // 7
		dancers.add(new Dancer("Allie")); // 8
		dancers.add(new Dancer("Alaina")); // 9

		dancers.get(0).addCannotMatch(dancers.get(6));
		dancers.get(0).addCannotMatch(dancers.get(7));
		dancers.get(1).addCannotMatch(dancers.get(3));
		dancers.get(1).addCannotMatch(dancers.get(2));
		dancers.get(2).addCannotMatch(dancers.get(3));
		dancers.get(2).addCannotMatch(dancers.get(1));
		dancers.get(3).addCannotMatch(dancers.get(1));
		dancers.get(3).addCannotMatch(dancers.get(2));
		dancers.get(3).addCannotMatch(dancers.get(4));
		dancers.get(4).addCannotMatch(dancers.get(5));
		dancers.get(4).addCannotMatch(dancers.get(3));
		dancers.get(5).addCannotMatch(dancers.get(4));
		dancers.get(5).addCannotMatch(dancers.get(8));
		dancers.get(6).addCannotMatch(dancers.get(0));
		dancers.get(6).addCannotMatch(dancers.get(9));
		dancers.get(7).addCannotMatch(dancers.get(0));
		dancers.get(7).addCannotMatch(dancers.get(8));
		dancers.get(7).addCannotMatch(dancers.get(9));
		dancers.get(8).addCannotMatch(dancers.get(7));
		dancers.get(8).addCannotMatch(dancers.get(9));
		dancers.get(8).addCannotMatch(dancers.get(5));
		dancers.get(9).addCannotMatch(dancers.get(8));
		dancers.get(9).addCannotMatch(dancers.get(7));
		dancers.get(9).addCannotMatch(dancers.get(6));

		for (Dancer dancer : dancers) {
			assignColor(dancer);
		}

		checkColors();

		printAll();
	}

	public static void printAll() {
		for (Dancer dancer : dancers) {
			System.out.println(dancer.getName() + " will wear " + dancer.getColor());
		}
	}

	public static void printNoMatches() {
		for (Dancer dancer : dancers) {
			for (Dancer noMatch : dancer.getCannotMatches()) {
				System.out.println(dancer.getName() + " cannot match colors with " + noMatch.getName());
			}
		}
	}

	public static void assignColor(Dancer dancer) {
		Random rand = new Random();
		int x = rand.nextInt(8);

		switch (x) {
			case 0: dancer.setColor(Color.DARK_BLUE);
			break;
			case 1: dancer.setColor(Color.GREEN);
			break;
			case 2: dancer.setColor(Color.LIGHT_BLUE);
			break;
			case 3: dancer.setColor(Color.ORANGE);
			break;
			case 4: dancer.setColor(Color.PINK);
			break;
			case 5: dancer.setColor(Color.PURPLE);
			break;
			case 6: dancer.setColor(Color.YELLOW);
			break;
			case 7: dancer.setColor(Color.RED);
			break;
		}
	}

	public static void checkColors() {
		for (Dancer dancer : dancers) {
			for (Dancer noMatch : dancer.getCannotMatches()) {
				checkTwo(dancer, noMatch);
			}
		}
	}

	// Could potentially get into an infinite loop (even though I think the chances
	// are small...
	public static void checkTwo(Dancer d1, Dancer d2) {
		if (d1.getColor() == d2.getColor()) {
			assignColor(d2);
			for (Dancer noMatch : d2.getCannotMatches()) {
				checkTwo(d2, noMatch);
			}
		}
	}
}
