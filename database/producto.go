package database

import (
	"encoding/json"
	"net/http"

	"github.com/M-2001/GoRestAPI/models"
	"github.com/go-chi/chi"
)

//catch funcion para errores
func catch(err error) {
	if err != nil {
		panic(err)
	}
}

//Deleteproduct servira para eliminar un producto de la base de datos
func Deleteproduct(w http.ResponseWriter, r *http.Request) {
	databaseConnection := ConnectDB()
	id := chi.URLParam(r, "id")
	query, err := databaseConnection.Prepare("delete from products where id = ?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	query.Close()
	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

//UpdateProduct servira para actulaizar un  producto en la base de datos
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	databaseConnection := ConnectDB()
	var producto models.Productos
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&producto)

	query, err := databaseConnection.Prepare("Update products SET product_code=?, description=? where id=?")
	catch(err)
	_, er := query.Exec(producto.ProductCode, producto.Description, id)
	catch(er)
	defer query.Close()
	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "update created"})
}

//CreateProduct servira para agregar un nuevo producto a la base de datos
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	databaseConnection := ConnectDB()
	var producto models.Productos
	json.NewDecoder(r.Body).Decode(&producto)

	query, err := databaseConnection.Prepare("Insert products SET product_code=?, description=?")
	catch(err)
	_, er := query.Exec(producto.ProductCode, producto.Description)
	catch(er)
	defer query.Close()
	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

//AllProductos nos devolvera un slice de productos
func AllProductos(w http.ResponseWriter, r *http.Request) {
	databaseConnection := ConnectDB()
	const sql = `SELECT id,product_code,COALESCE(description,'')FROM products`
	results, err := databaseConnection.Query(sql)
	catch(err)
	var products []*models.Productos

	for results.Next() {
		product := &models.Productos{}
		err = results.Scan(&product.ID, &product.ProductCode, &product.Description)
		catch(err)
		products = append(products, product)
	}
	respondwithJSON(w, http.StatusOK, products)
}
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type-", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
