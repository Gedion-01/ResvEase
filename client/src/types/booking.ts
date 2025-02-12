import type { Hotel, Room } from "./hotel";

export interface Booking {
  id: string;
  userID: string;
  roomID: string;
  firstName: string;
  lastName: string;
  email: string;
  phone: string;
  specialRequests: string;
  numPersons: number;
  fromDate: string;
  tillDate: string;
  cancelled: boolean;
  status: "pending" | "confirmed" | "cancelled";
  hotel: Hotel;
  room: Room;
  totalPrice: number;
  createdAt: string;
  updatedAt: string;
}

export interface BookingsResponse {
  results: number;
  data: Booking[];
  page: number;
}
