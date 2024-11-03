// src/components/ProtectedComponent.js
import React, { useEffect, useState } from 'react';
import axiosInstance from '../axiosInstance';

const ProtectedComponent = () => {
  const [data, setData] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axiosInstance.get('/protected-endpoint');
        setData(response.data);
      } catch (error) {
        console.error('Error fetching protected data:', error);
      }
    };

    fetchData();
  }, []);

  if (!data) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h1>Protected Data</h1>
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </div>
  );
};

export default ProtectedComponent;
