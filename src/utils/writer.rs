pub fn write_asnwer<T>(year: &str, day: &str, part: &str, answer: T)
where
    T: std::fmt::Debug,
{
    println!("Answer for year {year} day {day} part {part}: {:?}", answer);
}
