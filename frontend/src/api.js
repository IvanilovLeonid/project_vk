import axios from "axios";

const API_BASE_URL = "http://localhost:8081";

// Функция для получения всех контейнеров
export const getContainers = async () => {
    const response = await axios.get(`${API_BASE_URL}/containers`);
    return response.data;
};

// Функция для отправки POST-запроса с пингом
export const pingContainer = async (ipAddress, success) => {
    const response = await axios.post(`${API_BASE_URL}/ping`, {
        ip_address: ipAddress,
        success: success,
    });
    return response.data;
};
