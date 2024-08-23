"use client";
import React, { useEffect, useState } from "react";
import Infobox from "../../components/infobox/Infobox";
import LastTransaction from "../../components/lastTransaction/LastTransaction";
import Card from "../../components/card/Card";
import CreditCard from "../../components/creditCard/CreditCard";
import CardForCreditCards from "../../components/card/CardForCreditCards";
import DebitCreditOverviewChart from "../../components/charts/DebitCreditOverviewChart";
import InvoicesSent from "../../components/InvoiceSend/InvoiceSent";
import CardDisplay, {
  CardData,
  formatCardNumber,
  card,
} from "@/app/components/cardDisplay/CardDisplay";

const AccountsPage = () => {
  return (
    <div className="grid grid-cols-1 gap-2 pb-5">
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
          <CardDisplay numofcard={1} />
        </CardForCreditCards>
      </div>
      <div className="grid lg:grid-cols-[3fr_2fr] max-md:grid-cols-1 gap-7 p-4 w-auto">
        <Card title="Debit & Credit Overview" className="w-full">
          <DebitCreditOverviewChart />
        </Card>
        <Card title="Invoices Sent" className="w-full">
          <InvoicesSent />
        </Card>
      </div>
    </div>
  );
};

export default AccountsPage;
