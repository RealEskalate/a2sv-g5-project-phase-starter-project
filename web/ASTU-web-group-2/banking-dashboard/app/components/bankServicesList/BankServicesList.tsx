import React from "react";
import BankService from "../bankService/BankService";
import BankServiceMobile from "../bankService/BankServiceMobile";

const BankServicesList = () => {
  
  return (
    <div>
      <div className="flex flex-col gap-5 max-md:hidden">
        {data.map((bankService) => (
          <BankService
            logoLink={bankService.logo}
            category={bankService.category}
            description={bankService.description}
            details={bankService.details}
            action={bankService.action}
            key={bankService.id}
          />
        ))}
      </div>

      {/* Mobile view */}
      <div className="flex flex-col gap-5 md:hidden">
        {data.map((bankService) => (
          <BankServiceMobile
            logoLink={bankService.logo}
            category={bankService.category}
            description={bankService.description}
            details={bankService.details}
            action={bankService.action}
            key={bankService.id}
          />
        ))}
      </div>
    </div>
  );
};

export default BankServicesList;
