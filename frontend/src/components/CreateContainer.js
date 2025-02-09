import React, { useState } from 'react';
import { createContainer } from '../api/containerApi';

const CreateContainer = ({ onContainerCreated }) => {
    const [containerIP, setContainerIP] = useState('');
    const [error, setError] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        const response = await createContainer(containerIP);
        
        if (response.error) {
            setError(response.error);
        } else {
            setError('');
            setContainerIP('');
            onContainerCreated();
        }
    };

    return (
        <div className="container mt-4">
            <h2>Add Container</h2>
            {error && <div className="alert alert-danger">{error}</div>}
            <form onSubmit={handleSubmit}>
                <div className="mb-3">
                    <label htmlFor="containerIP" className="form-label">Container IP:</label>
                    <input
                        type="text"
                        id="containerIP"
                        value={containerIP}
                        onChange={(e) => setContainerIP(e.target.value)}
                        className="form-control"
                        required
                    />
                </div>
                <button type="submit" className="btn btn-primary">Create</button>
            </form>
        </div>
    );
};

export default CreateContainer;