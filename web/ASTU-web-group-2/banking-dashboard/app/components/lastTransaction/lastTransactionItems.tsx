export type Item = {
  transactionId: string;
  type: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string;
};

export const defaultItems: Item[] = [
  {
    receiverUserName: "Spotify Subscription",
    date: "25 Jan 2021",
    type: "Shopping",
    transactionId: "1234 ****",
    description: "Pending",
    amount: -159,
  },

  {
    receiverUserName: "Mobile Service",
    date: "25 Jan 2021",
    type: "Service",
    transactionId: "1234 ****",
    description: "Completed",
    amount: -340,
  },

  {
    receiverUserName: "Emily Watson",
    date: "25 Jan 2021",
    type: "Transfer",
    transactionId: "1234 ****",
    description: "Completed",
    amount: 780,
  },
];
