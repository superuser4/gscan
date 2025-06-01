use std::net::TcpStream;

pub fn scan_port(ip_addr: String, port: u16) -> bool {
    let ip_and_port: String = ip_addr + ":" + &port.to_string();
    if let Ok(_stream) = TcpStream::connect(ip_and_port) {
        true
    } else { 
        false
    }
}
