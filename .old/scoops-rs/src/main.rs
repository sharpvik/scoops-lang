extern crate regex;

use std::env;
use std::ffi::OsStr;
use std::path::Path;

mod util;


#[derive(PartialEq)]
enum Emode {
    Kwd,    // command line keyword
    Exe,    // end with execution
    Asm,    // end with assembling
    Com,    // just compile
    Err     // error!
}


fn main() {
    /* Collect command line arguments. */
    let args: Vec<String> = env::args().collect();

    /* Check if there are enough command line arguments to begin with. */
    if args.len() < 2 {
        util::hint();
        util::err("Not enough command line arguments.");
    }

    /* 
     * Regular expressions for
     *     1. input file  (ifile | source code | root)
     *     2. output fiel (ofile | target file )
     *     3. command line keyword ("help" | "hint")
     */
    let ifile_exp = regex::Regex::new(r"^[\w\\/]+\.scp(a|b)?$").unwrap();
    let ofile_exp = regex::Regex::new(r"^_[\w\\/]+\.scp(a|b)?$").unwrap();
    let command_exp = regex::Regex::new(r"^[a-z]+$").unwrap();

    /* 
     * First command line argument must represent a filename of code source 
     * (e.g. "hello.scp") or a command line keyword (e.g. "help").
     * Since it is so important, I called it "root".
     * 
     * 'target' is an optional filename for compilation output. In call to
     * 'scoops' one doesn't need to specify any compilation flags like in C --
     * instead, if you want your source code to be compiled or assembled rather
     * than executed straight away, you specify the destination filename
     * prefixed with '_'.
     */
    let root: String = String::from(&args[1]);
    let target: Option<String>;
    if args.len() > 2 && ofile_exp.is_match(&args[2]) {
        target = Some( String::from(&args[2]) );
    } else {
        target = None;
    }

    /* 
     * To account for modes of use, we must first of all determine the 'emode'.
     * 'emode' holds a value of enum type 'Emode' and is determined from the
     * first command line argument -- that is either a source file name or a
     * command line keyword like 'help' -- and second command line argument
     * (if present), that may serve as compilation output target.
     */
    let emode: Emode = get_emode(&root, &target, &ifile_exp, &command_exp);

    /* Pick execution branch based on 'emode' and 'root'. */
    match emode {
        Emode::Kwd => {
            let cmd: &str = &args[1];
            match cmd {
                "hint" => util::hint(),
                "help" => util::help(),
                     _ => util::hint()
            }
        },

        Emode::Err => util::err(
            "The 'run_scoops' function received erroneous 'emode'."
        ),

        /* 
         * At this point it is evident that no error occured while parsing
         * command line arguments and user didn't pass a command line keyword.
         * We are sure that user is trying to run some code here, so we proceed
         * into the 'run_scoops' function.
         */
        _ => run_scoops(root, target, emode, args)
    }
}


fn get_emode(first_arg: &str, second_arg: &Option<String>,
             ifile_exp: &regex::Regex, command_exp: &regex::Regex) -> Emode {

    if ifile_exp.is_match(first_arg) {

        let mut emode = Emode::Exe;
        if let Some(target) = second_arg {

            if !Path::new(&target).exists() {
                return emode;
            }

            /* If path specified in 'target' actually exists, let's ask user
               whether he/she wants to override that file or not. */
            let choice = util::ask(
                &format!(
                    "I figured that the file '{}' already exists. Do you want me to compile into it (this will override the existing file)?\n\nYour answer ('y' = 'yes', anything else = 'no'):",
                    &target
                )
            );

            util::warn(
                &format!(
                    "If you meant to pass '{0}' as a command line argument to your script, please consider one of the following next time:\n\t1. Prefix with `./` => './{0}'\n\t2. Remove _underscore from '{0}'",
                    &target
                )
            );

            if &choice == "y" {
                emode = get_compile_type(first_arg, target);
            } else {
                std::process::exit(0);
            }

        }    
        emode

    } else if command_exp.is_match(first_arg) {

        return Emode::Kwd;

    } else {

        util::log("Cannot determine 'emode'.");
        util::log(
"First command line argument is neither a file nor a command line keyword."
        );
        Emode::Kwd // display usage hint despite error

    }

}


fn get_extension(filename: &str) -> &str {
    match Path::new(filename).extension().and_then(OsStr::to_str) {
        Some(ext) => ext,
        None => "",
    }
}


fn get_compile_type(src: &str, trg: &str) -> Emode {
    let src_ext = get_extension(src);
    let trg_ext = get_extension(trg);

    let src_lvl = get_compile_level(src_ext);
    let trg_lvl = get_compile_level(trg_ext);

    if src_lvl - trg_lvl <= 0 {
        util::log(
            "Invalid compilation hierarchy. Cannot compile upside down."
        );
        return Emode::Err;
    }

    match src_lvl + trg_lvl {
        3     => Emode::Com,
        2 | 1 => Emode::Asm,
        _     => {
            util::log("The 'get_compile_type' function failed miserably.");
            Emode::Err
        }
    }
}


fn get_compile_level(extension: &str) -> i32 {
    match extension {
        "scp"  => 2,
        "scpa" => 1,
        "scpb" => 0,
        _      => -1
    }
}


fn run_scoops(root: String, target: Option<String>, emode: Emode,
              args: Vec<String>) {

    util::log("Running Scoops...");

    let root_compile_lvl = get_compile_level( get_extension(&root) );

    /* Defining all the necessary things Scoops will have to do based on 'root'
       compile level and 'emode'. */
    let must_compile = root_compile_lvl == 2;
    let must_assemble = root_compile_lvl > 0 && emode != Emode::Com;
    let must_execute = emode == Emode::Exe;

    if must_compile {
        util::log("Compiling...");
    }

    if must_assemble {
        util::log("Assembling...");
    }

    if must_execute {
        util::log("Executing...");
    }

}
