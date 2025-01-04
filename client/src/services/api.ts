import {
  type RoomResponse,
  type HotelResponse,
  type Hotel,
} from "@/types/hotel";
import axios from "axios";

const BASE_URL = "http://localhost:5000/api/v1";

const axiosInstance = axios.create({ baseURL: BASE_URL });

export const getHotel = async (hotelId: string) => {
  const response = await axiosInstance.get<Hotel>(`/hotel/${hotelId}`);
  return response.data;
};

export const getHotels = async () => {
  const response = await axiosInstance.get<HotelResponse>(
    "hotel?page=1&limit=10"
  );
  return response.data;
};

export const getRooms = async (
  hotelId: string,
  queryParams: Record<string, any>
) => {
  const queryString = new URLSearchParams(queryParams).toString();
  console.log(queryString);
  const response = await axiosInstance.get<RoomResponse>(
    `/hotel/${hotelId}/rooms?${queryString}`
  );
  return response.data;
};
// `/hotel/${hotelId}/rooms?checkIn=2025-01-01&checkOut=2025-01-05&page=1&limit=10&hotelRating=4.8&RoomCapacity=2`;
