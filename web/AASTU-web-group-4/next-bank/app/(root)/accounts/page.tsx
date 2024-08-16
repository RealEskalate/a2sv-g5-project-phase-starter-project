'use client'
import React from 'react'
import BalanceCard from '@/components/AccountSmallCard'
import App from '@/components/LastTransactionCard'
import DesktopCreditCart from '@/components/DesktopCreditCard';
import BarChart from '@/components/BarChartAccount';

const Accounts: React.FC = () => {
  const labels = ["January", "February", "March", "April", "May", "June", "July"];
  const data = [65, 59, 80, 81, 56, 55, 40];
  const title = "Sales Over Months";


// const Accounts = () => {
  return (
    <div>
      <div>
        <h1>Accounts</h1>
        <BalanceCard/>
      </div>
      <div className='flex flex-col text-center justify-center'>
        <h1>Last Transaction</h1>
        <App/>
      </div>
        <div>
          <h1>My Card</h1>
          <DesktopCreditCart/>
          <DesktopCreditCart/>
          <DesktopCreditCart/>
          <DesktopCreditCart/>
        </div>
        {/* <div>
          <h1>Debit & Credit Overview</h1>
          <BarChart labels={labels} title={title}/>
        </div> */}
    </div>
  )
}

export default Accounts