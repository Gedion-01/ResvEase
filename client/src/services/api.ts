import { type Hotel } from "@/types/hotel";
import axios from "axios";

const BASE_URL = "http://localhost:5000/api/v1";

const axiosInstance = axios.create({ baseURL: BASE_URL });

export const getHotels = async () => {
  return await axiosInstance.get<Hotel[]>("hotel?page=1&limit=10");
};
