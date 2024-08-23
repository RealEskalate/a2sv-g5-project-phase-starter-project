import React from 'react';
import Balance from '@/public/accountsimage/balance.svg';
import Income from '@/public/accountsimage/income.svg';
import Expense from '@/public/accountsimage/expense.svg';
import Saving from '@/public/accountsimage/saving.svg'
import Card from '@/components/accounts/card';
import BalanceCard from '@/components/commonalities/BalanceCard';
import spotify from '@/public/accountsimage/spotify.svg';
import service from '@/public/accountsimage/service.svg';
import person from '@/public/accountsimage/person.svg'
import LastTransactionsComp from '@/components/accounts/lastTransactions';
import AccountsBarChartComponent from '@/components/accounts/AccountsBarChart';


const page = () => {
  const Cards = [
    {
      image: Balance,
      title: "My Balance",
      value: "12,750"
    },
    {
      image: Income,
      title: "Income",
      value: "5,600"
    },
    {
      image: Expense,
      title: "Expense",
      value: "3,460"
    },
    {
      image: Saving,
      title:"Total Saving",
      value:"7,920"
    }
  ]

  const LastTransactions = [
    {
      image: spotify,
      title: "Spotify Subscription",
      date: "25 Jan 2021",
      type: "Shopping",
      phone: "1234 ****",
      status: "Pending",
      amount: "-150"
    },
    {
      image: service,
      title: "Mobile Service",
      date: "25 Jan 2022",
      type: "Service",
      phone: "1234 ****",
      status: "Completed",
      amount: "-340"
    },
    {
      image: person,
      title: "Emily Wilson",
      date: "20 Aug 2022",
      type: "Transfer",
      phone: "1234 ****",
      status: "Completed",
      amount: "780"
    },
  ]
  return (
    <div>
      <div className="flex items-center justify-between mx-auto p-6">
        {
          Cards.map((item, index) => {
              return <Card key={index} title={item.title} value={item.value} image={item.image} />;
          })
        }
      </div>
      <div className="grid grid-cols-12 gap-4 mt-5  md:mt-10">
        <div className="col-span-7 space-y-3 px-6">
          <h2 className='font-semibold'>Last Transaction</h2>
          <div className="flex flex-col bg-white rounded-xl">
          {
            LastTransactions.map((item, index) => {
              return <LastTransactionsComp key={index} item={item}  />
            })
          }
          </div>
        </div>
        <div className="col-span-5 mx-auto space-y-3">
          <span className='flex items-center space-x-40'>
            <p>My Card</p>
            <button>See All</button>
          </span>
          <BalanceCard property='blue'/>
        </div>
      </div>
      <div className="grid grid-cols-12 gap-4 px-6 mt-10">
        <div className="col-span-7 space-y-5">
          <h1 className='font-bold'>Debit & Credit Overview</h1>
          <AccountsBarChartComponent/>
        </div>
        <div className="col-span-5 border border-black">
          <h1 className='font-bold'>Invoices Sent</h1>
          <></>
        </div>
      </div>
    </div>
  )
}

export default page
