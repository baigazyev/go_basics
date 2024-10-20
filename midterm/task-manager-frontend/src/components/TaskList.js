import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

function TaskList() {
    const [tasks, setTasks] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8080/tasks')
            .then(response => {
                setTasks(response.data);
            })
            .catch(error => {
                console.error("There was an error fetching the tasks!", error);
            });
    }, []);

    return (
        <div className="container mt-5">
            <h2 className="text-center mb-4">Task List</h2>
            <ul className="list-group">
                {tasks.map(task => (
                    <li key={task.id} className="list-group-item d-flex justify-content-between align-items-center">
                        <Link to={`/task/${task.id}`} className="text-decoration-none">{task.title}</Link>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default TaskList;
