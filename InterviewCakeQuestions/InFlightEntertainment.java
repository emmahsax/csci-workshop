/**
 * There is an in-flight entertainment system with on-demand movie streaming. Users on long flights like to watch
 * multiple movies, but it's been found that their flight often will land before they can finish the second movie.
 * So write a method that will take the total length of the flight and an array of movie lengths available,
 * and return a simple boolean that states whether there exist two movies (which are not the same movie) where
 * the lengths equal the flight length. Assume your users will watch exactly two movies. Optimize for runtime over
 * memory.
 * Gotchas: this can be done in O(n) time and O(n) space
 */

public class InFlightEntertainment {
    public static boolean moviesMatchFlightLength(int flightLength, int[] movieLengths) {
        Set<Integer> movieLengthsSeen = new HashSet<Integer>();

        for (int firstMovieLength : movieLengths) {
            int matchingSecondMovieLength = flightLength - firstMovieLength;

            if (movieLengthsSeen.contains(matchingSecondMovieLength)) {
                return true;
            }

            movieLengthsSeen.add(firstMovieLength);
        }

        return false;
    }
}
