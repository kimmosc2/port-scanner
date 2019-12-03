package table

// port comparison
var portTable = make(map[int]string)

func init() {
    portTable[21] = "FTP"
    portTable[22] = "SSH"
    portTable[23] = "Telnet"
    portTable[25] = "SMTP"
    portTable[69] = "TFTP"
    portTable[80] = "HTTP"
    portTable[110] = "POP3"
    portTable[443] = "HTTPS"

    portTable[1433] = "SQLServer"
    portTable[1521] = "Oracle"

    portTable[3306] = "mysql"
    portTable[3389] = "RemoteDesktop"

    portTable[5000] = "DB2"
    portTable[5672] = "RabbitMQ"

    portTable[6379] = "Redis"

    portTable[8080] = "tomcat"
    portTable[9092] = "pointbase"
    portTable[9200] = "ElasticSearch"
    portTable[9300] = "ElasticSearch"

    portTable[15672] = "RabbitMQ UI"
    portTable[27017] = "MongoDB"
}

func GetPossibility(port int) (result string) {
    if result, exist := portTable[port]; exist {
        return result
    }
    return ""
}
