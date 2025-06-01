use std::net::TcpStream;

fn scan_port(ip_addr: &String, port: u16) -> bool {
    let ip_and_port: String = ip_addr.to_owned() + ":" + &port.to_string();
    TcpStream::connect(ip_and_port).is_ok()
}

pub fn scan_ports(ip_addr: &String, ports: Vec<u16>) {
    for port in ports {
        if scan_port(ip_addr, port) {
            println!("{}\topen", port);
        }
    }
}
