import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams, useNavigate } from 'react-router-dom';

function EditTask() {
    const [title, setTitle] = useState('');
    const [status, setStatus] = useState('');
    const { id } = useParams();
    const navigate = useNavigate();

    useEffect(() => {
        
        axios.get(`http://localhost:8080/tasks/${id}`)
            .then(response => {
                setTitle(response.data.title);
                setStatus(response.data.status);
            })
            .catch(err => {
                console.error("Error fetching task:", err);
            });
    }, [id]);

    const handleSubmit = (e) => {
        e.preventDefault();

        
        axios.put(`http://localhost:8080/tasks/${id}`, { title, status })
            .then(() => {
                alert("Task updated successfully!");
                navigate(`/task/${id}`); 
            })
            .catch(err => {
                console.error("Failed to update the task:", err);
                alert("Failed to update the task");
            });
    };

    return (
        <div className="container mt-5">
            <h2>Edit Task</h2>
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
                <button type="submit" className="btn btn-primary w-100">Update Task</button>
            </form>
        </div>
    );
}

export default EditTask;
