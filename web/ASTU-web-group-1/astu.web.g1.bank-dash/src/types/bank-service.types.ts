export interface BankServiceDataType {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
}
export interface BankServiceResponseType {
  success: boolean;
  message: string;
  data: {
    content: BankServiceDataType[];
    totalPages: number;
  };
}

export interface singleBankServiceResponseType {
  success: boolean;
  message: string;
  data: BankServiceDataType;
}
