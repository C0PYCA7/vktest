import React, { useState, useEffect } from 'react';
import { getContainers } from '../api/containerApi';

const ContainerList = () => {
    const [containers, setContainers] = useState([]);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchContainers = async () => {
            try {
                const data = await getContainers();
                console.log('Data received:', JSON.stringify(data, null, 2));
                if (data.error) {
                    setError(data.error);
                    setContainers([]);
                } else {
                    setContainers(data.containers);
                    setError(null);
                }
            } catch (err) {
                setError('Failed to fetch containers');
                setContainers([]);
                console.error('Error: ', err);
            }
        };
        fetchContainers();
    }, []);

    return (
        <div className="container mt-4">
            <h2>Container List</h2>
                <table className="table table-striped">
                    <thead>
                    <tr>
                        <th>Container IP</th>
                        <th>Ping Time</th>
                        <th>Last Success Date</th>
                    </tr>
                    </thead>
                    <tbody>
                    {containers.map((container, index) => (
                        <tr key={index}>
                            <td>{container.containerIP}</td>
                            <td>{container.pingTimeMKs} mks</td>
                            <td>{new Date(container.lastSuccessDate).toLocaleString()}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>
        </div>
    );
};

export default ContainerList;