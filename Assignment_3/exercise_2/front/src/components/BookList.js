// src/components/BookList.js
import React, { useEffect, useState } from 'react';
import axios from 'axios';

const BookList = () => {
  const [books, setBooks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        const response = await axios.get('http://localhost:9090/books');
        setBooks(response.data);
      } catch (err) {
        setError(err);
      } finally {
        setLoading(false);
      }
    };

    fetchBooks();
  }, []);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error fetching books: {error.message}</p>;

  return (
    <div>
      <h1>Book List</h1>
      <ul>
        {books.map((book) => (
          <li key={book.ID}>
            <h2>{book.Title}</h2>
            <p>{book.Author}</p>
            <p>{book.Year}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default BookList;
