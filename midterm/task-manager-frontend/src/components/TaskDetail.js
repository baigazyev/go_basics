import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams, useNavigate } from 'react-router-dom';

function TaskDetail() {
    const [task, setTask] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const { id } = useParams();
    const navigate = useNavigate();

    useEffect(() => {
        axios.get(`http://localhost:8080/tasks/${id}`)
            .then(response => {
                setTask(response.data);
                setLoading(false);
            })
            .catch(err => {
                console.error("Error fetching task:", err);
                setError("Failed to fetch the task");
                setLoading(false);
            });
    }, [id]);

    const handleDelete = () => {
        axios.delete(`http://localhost:8080/tasks/${id}`)
            .then(() => {
                alert("Task deleted successfully!");
                navigate("/");
            })
            .catch(err => {
                console.error("Failed to delete the task:", err);
                alert("Failed to delete the task");
            });
    };

    const handleEdit = () => {
        navigate(`/edit-task/${id}`);
    };

    if (loading) {
        return <div>Loading...</div>;
    }

    if (error) {
        return <div>{error}</div>;
    }

    return (
        <div className="container mt-5">
            <h2>Task Details</h2>
            <p><strong>Title:</strong> {task.title}</p>
            <p><strong>Status:</strong> {task.status}</p>
            <button className="btn btn-danger" onClick={handleDelete}>Delete Task</button>
            <button className="btn btn-primary ms-3" onClick={handleEdit}>Edit Task</button> {/* Add Edit Button */}
        </div>
    );
}

export default TaskDetail;
