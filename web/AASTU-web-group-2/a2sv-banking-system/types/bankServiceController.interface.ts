
export interface BankServiceData {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
}

export interface ApiResponse<T> {
  success: boolean;
  message: string;
  data: T;
}
