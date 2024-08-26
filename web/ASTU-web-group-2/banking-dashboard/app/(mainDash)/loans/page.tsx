import React from 'react'
import InfoboxForLoans from '../../components/infobox/InfoboxForLoans'
import ActiveLoansOverview from '../../components/activeLoans/ActiveLoansOverview'
import Card from '../../components/virtualCards/card/Card'

const LoansPage = () => {
  return (
    <div className="grid grid-cols-1 gap-2 pb-5">
      <InfoboxForLoans />
      <Card title="Active Loans Overview" className="">
        <ActiveLoansOverview />
      </Card>
    </div>
  )
}

export default LoansPage