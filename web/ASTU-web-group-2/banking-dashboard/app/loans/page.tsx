import React from 'react'
import InfoboxForLoans from '../components/infobox/InfoboxForLoans'

const LoansPage = () => {
  return (
    <div className="flex flex-col gap-2">
      <InfoboxForLoans />
      {/* <div className="flex max-sm:flex-col gap-[30px]">
        <Card
          title="My Expense"
          className="flex flex-col lg:w-[730px] lg:h-[300px] md:w-[487px] md:h-[299px] h-[254]"
        >
          <LastTransaction />
        </Card>
        <div className="flex max-sm:flex-col gap-[30px]">
          <CardForCreditCards
            title="Credit Cards"
            button="See All"
            link="/credit-cards"
            className="fflex flex-col"
          >
              <CreditCard
                balance={1250}
                cardHolder="John Doe"
                expiryDate="12/24"
                cardNumber="1234 5678 9012 3456"
                cardType="primary" // Can be "primary", "secondary", or "tertiary"
              />
          </CardForCreditCards>
        </div>
      </div>
      <div className="flex max-sm:flex-col gap-[30px]">
        <Card
          title="Debit & Credit Overview"
          className="flex flex-col lg:w-[730px] lg:h-[300px] md:w-[487px] md:h-[299px] h-[254]"
        >
          <DebitCreditOverviewChart />
        </Card>
        <Card
          title="Invoices Sent"
          className="flex flex-col max-w-[350px] lg:mx-auto h-auto"
        >
          <InvoicesSent />
        </Card>
      </div> */}
    </div>
  )
}

export default LoansPage