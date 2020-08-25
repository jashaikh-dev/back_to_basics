const rl = require('readline').createInterface({
  input: process.stdin,
  output: process.stdout
});

rl.question("Enter a number: ",  num => {
  if ( num < 0 )
    console.log(`${num} is negative`)
  else if (num > 0)
    console.log(`${num} is positive`)
  else
    console.log(`${num} is zero`)
  rl.close()
});
