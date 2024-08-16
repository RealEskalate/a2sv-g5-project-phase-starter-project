import React from 'react'
import InfoboxForServicePage from '../components/infobox/InfoboxForServicePage'
import BankServicesList from '../components/bankServicesList/BankServicesList'
import Card from '../components/card/Card'

const ServicesPage = () => {
  return (
    <div className="flex flex-col gap-2">
    <InfoboxForServicePage />
    <Card title="Active Loans Overview" className="flex flex-col max-sm:w-fit">
        <BankServicesList />
      </Card>
  </div>
  )
}

export default ServicesPage