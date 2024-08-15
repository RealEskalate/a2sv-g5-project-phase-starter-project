import React from 'react'
import BalanceCard from '@/components/AccountSmallCard'
import App from '@/components/LastTransactionCard'

const Accounts = () => {
  return (
    <div>
      <div>
        <h1>Accounts</h1>
        <div><BalanceCard/></div>
      </div>
      <div>
        <h1>Last Transaction</h1>
        <div><App/></div>
      </div>
    </div>
  )
}

export default Accounts