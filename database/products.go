package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
}

var ProductList []Product

func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red. I like orange",
		Price:       100,
		ImageURL:    "https://5.imimg.com/data5/SR/NF/GT/SELLER-2706786/orange-emulsion-1000x1000.jpg",
	}
	prod2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green. I don't like apple",
		Price:       50,
		ImageURL:    "https://static.wikia.nocookie.net/the-snack-encyclopedia/images/7/7d/Apple.png/revision/latest?cb=20200706145821",
	}
	prod3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow. I like it",
		Price:       50,
		ImageURL:    "https://fruitfortheoffice.co.uk/media/.renditions/wysiwyg/42e9as7nataai4a6jcufwg.jpeg",
	}

	ProductList = append(ProductList, prod1)
	ProductList = append(ProductList, prod2)
	ProductList = append(ProductList, prod3)
}
