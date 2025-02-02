import { useAuthStore } from "@/store/authStore";
import {
  type RoomResponse,
  type HotelResponse,
  type Hotel,
} from "@/types/hotel";
import axios from "axios";

const BASE_URL = "http://localhost:5000/api/v1";

const apiClient = axios.create({ baseURL: BASE_URL });

apiClient.interceptors.request.use((config) => {
  const authStore = useAuthStore();
  const token = authStore.token;

  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export const getHotel = async (hotelId: string) => {
  const response = await apiClient.get<Hotel>(`/hotel/${hotelId}`);
  return response.data;
};

export const getHotels = async (queryParams: Record<string, any>) => {
  try {
    const queryString = new URLSearchParams(queryParams).toString();
    console.log(queryString, queryParams);
    const response = await apiClient.get<HotelResponse>(
      `/hotel?page=1&limit=10&${queryString}`
    );
    return response.data;
  } catch (error) {
    if (axios.isAxiosError(error) && error.response?.status === 404) {
      console.error("Hotel resource not found:", error.response.data);
      return { data: [], results: 0, page: 1 };
    } else {
      console.error("An error occurred while fetching hotels:", error);
      throw error;
    }
  }
};

export const getRooms = async (
  hotelId: string,
  queryParams: Record<string, any>
) => {
  const queryString = new URLSearchParams(queryParams).toString();
  console.log(queryString);
  const response = await apiClient.get<RoomResponse>(
    `/hotel/${hotelId}/rooms?${queryString}`
  );
  return response.data;
};
// `/hotel/${hotelId}/rooms?checkIn=2025-01-01&checkOut=2025-01-05&page=1&limit=10&hotelRating=4.8&RoomCapacity=2`;

export default apiClient;
