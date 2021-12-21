# Some Assembly Required 1
## Category: Web Exploitation

#### Description:
* Website: http://mercury.picoctf.net:26318/index.html
  * Review `.js` code being used to calculate the flag
  * Use [JSnice.org](http://jsnice.org) to deobfuscate the variables.
  * Use the `navigatePop()` function to decode the integers called in this block:

```JS
let exports;
(async() => {
  const findMiddlePosition = _0x4e0e;
  let leftBranch = await fetch(findMiddlePosition('./JIFxzHyW8W'));
  let rightBranch = await WebAssembly[findMiddlePosition(instantiate)](await leftBranch[findMiddlePosition(arrayBuffer)]());
  let module = rightBranch[findMiddlePosition(instance)];
  exports = module["exports"];
})();
/**
 * @return {undefined}
 */
function onButtonPress() {
  const navigatePop = _0x4e0e;
  let params = document["getElementById"](navigatePop(input))[navigatePop(value)];
  for (let i = 0; i < params["length"]; i++) {
    exports[navigatePop(copy_char)](params[navigatePop(charCodeAt)](i), i);
  }
  exports["copy_char"](0, params["length"]);
  if (exports[navigatePop(check_flag)]() == 1) {
    document[navigatePop(getElementById)](navigatePop("result"))[navigatePop(innerHTML)] = navigatePop("Correct!");
  } else {
    document[navigatePop(getElementById)](navigatePop("result"))[navigatePop(innerHTML)] = navigatePop("Incorrect!");
  }
}
;
```
 * Downloaded the web assembly script found in `await fetch(findMiddlePosition('./JIFxzHyW8W'))`
 * flag found in the source of the `.wasm` code.

__Flag:__ picoCTF{8857462f9e30faae4d037e5e25fee1ce}
