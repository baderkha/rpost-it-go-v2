package view

type IView interface {
	Render() (string,error)
}

