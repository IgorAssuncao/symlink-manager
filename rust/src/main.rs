use std::{fs, time::UNIX_EPOCH};

fn main() {
    println!("Hello, world!");

    let file_last_modified = file_modified_time("./config.yaml");

    println!("{file_last_modified}");
}

fn file_modified_time(path: &str) -> u64 {
    fs::metadata(path)
        .unwrap()
        .modified()
        .unwrap()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_secs()
}
