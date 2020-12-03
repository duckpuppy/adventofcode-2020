// I hate rust. Maybe I'll come back to this

use std::{
    fs::File,
    io::{self, prelude::*, BufReader},
};

fn main() -> io::Result<()> {
    let file = File::open("input")?;
    let reader = BufReader::new(file);

    let lines: Vec<String> = reader.lines()
        .map(|l| l.expect("Could not parse line"))
        .collect();
    lines = lines.sort_by(|a, b| b.cmp(a));
    println!("{:?}", lines);

    Ok(())
}
