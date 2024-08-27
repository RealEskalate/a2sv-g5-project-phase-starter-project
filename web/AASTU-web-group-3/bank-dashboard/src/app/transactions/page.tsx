import React from "react";
import AllTransactions from "../components/transactions/AllTransactions";
import CreditCard from "../components/CreditCard";
import DashboardBarChart from "../components/Chart/DashboardBarChart";
import RecentTransactions from '../components/transactions/RecentTransactions';
import CreditCardTransaction from '../components/transactions/CreditCardTransaction';

const TransactionsPage = () => {
  return (
    <section className="xl:w-11/12 xl:mx-7">
      <div className="credit-cards expenses flex flex-col lg:flex-row justify-between lg:gap-1 xl:gap-6 ">
      
          <CreditCardTransaction/>
          <DashboardBarChart />
      </div>
      <RecentTransactions/>

    </section>
  );
};

export default TransactionsPage;
