/**
 * If we execute the following code in Javascript, what will the browser's console show?
 * Gotchas: it's not 'outside', or 'inside'
 */

var text = 'outside';
function logIt() {
	console.log(text);
	var text = 'inside';
};
logIt();

/**
 * The console will log 'undefined'. This is because the declaration of `text` is hoisted to the top of
 * of the `logIt()` scope, but the assignment of `text` to 'outside' is not hoisted. Therefore, when
 * the code evaluates `console.log(text)`, it does not know what `text` is yet.
 */
