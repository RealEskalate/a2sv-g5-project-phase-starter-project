import React from "react";
import Card from "../components/loans/Card";
import ActiveLoans from '../components/loans/ActiveLoans';

const LoansPage = () => {
  return (
    <div>
      <div className="card-holder flex gap-16 px-10 py-4">
        <Card />
        <Card />
        <Card />
        <Card />
      </div>
      
      <ActiveLoans />
    </div>
  );
};

export default LoansPage;
