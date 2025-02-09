export const getContainers = async () => {
    try {
        const response = await fetch('/containers');
        if (!response.ok) {
            throw new Error('Failed to fetch containers');
        }
        return await response.json();
    } catch (error) {
        console.error('Error: ', error);
        return { containers: [], error: error.message };
    }
};

export const createContainer = async (containerIP) => {
    try {
        const response = await fetch('/createcontainer', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ containerIP }),
        });
        if (!response.ok) {
            throw new Error('Error creating container');
        }
        return await response.json();
    } catch (error) {
        console.error('Error:', error);
        return { error: error.message };
    }
};