use std::env;
use std::fs;
use std::path::Path;
use std::path::PathBuf;

pub fn read_input<'life>(year: &'life str, day: &'life str) -> String {
    let project_dir = get_project_dir();
    let input_file_path = get_input_file_path(&project_dir, &year, &day);

    fs::read_to_string(input_file_path)
        .unwrap()
        .trim()
        .to_string()
}

fn get_input_file_path(project_dir: &str, year: &str, day: &str) -> PathBuf {
    Path::new(project_dir)
        .join("src")
        .join(["year_", year].join(""))
        .join(["day", day].join(""))
        .join("input")
}

fn get_project_dir() -> String {
    let binding = env::current_dir().unwrap();
    String::from(binding.as_path().to_str().unwrap())
}
