import { expect, test } from "vitest";

const add = (a, b) => a + b;

test("adds", () => {
  const x = add(1, 2);

  expect(x).toBe(3);
});
