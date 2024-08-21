import React from "react";
import Card from "../components/loans/Card";
import ActiveLoans from '../components/loans/ActiveLoans';

const LoansPage = () => {
  return (
    <div>
        <Card />
      <ActiveLoans />
    </div>
  );
};

export default LoansPage;
