import React, { useState } from 'react';
import axios from 'axios';

function AddTask() {
    const [title, setTitle] = useState('');
    const [status, setStatus] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();

        axios.post('http://localhost:8080/tasks', { title, status })
            .then(response => {
                console.log("Task added:", response.data);
                setTitle('');
                setStatus('');
                window.location.reload(); 
            })
            .catch(error => {
                console.error("There was an error adding the task!", error);
            });
    };

    return (
        <div className="container mt-5">
            <h2 className="text-center mb-4">Add Task</h2>
            <form onSubmit={handleSubmit}>
                <div className="form-group mb-3">
                    <input
                        type="text"
                        className="form-control"
                        placeholder="Task Title"
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                    />
                </div>
                <div className="form-group mb-3">
                    <input
                        type="text"
                        className="form-control"
                        placeholder="Task Status"
                        value={status}
                        onChange={(e) => setStatus(e.target.value)}
                    />
                </div>
                <button type="submit" className="btn btn-primary w-100">Add Task</button>
            </form>
        </div>
    );
}

export default AddTask;
