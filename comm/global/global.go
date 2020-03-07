package global 

type EndPoint struct{
	Ip string
	Port uint32
}

func GetPortForSvr(strSvrName  string)int32{
		switch strSvrName{
			case "whinvitesvr":
				return 50051
			default:
				return 0
		}
}

func GetEndPointForClient(strSvrName string, uId uint64) *EndPoint{
		switch strSvrName{
			case "whinvitesvr":
				return &EndPoint{"49.235.67.28", 50051}
			default:
				return nil
		}
}
