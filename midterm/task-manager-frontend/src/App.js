import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import TaskList from './components/TaskList';
import TaskDetail from './components/TaskDetail';
import AddTask from './components/AddTask';
import EditTask from './components/EditTask'; // Import the new EditTask component

function App() {
    return (
        <Router>
            <div className="App">
                <h1>Task Manager</h1>
                <Routes>
                    <Route path="/" element={<TaskList />} />
                    <Route path="/task/:id" element={<TaskDetail />} />
                    <Route path="/add-task" element={<AddTask />} />
                    <Route path="/edit-task/:id" element={<EditTask />} /> {/* Add the new edit route */}
                </Routes>
            </div>
        </Router>
    );
}

export default App;
