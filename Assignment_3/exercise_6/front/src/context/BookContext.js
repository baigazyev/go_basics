// src/context/BookContext.js
import React, { createContext, useState, useEffect } from 'react';
import axios from 'axios';

export const BookContext = createContext();

const BookProvider = ({ children }) => {
  const [books, setBooks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        const response = await axios.get('http://localhost:9090/books');
        setBooks(response.data);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchBooks();
  }, []);

  const addBook = async (newBook) => {
    setLoading(true); // Set loading to true when adding a book
    setError(null); // Reset the error before the new request

    try {
      const response = await axios.post('http://localhost:9090/books', newBook);
      setBooks((prevBooks) => [...prevBooks, response.data]);
    } catch (err) {
      setError(err.message); // Capture the error message
    } finally {
      setLoading(false); // Reset loading after the request
    }
  };

  return (
    <BookContext.Provider value={{ books, loading, error, addBook }}>
      {children}
    </BookContext.Provider>
  );
};

export default BookProvider;
