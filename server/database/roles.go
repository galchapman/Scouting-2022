package database

func parseRole(role string) int {
	switch role {
	case "Scouter":
		return ScouterRole
	case "Viewer":
		return ViewerRole
	case "FieldPlayer":
		return FieldPlayerRole
	case "Manager":
		return Manager
	case "Admin":
		return AdminRole
	default:
		return -1 // error
	}
}
