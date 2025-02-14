export interface HotelResponse {
  results: number;
  data: Hotel[];
  page: number;
  totalPage: number;
}

export interface RoomResponse {
  results: number;
  data: Room[];
  page: number;
  totalPage: number;
}

export interface Hotel {
  id: string;
  name: string;
  images: string[];
  location: string;
  description: string;
  rating: number;
  amenities: string[];
  rooms: null;
  prices: number[];
}

export interface Room {
  id: string;
  name: string;
  description: string;
  price: number;
  capacity: number;
  bedType: string;
  bedRooms: number;
  availableCount: number;
  amenities: string[];
  images: string[];
}
