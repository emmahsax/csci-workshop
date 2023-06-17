import java.util.ArrayList;

public class Dancer {
  private Color color;
  private String name;
  private char size;
  private ArrayList<Dancer> cannotMatch = new ArrayList<Dancer>();

  public Dancer(String name) {
    this.name = name;
  }

  public String getName() {
    return name;
  }

  public Color getColor() {
    return color;
  }

  public char getSize() {
    return size;
  }

  public ArrayList<Dancer> getCannotMatches() {
    return cannotMatch;
  }

  public void setColor(Color color) {
    this.color = color;
  }

  public void setSize(char size) {
    this.size = size;
  }

  public void addCannotMatch(Dancer dancer) {
    cannotMatch.add(dancer);
  }
}
