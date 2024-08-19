import React from "react";
import Infobox from "../components/infobox/Infobox";
import LastTransaction from "../components/lastTransaction/LastTransaction";
import Card from "../components/card/Card";
import CreditCard from "../components/creditCard/CreditCard";
import CardForCreditCards from "../components/card/CardForCreditCards";
import DebitCreditOverviewChart from "../components/charts/DebitCreditOverviewChart";
import InvoicesSent from "../components/InvoiceSend/InvoiceSent";

const AccountsPage = () => {
  return (
    <div className="flex flex-col gap-2  pb-5">
      <Infobox />
      <div className="flex max-sm:flex-col justify-between gap-8">
        <Card
          title="Last Transaction"
          className="flex flex-col md:h-[299px] h-[254] w-full"
        >
          <LastTransaction />
        </Card>
        <CardForCreditCards
          title="Credit Cards"
          button="See All"
          link="/credit-cards"
          className="flex flex-col lg:h-[300px] w-[350px]"
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
      <div className="flex max-sm:flex-col justify-between gap-8">
        <Card
          title="Debit & Credit Overview"
          className="flex flex-col lg:w-[730px] w-[350px]"
        >
          <DebitCreditOverviewChart />
        </Card>
        <Card
          title="Invoices Sent"
          className="flex flex-col lg:w-[350px] lg:h-[300px] max-md:w-[350px]"
        >
          <InvoicesSent />
        </Card>
      </div>
    </div>
  );
};

export default AccountsPage;
