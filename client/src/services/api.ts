import { useAuthStore } from "@/store/authStore";
import { type Booking, type BookingsResponse } from "@/types/booking";
import {
  type RoomResponse,
  type HotelResponse,
  type Hotel,
} from "@/types/hotel";
import axios from "axios";

export const BASE_URL = import.meta.env.VITE_API_ENDPOINT;

const apiClient = axios.create({ baseURL: BASE_URL });

apiClient.interceptors.request.use((config) => {
  const authStore = useAuthStore();
  const token = authStore.token;

  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    const authStore = useAuthStore();

    if (error.response.status === 401) {
      // Token invalid or expired
      authStore.logout();
    }
    return Promise.reject(error);
  }
);

export const getHotel = async (hotelId: string) => {
  const response = await apiClient.get<Hotel>(`/hotel/${hotelId}`);
  return response.data;
};

export const getHotels = async (queryParams: Record<string, any>) => {
  try {
    const queryString = new URLSearchParams(queryParams).toString();
    const response = await apiClient.get<HotelResponse>(
      `/hotel?${queryString}`
    );

    const { data, results, page, totalPage } = response.data;
    return {
      pageData: data || [],
      cursor: page < totalPage ? page + 1 : undefined,
    };
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
  const response = await apiClient.get<RoomResponse>(
    `/hotel/${hotelId}/rooms?${queryString}`
  );
  const { data, page, totalPage } = response.data;
  return {
    pageData: data || [],
    cursor: page < totalPage ? page + 1 : undefined,
  };
};

interface BookingData {
  hotelID: string;
  roomName: string;
  firstName: string;
  lastName: string;
  phone: string;
  specialRequests: string;
  fromDate: string;
  tillDate: string;
  numPersons: number;
}
export const bookRoom = async (data: BookingData) => {
  try {
    const response = await apiClient.post("/room/book", data);
    return response;
  } catch (err) {
    throw err;
  }
};

export const getUserBookings = async (userID: string) => {
  const response = await apiClient.get<BookingsResponse>(
    `/bookings/user/${userID}?page=1&limit=10`
  );
  return response.data;
};

export const getBooking = async (bookingID: string) => {
  const response = await apiClient.get<Booking>(`/booking/${bookingID}`);
  return response.data;
};
export default apiClient;
