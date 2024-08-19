export interface LoanType {
    serialNumber: string;
    loanAmount: number;
    amountLeftToRepay: number;
    duration: number;
    interestRate: number;
    installment: number;
    type: 'personal' | 'corporate';
    activeLoneStatus: 'pending' | string; // Add other possible statuses if known
    userId: string;
}

export interface ApiResponse {
    success: boolean;
    message: string;
    data: LoanType[];
}