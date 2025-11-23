package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"shopping-cart-backend/models"
	"shopping-cart-backend/routes"
)

var _ = Describe("Shopping Cart API", func() {
	var (
		router *gin.Engine
		db     *gorm.DB
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		db = setupTestDB()
		router = routes.SetupRouter(db)
	})

	It("allows a user to sign up", func() {
		resp := performRequest(router, "POST", "/users", map[string]string{
			"username": "alice",
			"password": "pass",
		}, "")

		Expect(resp.Code).To(Equal(http.StatusCreated))
	})

	It("returns a token on login", func() {
		performRequest(router, "POST", "/users", map[string]string{
			"username": "bob",
			"password": "1234",
		}, "")

		resp := performRequest(router, "POST", "/users/login", map[string]string{
			"username": "bob",
			"password": "1234",
		}, "")

		Expect(resp.Code).To(Equal(http.StatusOK))
		var body map[string]string
		json.Unmarshal(resp.Body.Bytes(), &body)
		Expect(body["token"]).ToNot(BeEmpty())

		var user models.User
		err := db.Where("username = ?", "bob").First(&user).Error
		Expect(err).To(BeNil())
		Expect(user.CurrentToken).To(Equal(body["token"]))
	})

	It("creates a store item", func() {
		resp := performRequest(router, "POST", "/items", map[string]interface{}{
			"name":  "Pen",
			"price": 2.5,
		}, "")
		Expect(resp.Code).To(Equal(http.StatusCreated))

		var stored models.Item
		err := db.First(&stored, "name = ?", "Pen").Error
		Expect(err).To(BeNil())
		Expect(stored.Price).To(Equal(2.5))
	})

	It("adds an item to the cart", func() {
		token := seedUserAndLogin(router, "cara", "pass")
		itemID := createItemAndGetID(router, "Marker", 3.0)

		resp := performRequest(router, "POST", "/carts", map[string]interface{}{
			"item_id":  itemID,
			"quantity": 2,
		}, token)

		Expect(resp.Code).To(Equal(http.StatusOK))

		var cartItem models.CartItem
		err := db.First(&cartItem).Error
		Expect(err).To(BeNil())
		Expect(cartItem.Quantity).To(Equal(2))
	})

	It("checks out a cart into an order and clears it", func() {
		token := seedUserAndLogin(router, "dave", "pass")
		itemID := createItemAndGetID(router, "Notebook", 5.0)

		performRequest(router, "POST", "/carts", map[string]interface{}{
			"item_id":  itemID,
			"quantity": 1,
		}, token)

		resp := performRequest(router, "POST", "/orders", nil, token)
		Expect(resp.Code).To(Equal(http.StatusOK))

		var orders []models.Order
		db.Find(&orders)
		Expect(len(orders)).To(Equal(1))

		var cartItems []models.CartItem
		db.Find(&cartItems)
		Expect(cartItems).To(BeEmpty())
	})
})

func setupTestDB() *gorm.DB {
	tmp, err := os.CreateTemp("", "cart-test-*.db")
	Expect(err).To(BeNil())
	path := tmp.Name()
	tmp.Close()
	os.Remove(path)

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	Expect(err).To(BeNil())

	err = db.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)
	Expect(err).To(BeNil())

	DeferCleanup(func() {
		os.Remove(path)
	})

	return db
}

func performRequest(router *gin.Engine, method, path string, payload interface{}, token string) *httptest.ResponseRecorder {
	var body *bytes.Buffer
	if payload != nil {
		data, _ := json.Marshal(payload)
		body = bytes.NewBuffer(data)
	} else {
		body = bytes.NewBuffer([]byte{})
	}

	req, _ := http.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func seedUserAndLogin(router *gin.Engine, username, password string) string {
	performRequest(router, "POST", "/users", map[string]string{
		"username": username,
		"password": password,
	}, "")

	resp := performRequest(router, "POST", "/users/login", map[string]string{
		"username": username,
		"password": password,
	}, "")

	var body map[string]string
	json.Unmarshal(resp.Body.Bytes(), &body)
	return body["token"]
}

func createItemAndGetID(router *gin.Engine, name string, price float64) interface{} {
	resp := performRequest(router, "POST", "/items", map[string]interface{}{
		"name":  name,
		"price": price,
	}, "")
	Expect(resp.Code).To(Equal(http.StatusCreated))

	var item map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &item)
	return item["ID"]
}

