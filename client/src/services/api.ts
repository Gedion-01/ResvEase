import { useAuthStore } from "@/store/authStore";
import { type Booking, type BookingsResponse } from "@/types/booking";
import {
  type RoomResponse,
  type HotelResponse,
  type Hotel,
} from "@/types/hotel";
import axios from "axios";

export const BASE_URL = "http://localhost:5000/api/v1";

const apiClient = axios.create({ baseURL: BASE_URL });

apiClient.interceptors.request.use((config) => {
  const authStore = useAuthStore();
  const token = authStore.token;

  console.log("token", token);

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
      // Token is invalid or expired
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
    console.log(queryString, queryParams);
    const response = await apiClient.get<HotelResponse>(
      `/hotel?${queryString}`
    );
    // return response.data;
    const { data, results, page } = response.data;
    return {
      pageData: data || [],
      cursor: results > 0 ? page + 1 : undefined,
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
  console.log(queryString);
  const response = await apiClient.get<RoomResponse>(
    `/hotel/${hotelId}/rooms?${queryString}`
  );
  return response.data;
};
// `/hotel/${hotelId}/rooms?checkIn=2025-01-01&checkOut=2025-01-05&page=1&limit=10&hotelRating=4.8&RoomCapacity=2`;

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
    console.log(err);
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
