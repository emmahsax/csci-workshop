/**
 * We're building a web game where everybody wins and everyone is friends forever. All a player needs
 * to do is to click a button to win something. And every prize is a good prize. But something
 * is going wrong. What is it?
 * Gotchas: the syntax is just fine... the problem is some unexpected behavior
 */

<button id="btn-0">Button 1!</button>
<button id="btn-1">Button 2!</button>
<button id="btn-2">Button 3!</button>

<script type="text/javascript">
  var prizes = ['A Unicorn!', 'A Hug!', 'Clean Laundry!'];
  for (var btnNum = 0; btnNum < prizes.length; btnNum++) {
    // for each of our buttons, when the user clicks it...
    document.getElementByID('btn-' + btnNum).onClick = function() {
      // tell the player what they've won!
      alert(prizes[btnNum]);
    };
  }
</script>

/**
 * The user's prize is always undefined. The anonymous function that is assigned to the button's onclicks has
 * access to variables in the scope outside of it (called closure). In this case, it has access to btnNum.
 * And when the anonymous function accesses the variable btnNum outside of its scope, it will access whatever
 * the variable is at that time, not a "frozen copy". Therefore, when a player actually clicks on a button,
 * that's when it'll run line 18 and btnNum will be 3 because it has already iterated through each button.
 * Therefore, the way to solve this is to pass in a variable to the anonymous function. Doing this will create a new
 * variable inside of the anonymous function which will never change even as the loop iterates:
 */

<button id="btn-0">Button 1!</button>
<button id="btn-1">Button 2!</button>
<button id="btn-2">Button 3!</button>

<script type="text/javascript">
  var prizes = ['A Unicorn!', 'A Hug!', 'Clean Laundry!'];
  for (var btnNum = 0; btnNum < prizes.length; btnNum++) {
    // for each of our buttons, when the user clicks it...
    document.getElementByID('btn-' + btnNum).onClick = function(frozenBtnNum) {
      // tell the player what they've won!
      alert(prizes[frozenBtnNum]);
    }(btnNum);
  }
</script>
