package database

func ParseRole(role string) int {
	switch role {
	case "Scouter":
		return ScouterRole
	case "Viewer":
		return ViewerRole
	case "Supervisor":
		return SupervisorRole
	case "Manager":
		return ManagerRole
	case "Admin":
		return AdminRole
	default:
		return -1 // error
	}
}

func GetRole(role int) string {
	switch role {
	case ScouterRole:
		return "Scouter"
	case ViewerRole:
		return "Viewer"
	case SupervisorRole:
		return "Supervisor"
	case ManagerRole:
		return "Manager"
	case AdminRole:
		return "Admin"
	default:
		return "" // error
	}
}
