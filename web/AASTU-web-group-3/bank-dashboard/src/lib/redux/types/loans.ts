export interface Loan {
    success: boolean,
    message: string,
    data: {
      content: [
        {
          serialNumber: string,
          loanAmount: number,
          amountLeftToRepay: number,
          duration: number,
          interestRate: number,
          installment: number,
          type: string,
          activeLoneStatus: string,
          userId: string
        }
      ],
      totalPages: number
    }
  }