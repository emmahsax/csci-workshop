/**
 * Given an array of integers 1..n, where the length of the array is n+1, find where the
 * 1 (at least) duplicate is in the array (if there are multiple duplicates, find the first
 * one). Try to complete this in O(n) time and O(1) space by treating the array like a linked list
 * where the value of the node is the integer from the array and the "next" pointer points to the
 * "value-eth" node in the list:
 * [2, 3, 1, 3]
 * 2 points to the first 3, first 3 points to the 1, 1 points to the 2, and last 3 points to the 1
 */

// We'll solve this by realizing that by treating the array like a linked list,
// there must be at least one loop somewhere in the linked list. So now we can
// traverse the array once to find the start of the loop. If we can find the start
// of the loop, then that position where position = index + 1 is the first repeat.

public static int findADuplicate (int[] intArray) {
  // Get somewhere inside the loop
  int n = intArray.length - 1;
  int position = n + 1;

  for (int x = 0; x < n; n++) {
    position = intArray[position - 1]; // last number in array
  }

  // find the entire length of the loop
  int remember = position;
  int moving = intArray[position - 1]; // a step ahead
  int count = 1;

  while (moving != remember) {
    moving = intArray[moving - 1]; // keep stepping ahead
    count++;
  }

  // find the first node in the loop
  int start = n + 1;
  int ahead = n + 1;

  for (int x = 0; x < count; x++) {
    ahead = intArray[ahead - 1];
  }

  while (start != ahead) {
    start = intArray[start - 1];
    ahead = intArray[ahead - 1];
  }

  return start;
}
