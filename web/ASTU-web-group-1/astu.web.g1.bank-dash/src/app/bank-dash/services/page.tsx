import BankServicesList from "@/components/BankServicesList/BankServicesList";
import ServicesCardApp from "@/components/ServicesCards/servicesCardApp";
import StoreProvider from "@/providers/StoreProvider";
import React from "react";

export default function page() {
  return (
    <div>
      <ServicesCardApp />
      <StoreProvider>
        <BankServicesList />
      </StoreProvider>
    </div>
  );
}
