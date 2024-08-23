export interface BankService {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
}

export interface BankServiceData {
  content: BankService[];
  totalPages: number;
}

export interface ApiResponse {
  success: boolean;
  message: string;
  data: BankServiceData;
}

export interface DeleteResponse {
  success: boolean;
  message: string;
  data: Record<string, never>;
}

export interface PostResponseData {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
}

export interface PostResponse {
  success: boolean;
  message: string;
  data: PostResponseData;
}
export interface SearchResponseData {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
}

export interface SearchResponse {
  success: boolean;
  message: string;
  data: SearchResponseData[];
}
export interface PutBankServiceData {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
}

export interface PutBankServiceResponse {
  success: boolean;
  message: string;
  data: PutBankServiceData;
}
export interface GetByIdResponseData {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
}

export interface GetByIdResponse {
  success: boolean;
  message: string;
  data: GetByIdResponseData;
}
