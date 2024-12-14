// src/api/products.js

export const fetchProducts = async () => {
    const response = await fetch("http://localhost:8080/api/products");
    if (!response.ok) {
      throw new Error("Failed to fetch products");
    }
    return await response.json();
  };
  
  export const fetchProductById = async (id) => {
    const response = await fetch(`http://localhost:8080/api/products/${id}`);
    if (!response.ok) {
      throw new Error("Failed to fetch product details");
    }
    return await response.json();
  };
  