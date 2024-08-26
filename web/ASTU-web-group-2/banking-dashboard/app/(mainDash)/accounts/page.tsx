"use client";
import React, { useEffect, useState } from "react";
import Infobox from "../../components/infobox/Infobox";
import LastTransaction from "../../components/transactions/lastTransaction/LastTransaction";
import Card from "../../components/virtualCards/card/Card";
import CreditCard from "../../components/virtualCards/creditCard/CreditCard";
import CardForCreditCards from "../../components/virtualCards/card/CardForCreditCards";
import DebitCreditOverviewChart from "../../components/charts/DebitCreditOverviewChart";
import InvoicesSent from "../../components/InvoiceSend/InvoiceSent";
import CardDisplay, {
  CardData,
  formatCardNumber,
  card,
} from "@/app/components/virtualCards/cardDisplay/CardDisplay";

import {
  useGetAllCardInfoQuery,
  useRetiriveCardInfoQuery,
} from "@/lib/service/CardService";
import { useSession } from "next-auth/react";
import CardSkeleton from "@/app/components/virtualCards/creditCard/CardSkeleton";

// import { useRouter } from "next/navigation";

const AccountsPage = () => {
  return (
    <div className="grid grid-cols-1 gap-2 pb-5">
      <Infobox />
      <div className="flex max-sm:flex-col justify-between gap-8">
        <Card
          title="Last Transaction"
          className="flex flex-col md:h-[299px] h-[254] w-full h-[100%]"
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
