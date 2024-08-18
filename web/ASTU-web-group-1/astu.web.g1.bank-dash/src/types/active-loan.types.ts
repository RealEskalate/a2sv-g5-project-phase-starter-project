export interface ActiveLoanDataType {
  serialNumber: string;
  loanAmount: number;
  amountLeftToRepay: number;
  duration: number;
  interestRate: number;
  installment: number;
  type: string;
  activeLoneStatus: string;
  userId: string;
}

export interface ActiveLoanResponseType {
  success: boolean;
  message: string;
  data: ActiveLoanDataType[];
}

export interface singleActiveLoanResponseType {
  success: boolean;
  message: string;
  data: ActiveLoanDataType;
}
