package Model

type Config struct {
	jmxPath string
	jtlPath string
}

var Conf = Config{"Data/jmx", "Data/jtl"}
