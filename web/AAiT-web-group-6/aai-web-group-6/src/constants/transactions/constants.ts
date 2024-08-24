import { TransactionTableProps } from "@/constants/transactions/types";

export const CredidtCardsDemo: Props[] = [
  {
    balance: 1000,
    cardHolder: "John Doe",
    validTHRU: "12/23",
    Code: "3778 **** **** 1234",
    isWhite: false,
  },
  {
    balance: 2000,
    cardHolder: "Sona Hanks",
    validTHRU: "10/22",
    Code: "3778 **** **** 5678",
    isWhite: true,
  },
];

type Props = {
  balance: Number;
  cardHolder: string;
  validTHRU: string;
  Code: string;
  isWhite: boolean;
};

export const transactionTableData: TransactionTableProps[] = [
  {
    description: "Spotify Subscription",
    transactionId: "#12548796",
    type: "Shopping",
    card: "1234 ****",
    date: "28 Jan, 12.30 AM",
    amount: -2500,
  },
  {
    description: "Freepik Sales",
    transactionId: "#12548796",
    type: "Transfer",
    card: "1234 ****",
    date: "25 Jan, 10.40 PM",
    amount: +750,
  },
  {
    description: "Mobile Service",
    transactionId: "#12548796",
    type: "Service",
    card: "1234 ****",
    date: "20 Jan, 10.40 PM",
    amount: -150,
  },
  {
    description: "Wilson",
    transactionId: "#12548796",
    type: "Transfer",
    card: "1234 ****",
    date: "15 Jan, 03.29 PM",
    amount: -1050,
  },
  {
    description: "Emilly",
    transactionId: "#12548796",
    type: "Transfer",
    card: "1234 ****",
    date: "14 Jan, 10.40 PM",
    amount: +840,
  },
];
