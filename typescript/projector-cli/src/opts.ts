import { cwd } from "node:process";
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
      defaultOption: true,
      multiple: true,
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
      defaultValue: cwd(),
    },
  ]) as Opts;
}
