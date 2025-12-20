package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
}

type ProductRepo interface {
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Create(p Product) (*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

// constructor or constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}

	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Delete(productID int) error {
	var tempList []*Product
	for _, product := range r.productList {
		if product.ID != productID {
			tempList = append(tempList, product)
		}
	}

	r.productList = tempList
	return nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	for index, p := range r.productList {
		if p.ID == product.ID {
			r.productList[index] = &product
		}
	}

	return &product, nil
}

func generateInitialProducts(r *productRepo) {
	prod1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red. I like orange",
		Price:       100,
		ImageURL:    "https://5.imimg.com/data5/SR/NF/GT/SELLER-2706786/orange-emulsion-1000x1000.jpg",
	}
	prod2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green. I don't like apple",
		Price:       50,
		ImageURL:    "https://static.wikia.nocookie.net/the-snack-encyclopedia/images/7/7d/Apple.png/revision/latest?cb=20200706145821",
	}
	prod3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow. I like it",
		Price:       50,
		ImageURL:    "https://fruitfortheoffice.co.uk/media/.renditions/wysiwyg/42e9as7nataai4a6jcufwg.jpeg",
	}

	r.productList = append(r.productList, prod1)
	r.productList = append(r.productList, prod2)
	r.productList = append(r.productList, prod3)
}
