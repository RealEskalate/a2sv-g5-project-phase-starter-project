import React from 'react'
import InfoboxForLoans from '../components/infobox/InfoboxForLoans'
import ActiveLoansOverview from '../components/activeLoans/ActiveLoansOverview'
import Card from '../components/card/Card'

const LoansPage = () => {
  return (
    <div className="flex flex-col gap-2">
      <InfoboxForLoans />
      <Card title="Active Loans Overview" className="flex flex-col max-sm:w-fit">
        <ActiveLoansOverview />
      </Card>
    </div>
  )
}

export default LoansPage