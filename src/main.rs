use std::env;
use year_2015::day01::lib as lib_2015_01;

pub mod utils;
pub mod year_2015;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 4 {
        println!(
            "Insufficient arguments! Please, pass year, day and part e.g. cargo run 2015 01 1"
        );
        return;
    }
    let (year, day, part) = (args[1].clone(), args[2].clone(), args[3].clone());
    match (year.as_str(), day.as_str(), part.as_str()) {
        ("2015", "01", "1") => lib_2015_01::part_one(),
        ("2015", "01", "2") => lib_2015_01::part_two(),
        _ => println!("Invalid input"),
    }
}
