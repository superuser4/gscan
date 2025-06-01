use std::net::TcpStream;

pub fn scan_port(ip_addr: String, port: u16) -> bool {
    let ip_and_port: String = ip_addr + ":" + &port.to_string();
    TcpStream::connect(ip_and_port).is_ok()
}
