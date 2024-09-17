use anyhow::{anyhow, Context, Ok, Result};
use std::path::PathBuf;

use crate::opts::Opts;

#[derive(Debug)]
pub struct Config {
    pub operation: Operation,
    pub pwd: PathBuf,
    pub config: PathBuf,
}
#[derive(Debug)]
pub enum Operation {
    Print(Option<String>),
    Add(String, String),
    Remove(String),
}

impl TryFrom<Vec<String>> for Operation {
    type Error = anyhow::Error;

    fn try_from(value: Vec<String>) -> Result<Self, Self::Error> {
        let mut value = value;
        if value.len() == 0 {
            return Ok(Operation::Print(None));
        }

        let term = value.get(0).expect("expect it exist");

        if term == "add" {
            if value.len() != 3 {
                let err = anyhow!("operation add expects 2 arguments but got {}", value.len());
                return Err(err);
            }

            let mut drain = value.drain(1..=2);

            return Ok(Operation::Add(
                drain.next().expect("to exist"),
                drain.next().expect("to exist"),
            ));
        }

        if term == "rm" {
            if value.len() != 1 {
                let err = anyhow!("operation rm expects 1 argument but got {}", value.len());
                return Err(err);
            }

            let arg = value.pop().expect("to exist");
            return Ok(Operation::Remove(arg));
        }

        if value.len() > 1 {
            let err = anyhow!(
                "operation print expects 0 or 1 argument but got {}",
                value.len()
            );
            return Err(err);
        }

        let arg = value.pop().expect("to exist");
        return Ok(Operation::Print(Some(arg)));
    }
}

impl TryFrom<Opts> for Config {
    type Error = anyhow::Error;

    fn try_from(value: Opts) -> Result<Self> {
        return Ok(Config {
            operation: value.args.try_into()?,
            config: get_config(value.config)?,
            pwd: get_pwd(value.pwd)?,
        });
    }
}

fn get_config(config: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(v) = config {
        return Ok(v);
    }

    let loc = std::env::var("XDG_CONFIG_HOME").context("unable to get XDG_CONFIG_HOME")?;
    let mut loc = PathBuf::from(loc);

    loc.push("projector");
    loc.push("projector.json");

    return Ok(loc);
}

fn get_pwd(pwd: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(v) = pwd {
        return Ok(v);
    }

    return Ok(std::env::current_dir().context("error getting current directory")?);
}
