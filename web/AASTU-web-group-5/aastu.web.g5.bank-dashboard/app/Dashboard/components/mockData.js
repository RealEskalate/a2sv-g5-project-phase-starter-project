// mockData.ts
import Deposit from "@/public/assets/icons/deposit.jpg"
import Transfer from "@/public/assets/icons/transfer.jpg"
import Paypal from '@/public/assets/icons/paypal.jpg'



export const RecentTransaction = [
    {
      name: "Deposit from my Card",
      date: "28 January 2021",
      amount: "-$850",
      image: Deposit, // Replace with the correct path to your image
    },
    {
      name: "Payment for Invoice #1234",
      date: "15 February 2021",
      amount: "-$150",
      image: Transfer,
    },
    {
      name: "Salary Payment",
      date: "1 March 2021",
      amount: "+$3000",
      image: Paypal,
    },
    {
      name: "Refund from Vendor",
      date: "10 March 2021",
      amount: "+$50",
      image: Deposit,
    },
  ];
