// src/App.js
import React, { useState } from 'react';
import './App.css';
import BookList from './components/BookList';
import LoginForm from './components/LoginForm'; // Import your LoginForm component
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false); // State to manage login status

  const handleLogin = () => {
    setIsLoggedIn(true); // Set logged in state to true
  };

  const handleLogout = () => {
    localStorage.removeItem('token'); // Clear the token from local storage
    setIsLoggedIn(false); // Set logged in state to false
    // Additional logout logic (like redirecting)
  };

  return (
    <div className="App">
      {!isLoggedIn ? (
        <LoginForm onLogin={handleLogin} /> // Render LoginForm if not logged in
      ) : (
        <>
          <BookList />
          <button className="btn btn-danger" onClick={handleLogout}>Logout</button> {/* Logout button */}
        </>
      )}
    </div>
  );
}

export default App;
