use crate::utils::{input_reader, writer};

const YEAR: &str = "2015";
const DAY: &str = "01";

pub fn part_one() {
    let input = input_reader::read_input(YEAR, DAY);
    let mut floor = 0;
    for c in input.chars() {
        if c == '(' {
            floor += 1;
        } else if c == ')' {
            floor -= 1;
        }
    }
    writer::write_asnwer(YEAR, DAY, "1", floor);
}

pub fn part_two() {
    let input = input_reader::read_input(YEAR, DAY);
    let mut floor = 0;
    for (index, character) in input.chars().enumerate() {
        if character == '(' {
            floor += 1;
        } else if character == ')' {
            floor -= 1;
        }
        if floor == -1 {
            // because the index starts at 0
            writer::write_asnwer(YEAR, DAY, "2", index + 1);
            return;
        }
    }
    writer::write_asnwer(YEAR, DAY, "2", "None");
}
