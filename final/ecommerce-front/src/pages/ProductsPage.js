// src/pages/ProductsPage.js
import React, { useEffect, useState } from "react";
import { fetchProducts } from "../api/products";

const ProductsPage = () => {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    fetchProducts()
      .then((data) => setProducts(data))
      .catch((err) => console.error("Error fetching products:", err));
  }, []);

  return (
    <div>
      <h1>Products</h1>
      <ul>
        {products.map((product) => (
          <li key={product.id}>
            {product.name} - ${product.price}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ProductsPage;
