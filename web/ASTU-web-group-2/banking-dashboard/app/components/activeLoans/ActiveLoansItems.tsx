export type loan = {
  serialNumber: string | null;
  loanAmount: string | null;
  amountLeftToRepay: string | null;
  duration: number | null;
  interestRate: number | null;
  installment: string | null;
  type: string | null;
  activeLoneStatus: string | null;
  userId: string | null;
};

export const defaultloans: loan[] = [
  {
    serialNumber: "01",
    loanAmount: "100,000",
    amountLeftToRepay: "40,500",
    duration: 8,
    interestRate: 12,
    installment: "2,000",
    type: "",
    activeLoneStatus: "",
    userId: "",
  },

  {
    serialNumber: "02",
    loanAmount: "500,000",
    amountLeftToRepay: "250,000",
    duration: 36,
    interestRate: 10,
    installment: "8,000",
    type: "",
    activeLoneStatus: "",
    userId: "",
  },

  {
    serialNumber: "03",
    loanAmount: "900,000",
    amountLeftToRepay: "40,500",
    duration: 12,
    interestRate: 12,
    installment: "5,000",
    type: "",
    activeLoneStatus: "",
    userId: "",
  },

  {
    serialNumber: "04",
    loanAmount: "50,000",
    amountLeftToRepay: "40,500",
    duration: 25,
    interestRate: 5,
    installment: "2,000",
    type: "",
    activeLoneStatus: "",
    userId: "",
  },

  {
    serialNumber: "05",
    loanAmount: "50,000",
    amountLeftToRepay: "40,500",
    duration: 5,
    interestRate: 16,
    installment: "10,000",
    type: "",
    activeLoneStatus: "",
    userId: "",
  },

  {
    serialNumber: "06",
    loanAmount: "80,000",
    amountLeftToRepay: "25,500",
    duration: 14,
    interestRate: 8,
    installment: "2,000",
    type: "",
    activeLoneStatus: "",
    userId: "",
  },

  {
    serialNumber: "07",
    loanAmount: "12,000",
    amountLeftToRepay: "5,500",
    duration: 9,
    interestRate: 13,
    installment: "500",
    type: "",
    activeLoneStatus: "",
    userId: "",
  },

  {
    serialNumber: "08",
    loanAmount: "160,000",
    amountLeftToRepay: "100,800",
    duration: 3,
    interestRate: 12,
    installment: "900",
    type: "",
    activeLoneStatus: "",
    userId: "",
  },

  {
    serialNumber: "Total",
    loanAmount: "125,0000",
    amountLeftToRepay: "750,000",
    duration: null,
    interestRate: null,
    installment: "50,000",
    type: "",
    activeLoneStatus: "",
    userId: "",
  },
];
