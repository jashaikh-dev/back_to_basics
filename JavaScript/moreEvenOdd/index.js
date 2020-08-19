let num = 20


for (let x = 1; x <= num; x++) {
  if (num % 2 === 0)
    console.log(num + 'is even');
  else
    console.log(num + 'is odd');
}

for (let x = 1; x <= num; x++) {
  console.log(  `${x}  ${(x % 2 === 0) ? 'is even' : 'is odd'}` );
}

// Love JS
// `> node index` to run the code