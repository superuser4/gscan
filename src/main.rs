mod scan;

fn main() {
    let ip_addr = String::from("127.0.0.1");
    println!("Started Scanning IP: {}", ip_addr);
    if scan::scan_port(ip_addr, 631) {
        println!("Port: 631 open");
    } else {
        println!("Port closed");
    }
}
