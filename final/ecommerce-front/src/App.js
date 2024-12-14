import React, { useState, useEffect } from "react";
import { fetchProducts, fetchProductById } from "./api/products";
import "./styles/App.css";

function App() {
  const [products, setProducts] = useState([]);
  const [selectedProduct, setSelectedProduct] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchProducts()
      .then((data) => setProducts(data))
      .catch((err) => setError("Failed to fetch products."));
  }, []);

  const handleFetchProduct = (id) => {
    fetchProductById(id)
      .then((data) => setSelectedProduct(data))
      .catch((err) => setError("Failed to fetch product details."));
  };

  return (
    <div>
      <h1>Products</h1>

      {error && <p style={{ color: "red" }}>{error}</p>}

      <ul>
        {products.map((product) => (
          <li key={product.id}>
            {product.name} - ${product.price.toFixed(2)}
            <button onClick={() => handleFetchProduct(product.id)} style={{ marginLeft: "10px" }}>
              Details
            </button>
          </li>
        ))}
      </ul>

      {selectedProduct && (
        <div style={{ marginTop: "20px" }}>
          <h2>Product Details</h2>
          <p><strong>Name:</strong> {selectedProduct.name}</p>
          <p><strong>Description:</strong> {selectedProduct.description}</p>
          <p><strong>Price:</strong> ${selectedProduct.price.toFixed(2)}</p>
          <p><strong>Stock:</strong> {selectedProduct.stock}</p>
        </div>
      )}
    </div>
  );
}

export default App;
