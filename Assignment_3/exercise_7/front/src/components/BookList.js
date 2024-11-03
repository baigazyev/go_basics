// src/components/BookList.js
import React, { useContext, useState } from 'react';
import { BookContext } from '../context/BookContext';

const BookList = () => {
  const { books, loading, error, addBook } = useContext(BookContext);
  const [newBook, setNewBook] = useState({
    title: '',
    author: '',
    description: '',
  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setNewBook((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    addBook(newBook);
    setNewBook({ title: '', author: '', description: '' }); // Reset form
  };

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error fetching books: {error}</p>;

  return (
    <div className="container mt-4">
      <h1>Book List</h1>
      <ul className="list-group mb-4">
        {books.map((book) => (
          <li key={book.ID} className="list-group-item">
            <h2>{book.Title}</h2>
            <p>{book.Author}</p>
            <p>{book.Year}</p>
          </li>
        ))}
      </ul>

      <h2>Add a New Book</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <input
            type="text"
            name="title"
            placeholder="Title"
            value={newBook.title}
            onChange={handleInputChange}
            required
            className="form-control"
          />
        </div>
        <div className="mb-3">
          <input
            type="text"
            name="author"
            placeholder="Author"
            value={newBook.author}
            onChange={handleInputChange}
            required
            className="form-control"
          />
        </div>
        <div className="mb-3">
          <textarea
            name="description"
            placeholder="Description"
            value={newBook.description}
            onChange={handleInputChange}
            required
            className="form-control"
          />
        </div>
        <button type="submit" className="btn btn-primary">Add Book</button>
      </form>
    </div>
  );
};

export default BookList;
