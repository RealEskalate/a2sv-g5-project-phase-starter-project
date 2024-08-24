import React from "react";
import Cards from "./Card";
import ActiveLoans from "./ActiveLoans";

const LoansPage: React.FC = () => {
  return (
    <div
      style={{
        backgroundColor: "#F5F7FA",
        minHeight: "100vh",
        padding: "20px",
      }}
    >
      <Cards />
      <ActiveLoans />
    </div>
  );
};

export default LoansPage;
