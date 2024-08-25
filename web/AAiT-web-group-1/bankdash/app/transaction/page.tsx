"use client";

import { useState, useEffect } from 'react';
import Cards from "../components/Common/Card";
import TransactionsTable from "../components/Transactions/transactions_table";
import ExpenseBarchart from "../components/Transactions/MyExpenseBarChart";
// card data
const card_data = [
  { balance: "$5,708", cardHolder: "Eddy Cusuma", cardNumber: "3778 **** **** 1234", validThru: "12/26" },
  { balance: "$908", cardHolder: "Jatani Iya", cardNumber: "3778 **** **** 5678", validThru: "11/25" },
  { balance: "$2,908", cardHolder: "Dida Jaldo", cardNumber: "3778 **** **** 9101", validThru: "02/30" },
  { balance: "$1,500", cardHolder: "Mikael Tesfaye", cardNumber: "3778 **** **** 1121", validThru: "05/27" },
  { balance: "$3,200", cardHolder: "Sara Yared", cardNumber: "3778 **** **** 3141", validThru: "08/28" },
  { balance: "$4,750", cardHolder: "Liya Alemu", cardNumber: "3778 **** **** 5161", validThru: "09/29" }
];

//Expense bar chart data
const expenses = [11200, 16200, 12400, 6400, 15000, 12000, 12100]
const months = ['Aug', 'Sep', 'Oct', 'Nov', 'Dec', 'Jan']
// transaction table data
const transactions = [
  {
    description: 'Spotify Subscription',
    transactionId: '#12548796',
    type: 'Shopping',
    card: '1234 ****',
    date: '28 Jan, 12.30 AM',
    amount: '-$9.99',
    amountColor: '',
  },
  {
    description: 'Amazon Purchase',
    transactionId: '#12548797',
    type: 'Shopping',
    card: '5678 ****',
    date: '27 Jan, 3.15 PM',
    amount: '-$45.00',
    amountColor: '',
  },
  {
    description: 'Salary Credit',
    transactionId: '#12548798',
    type: 'Transfer',
    card: '1234 ****',
    date: '25 Jan, 10.40 PM',
    amount: '+$3,000',
    amountColor: '',
  },
  {
    description: 'Electricity Bill',
    transactionId: '#12548799',
    type: 'Utilities',
    card: '5678 ****',
    date: '24 Jan, 8.00 AM',
    amount: '-$120.50',
    amountColor: '',
  },
  {
    description: 'Freelance Payment',
    transactionId: '#12548800',
    type: 'Transfer',
    card: '1234 ****',
    date: '23 Jan, 5.30 PM',
    amount: '+$500',
    amountColor: '',
  },
  {
    description: 'Netflix Subscription',
    transactionId: '#12548801',
    type: 'Shopping',
    card: '5678 ****',
    date: '22 Jan, 9.00 PM',
    amount: '-$15.99',
    amountColor: '',
  },
  {
    description: 'Gym Membership',
    transactionId: '#12548802',
    type: 'Shopping',
    card: '1234 ****',
    date: '21 Jan, 7.00 AM',
    amount: '-$50.00',
    amountColor: '',
  },
  {
    description: 'Grocery Shopping',
    transactionId: '#12548803',
    type: 'Shopping',
    card: '5678 ****',
    date: '20 Jan, 4.00 PM',
    amount: '-$200.75',
    amountColor: '',
  },
  {
    description: 'Book Sale',
    transactionId: '#12548804',
    type: 'Transfer',
    card: '1234 ****',
    date: '19 Jan, 11.30 AM',
    amount: '+$30.00',
    amountColor: '',
  },
  {
    description: 'Coffee Shop',
    transactionId: '#12548805',
    type: 'Shopping',
    card: '5678 ****',
    date: '18 Jan, 2.00 PM',
    amount: '-$5.50',
    amountColor: '',
  },
  {
    description: 'Spotify Subscription',
    transactionId: '#12548806',
    type: 'Shopping',
    card: '1234 ****',
    date: '17 Jan, 12.30 AM',
    amount: '-$9.99',
    amountColor: '',
  },
  {
    description: 'Freelance Payment',
    transactionId: '#12548807',
    type: 'Transfer',
    card: '5678 ****',
    date: '16 Jan, 10.40 PM',
    amount: '+$750',
    amountColor: '',
  },
  {
    description: 'Electricity Bill',
    transactionId: '#12548808',
    type: 'Utilities',
    card: '1234 ****',
    date: '15 Jan, 8.00 AM',
    amount: '-$120.50',
    amountColor: '',
  },
  {
    description: 'Amazon Purchase',
    transactionId: '#12548809',
    type: 'Shopping',
    card: '5678 ****',
    date: '14 Jan, 3.15 PM',
    amount: '-$45.00',
    amountColor: '',
  },
  {
    description: 'Salary Credit',
    transactionId: '#12548810',
    type: 'Transfer',
    card: '1234 ****',
    date: '13 Jan, 10.40 PM',
    amount: '+$3,000',
    amountColor: '',
  },
  {
    description: 'Spotify Subscription',
    transactionId: '#12548796',
    type: 'Shopping',
    card: '1234 ****',
    date: '28 Jan, 12.30 AM',
    amount: '-$9.99',
    amountColor: '',
  },
  {
    description: 'Amazon Purchase',
    transactionId: '#12548797',
    type: 'Shopping',
    card: '5678 ****',
    date: '27 Jan, 3.15 PM',
    amount: '-$45.00',
    amountColor: '',
  },
  // Add more transactions as needed
];


function whichIndex(){
  return Math.floor(Math.random() * card_data.length);
}
const Page: React.FC = () => {
  const [index_arr, setIndex] = useState<number[]>([0, 1])

  // Generate random index
  useEffect(() => {
    let f_index: number = whichIndex();
    let s_index: number = whichIndex();
    while (f_index === s_index) {
        s_index = whichIndex();
    }
    setIndex([f_index, s_index]);
  }, [])


return (
    <div className="mt-[50px] mr-5">
        <div className="bg-[#F5F7FA] flex">

        {/*...... cards part ......*/}
        <div className="w-[35%]">
            <p className="font-inter text-[21px] font-semibold leading-[21.78px] text-left text-[#343C6A] mb-5">My Cards</p>
            <Cards
            cData1={card_data[index_arr[0]]}
            cData2={card_data[index_arr[1]]}
            bgCol="bg-custom-gradient"
            bbgCol="bg-bottom-gradient"
            isBlue={true} //this is for card-color
            imageSrc="./images/Chip_Card (1).png"
            iconSrc="/images/Group 17 (1).png"
            />
        </div>

        <div className="text-right ml-[20px] w-[35%]">
            <p className="font-inter text-[17px] font-semibold leading-[19.15px] text-[#343C6A] mb-5">+ Add Card</p>
            <Cards
            cData2={card_data[index_arr[1]]}
            cData1={card_data[index_arr[0]]}
            imageSrc="./images/Chip_Card.png"
            iconSrc="/images/Group 17.png"
            />
        </div>

        {/*...... Expense-Barchart part ......*/}
        <div className="ml-[20px] w-[35%]">
            <p className="font-inter text-[21px] font-semibold leading-[21.78px] text-left text-[#343C6A] mb-5">My Expense</p>
            <ExpenseBarchart Expenses={expenses} months={months}/>
        </div>
        </div>

        {/* table part */}
        <div>
        <TransactionsTable transactions={transactions} />
        </div>
    </div>
);
    };

export default Page;