# 🍦 CRUD-REST-API-GO  

🚀 A **simple and efficient** REST API built with **Go (Gin Framework)** and **MongoDB** to manage ice cream flavors!  

---

## 🏗️ Tech Stack  
🔹 **Go** – High-performance backend language  
🔹 **Gin** – Fast and lightweight web framework  
🔹 **MongoDB** – NoSQL database for storing flavors  
🔹 **Go Modules** – Dependency management  

---

## 🌟 Features  
✅ **Flavour Management** – Add, update, delete, and retrieve flavors  
✅ **MongoDB Integration** – Stores flavors with unique IDs  
✅ **RESTful API** – Clean and structured API design  
✅ **Fast & Scalable** – Optimized for performance  

---

## 📌 API Endpoints  

### 📜 Flavour Management  
🔹 **Get All Flavours:** `GET /api/flavours`  
🔹 **Get Flavour by ID:** `GET /api/flavours/:id`  
🔹 **Add a Flavour:** `POST /api/flavours`  
🔹 **Update a Flavour:** `PUT /api/flavours/:id`  
🔹 **Delete a Flavour:** `DELETE /api/flavours/:id`  

---

## 🛠 Database Schema (MongoDB)  
```go
type Flavour struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    FlavourName string             `bson:"flavour_name" json:"flavour_name"`
    Quantity    string             `bson:"qty" json:"qty"`
}
