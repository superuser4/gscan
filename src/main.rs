mod portscan;
use clap::Parser;
use portscan::scan_ports;

#[derive(Parser,Debug)]
#[command(version, about, long_about = None)]
struct Args {
    ip_addr: String,
    port_range: u16,
}

fn main() {
    println!("GScan v0.1.0 Started at {}", chrono::Local::now().format("%Y-%m-%d %H:%M:%S"));
    
    let args = Args::parse();
    let ip_addr: String = args.ip_addr; 
    let ports: u16 = args.port_range;

    println!("Started Scanning IP: {}", ip_addr);
    println!("Port\tStatus\tService"); 
    println!("----\t------\t-------");
    
    let arr: Vec<u16> = (0..ports).collect();
    scan_ports(&ip_addr, arr); 
}
