let num = 20;

for(let x = 1; x <= 20; x++) {
  if(x % 3 === 0 && x % 5 === 0)
    console.log("FIZZBUZZ")
  else if(x % 3 === 0)
    console.log("FIZZ")
  else if(x % 5 === 0)
    console.log("BUZZ")
  else
    console.log("-")
}

// Love JS 
// `> node index` to run the code