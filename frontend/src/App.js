import React, { useState } from 'react';
import ContainerList from './components/ContainerList';
import CreateContainer from './components/CreateContainer';

const App = () => {
    const [isContainerCreated, setIsContainerCreated] = useState(false);

    const handleContainerCreated = () => {
        setIsContainerCreated(!isContainerCreated);
    };

    return (
        <div className="container d-flex flex-column align-items-center justify-content-center min-vh-100">
            <h1 className="mb-4">Container Management</h1>
            <CreateContainer onContainerCreated={handleContainerCreated} />
            <ContainerList />
        </div>
    );
};

export default App;