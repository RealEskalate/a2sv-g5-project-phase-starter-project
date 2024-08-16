'use client'
import React from 'react'
import BalanceCard from '@/components/AccountSmallCard'
import App from '@/components/LastTransactionCard'
import DesktopCreditCart from '@/components/DesktopCreditCard';

const Accounts = () => {
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
        {/* <div>
          <h1>My Card</h1>
          <DesktopCreditCart/>
        </div> */}
        {/* <div>
          <h1>Debit & Credit Overview</h1>   
        </div> */}
    </div>
  )
}

export default Accounts