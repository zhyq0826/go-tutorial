package lib

var favorites []string

func init() {
	favorites = make([]string, 3)
	favorites[0] = "github.com/gorilla/mux"
	favorites[1] = "github.com/codegangsta/negroni"
	favorites[1] = "gopkg.in/mgo.v2"
}

func Add(favorite string) {
	favorites = append(favorites, favorite)
}

func GetAll() []string {
	return favorites
}
