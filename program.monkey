let five = 5;
let ten = 10;

let add = fn(x, y) {
    x + y;
};

let result = add(five, ten);


let map = fn(arr, f) { let iter = fn(arr, acc) { if (len(arr) == 0) { acc; } else { iter(rest(arr), push(acc, f(first(arr)))); } } iter(arr, []); };
let reduce = fn(arr, init, f) { let iter = fn(arr, res) { if (len(arr) == 0) { res; } else { iter(rest(arr), f(res, first(arr))); } } iter(arr, init); };
let sum = fn(arr) { reduce(arr, 0, fn(init, el) { init + el }); };

sum([1, 2, 3, 4, 5]);
