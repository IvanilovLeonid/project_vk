import React, { useState } from "react";
import Table from "./Table";
import { pingContainer } from "./api";

function App() {
    const [ipAddress, setIpAddress] = useState("");
    const [message, setMessage] = useState("");
    const [success, setSuccess] = useState(true);  // Добавим состояние для успеха пинга

    const handlePingContainer = async (e) => {
        e.preventDefault();
        try {
            const response = await pingContainer(ipAddress, success);
            setMessage(response.message);
            setIpAddress("");
        } catch (err) {
            setMessage("Ошибка при отправке пинга.");
        }
    };

    return (
        <div style={{ padding: "20px" }}>
            <h1>Контейнеры</h1>

            {/* Форма для отправки пинга контейнера */}
            <form onSubmit={handlePingContainer} style={{ marginBottom: "20px" }}>
                <input
                    type="text"
                    placeholder="Введите IP-адрес"
                    value={ipAddress}
                    onChange={(e) => setIpAddress(e.target.value)}
                    required
                    style={{ padding: "8px", marginRight: "10px", width: "300px" }}
                />
                <select
                    value={success ? "true" : "false"}
                    onChange={(e) => setSuccess(e.target.value === "true")}
                    style={{ padding: "8px", marginRight: "10px" }}
                >
                    <option value="true">Успешный пинг</option>
                    <option value="false">Неудачный пинг</option>
                </select>
                <button type="submit" style={{ padding: "8px 16px" }}>
                    Отправить пинг
                </button>
            </form>

            {message && <p style={{ color: "green" }}>{message}</p>}
            <Table />
        </div>
    );
}

export default App;
