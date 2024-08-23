import CardS from '@/components/CreditCards/CardS'
import RecentTransactionTable from '@/components/RecentTable/RecentTransactionTable'
import Top from '@/components/Top'
import React from 'react'

function Transaction() {
  return (
    <div>
        <Top topicName='Transaction'/>
        <CardS />
        <RecentTransactionTable />
    </div>
  )
}

export default Transaction