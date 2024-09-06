import path from "node:path";
import cli from "command-line-args";

export type Opts = {
  args?: string[];
  // specified working directory
  pwd?: string;
  // json config that contains all values
  config?: string;
};

export default function getOps(): Opts {
  return cli([
    {
      name: "args",
      defaultValue: true,
      type: String,
    },
    {
      name: "config",
      alias: "c",
      type: String,
    },
    {
      name: "pwd",
      alias: "p",
      type: String,
      defaultOption: path.cwd(),
    },
  ]) as Opts;
}
