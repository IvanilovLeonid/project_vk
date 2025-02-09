import React, { useEffect, useState } from 'react';

const Table = () => {
    const [data, setData] = useState([]);

    useEffect(() => {
        // Замените URL на ваш Backend-сервис
        fetch('http://localhost:8081/containers')
            .then((response) => response.json())
            .then((data) => setData(data))
            .catch((error) => console.error('Error fetching data:', error));
    }, []);

    return (
        <div>
            <h1>Container Data</h1>
            <table>
                <thead>
                <tr>
                    <th>IP Address</th>
                    <th>Last Ping Time</th>
                    <th>Last Successful Ping</th>
                </tr>
                </thead>
                <tbody>
                {data.map((item) => (
                    <tr key={item.id}>
                        <td>{item.ip_address}</td>
                        <td>{new Date(item.last_ping_time).toLocaleString()}</td>
                        <td>{new Date(item.last_successful_ping).toLocaleString()}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

export default Table;
