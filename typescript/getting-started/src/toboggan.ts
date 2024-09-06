function getInput() {
  return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`;
}

enum Thing {
  Tree,
  Snow,
}

const things = getInput()
  .split("\n")
  .map((x) => x.split("").map((x) => (x === "." ? Thing.Snow : Thing.Tree)));

const colLength = things[0].length;
let treeCount = 0;

things.forEach((thingRow, i) => {
  if (thingRow[(i * 3) % colLength] === Thing.Tree) {
    treeCount++;
  }
});

console.log(treeCount);
