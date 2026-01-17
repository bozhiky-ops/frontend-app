// types.ts

export enum Status {
  PENDING = 'PENDING',
  SUCCESS = 'SUCCESS',
  ERROR = 'ERROR',
}

export type Product = {
  id: number;
  name: string;
  price: number;
  status: Status;
};

export type User = {
  id: number;
  name: string;
  email: string;
};

export type Order = {
  id: number;
  userId: number;
  productId: number;
  quantity: number;
  status: Status;
};

export type ProductResponse = {
  data: Product[];
  error?: string;
};

export type UserResponse = {
  data: User[];
  error?: string;
};

export type OrderResponse = {
  data: Order[];
  error?: string;
};