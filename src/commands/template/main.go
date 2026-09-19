package template

func Main(args []string, operation string) {

	templateName := args[0]

	switch operation {
	case "add":
		AddTemplate(templateName)
	case "create":
		CreateTemplate(templateName)
	case "delete":
		DeleteTemplate(templateName)
	}
}