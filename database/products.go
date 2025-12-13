package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
}

var productList []Product

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}

	return nil
}

func Update(p Product) {
	for index, product := range productList {
		if product.ID == p.ID {
			productList[index] = product
		}
	}
}

func Delete(pID int) {
	var temp []Product
	for index, product := range productList {
		if product.ID != pID {
			temp[index] = product
		}
	}

	productList = temp
}

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

	productList = append(productList, prod1)
	productList = append(productList, prod2)
	productList = append(productList, prod3)
}
