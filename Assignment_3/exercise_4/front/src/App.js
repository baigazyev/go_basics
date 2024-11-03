// src/App.js
import React from 'react';
import './App.css';
import BookList from './components/BookList';
// src/index.js or src/App.js
import 'bootstrap/dist/css/bootstrap.min.css';


function App() {
  return (
    <div className="App">
      <BookList />
    </div>
  );
}

export default App;
