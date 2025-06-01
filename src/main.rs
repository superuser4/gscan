use clap::Parser;
use std::borrow::Cow;
mod gscan;

#[derive(Parser,Debug)]
#[command(name = "GScan",version = "0.1.0", about="Concurrent Port Scanner", long_about = None)]
struct Args {

    #[arg(short, long)]
    ip: String,

    #[arg(short,long)]
    port: Option<u16>,
}

fn main() { 
    let args = Args::parse();
    let ip_addr: String = args.ip; 
    let ports: Cow<[u16]> = match args.port {
        Some(port) => {
            let range_vec: Vec<u16> = (1..port).collect();
            Cow::Owned(range_vec)
        }
        None => Cow::Borrowed(gscan::scanner::MOST_COMMON_TCP), 
    };

    
    println!("GScan Started at {}", chrono::Local::now().format("%Y-%m-%d %H:%M:%S"));
    println!("Started Scanning IP: {}", ip_addr);
    println!("Port\tStatus\tService"); 
    println!("----\t------\t-------");
    
    gscan::scanner::scan_ports(&ip_addr, ports); 
}
