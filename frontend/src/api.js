import axios from "axios";

const API_BASE_URL = "http://localhost:8081";

export const addContainer = async (ipAddress) => {
    const response = await axios.post(`${API_BASE_URL}/containers`, { IPAddress: ipAddress });
    return response.data;
};