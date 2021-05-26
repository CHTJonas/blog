package blog

type templateData struct {
	Description string
	Author      string
	Title       string
	ContentData interface{}
}

func (serv *Server) getTemplateData() *templateData {
	return &templateData{
		Description: serv.b.Description,
		Author:      serv.b.OwnerName,
		Title:       serv.b.Title,
	}
}
