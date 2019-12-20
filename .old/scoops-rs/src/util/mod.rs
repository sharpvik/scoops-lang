/* Box Drawing
│ └ ├ ─ ┬
*/

use std::io::{self, BufRead};


pub fn ask(msg: &str) -> String {
    println!("=====> QUESTION PROMPT\n{}", msg);
    
    let stdin = io::stdin();
    let answer = stdin.lock().lines().next().unwrap().unwrap();

    println!();
 
    String::from( answer.trim() )
}

pub fn warn(msg: &str) {
    println!("=====> WARNING\n{}\n", msg);

}

pub fn log(msg: &str) {
    println!("=====> LOG MESSAGE\n{}\n", msg);
}

pub fn err(msg: &str) {
    eprintln!("=====> ERROR\n{}\n", msg);
    std::process::exit(1);
}

pub fn hint() {
    println!(
"
│ Usage hint:
├──── scoops inputFile.scp
└──── scoops help
"
    );
}

pub fn help() {
    println!(
"
│ Hello, I am Embedded Helper! Welcome!
│
│ Scoops requires at most two command line arguments to know what to do.
│ First one is an input file name (must have).
│ Second one is an output file name (can be ommitted in most cases).
│
│ Let's look at some cases:
├───┬ scoops inputFile.scp [optional command line arguments for your script]
│   └───┬ This is the simplest execution type. Scoops will simply execute
│       └ whatever source code happens to be in the 'inputFile.scp'.
│
├───┬ scoops inputFile.scp _outputFile.scpb [optional command line arguments]
│   └───┬ In this case, Scoops will know to compile, assemble and output code
│       ├ into the outputFile.scpb. You must put '_' in front of the file name,
│       ├ otherwise, Scoops will consider it to be a regular command line 
│       └ argument that your script intends to utilise.
│
├───┬ scoops inputFile.scpb _outputFile.scp [optional command line arguments]
│   └───┬ Provided such command, Scoops will exit with error as it doesn't make
│       └ sense to convert lower level code upwards.
│
│ Valid Scoops file formats:
├──── *.scp  ── (Source)   ── Basic Scoops Source Code file.
├──── *.scpa ── (Assembly) ── Scoops Assembly file.
├──── *.scpb ── (Bytecode) ── Scoops Bytecode file.
│
│ You may also use command line keywords.
│ To use command line keyword, type 'scoops [keyword]'.
├──── help ── Embedded helper.
└──── hint ── Brief usage hint.
"
    );
}
