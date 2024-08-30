enum TS {
  FOO,
  BAR,
  BAZ,
}

const a: number[] = [];
const b = a;

b.push(1);

console.log(b);
