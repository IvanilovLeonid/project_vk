import React, { useEffect, useState } from "react";
import { getContainers } from "./api";

function Table() {
    const [containers, setContainers] = useState([]);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const data = await getContainers();
                setContainers(data);
            } catch (err) {
                setError("Не удалось загрузить данные.");
            }
        };

        fetchData();
    }, []);

    if (error) {
        return <div style={{ color: "red" }}>{error}</div>;
    }

    return (
        <table border="1" cellPadding="10" style={{ width: "100%", textAlign: "left" }}>
            <thead>
            <tr>
                <th>ID</th>
                <th>IP-адрес</th>
                <th>Последний пинг</th>
                <th>Последний успешный пинг</th>
            </tr>
            </thead>
            <tbody>
            {containers.map((container) => (
                <tr key={container.ID}>
                    <td>{container.ID}</td>
                    <td>{container.IPAddress}</td>
                    <td>{new Date(container.LastPingTime).toLocaleString()}</td>
                    <td>
                        {container.LastSuccessfulPing
                            ? new Date(container.LastSuccessfulPing).toLocaleString()
                            : "Нет данных"}
                    </td>
                </tr>
            ))}
            </tbody>
        </table>
    );
}

export default Table;
