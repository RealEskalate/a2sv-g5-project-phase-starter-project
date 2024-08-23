import React from 'react'
import Link from 'next/link'
const Sidebar = () => {
  return (
    <div className='flex flex-col gap-5 px-5'>
      <Link href= "./dashboard"> Dashboard</Link>
      <Link href= "./transactions"> Transactions</Link>
      <Link href= "./accounts"> Accounts</Link>
      <Link href= "./investments"> Investments</Link>
      <Link href= "./creditCards"> Credit Cards</Link>
      <Link href= "./loans"> Loans</Link>
      <Link href= "./bankingServices"> Services</Link>
      <Link href= "./privileges"> My Privileges</Link>
      <Link href= "./bankingSettings"> Settings</Link>
      
    </div>
  )
}

export default Sidebar
