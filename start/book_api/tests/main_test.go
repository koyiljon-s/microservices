package tests

import (
	// other imports
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("SECRET_TOKEN"))

func setupTestDB() {
	// setupTestDB goes here
}

func addBook() api.Book {
	// add book code goes here
}

func generateValidToken() string {
	expirationTime := time.Now().Add(15 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expirationTime.Unix(),
	})
	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}

func TestGenerateJWT(t *testing.T) {
	router := gin.Default()
	router.POST("/token", api.GenerateJWT)

	loginRequest := map[string]string{
		"username": "admin",
		"password": "password",
	}

	jsonValue, _ := json.Marshal(loginRequest)
	req, _ := http.NewRequest("POST", "/token", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}

	var response api.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil || response.Data.(map[string]interface{})["token"] == "" {
		t.Errorf("Expected token in response, got nil or empty")
	}
}


func TestCreateBook(t *testing.T) {
	setupTestDB()
	router := gin.Default()
	protected := router.Group("/", api.JWTAuthMiddleware()) // add
	protected.POST("/book", api.CreateBook)                 // add

	token := generateValidToken() // add

	book := api.Book{
		Title: "Demo Book name", Author: "Demo Author name", Year: 2021,
	}
	jsonValue, _ := json.Marshal(book)

	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", token) // add

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, status)
	}

	var response api.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil {
		t.Errorf("Expected book data, got nil")
	}
}
