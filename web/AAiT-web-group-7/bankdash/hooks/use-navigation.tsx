"use client";

import { useEffect, useState } from "react";

import { usePathname } from "next/navigation";

const useNavigation = (): any => {
  const pathname = usePathname();
  const [isDashboardActive, setIsDashboardActive] = useState<boolean>(false);
  const [isInvestmentsActive, setIsInvestmentsActive] = useState<boolean>(false);
const [isTransactionsActive, setIsTransactionsActive] = useState<boolean>(false);
const [isAccountsActive, setIsAccountsActive] = useState<boolean>(false);
const [isLoansActive, setIsLoansActive] = useState<boolean>(false);
const [isServicesActive, setIsServicesActive] = useState<boolean>(false);
const [isSettingsActive, setIsSettingsActive] = useState<boolean>(false);
const [isPrevilegesActive, setIsPrevilegesActive] = useState<boolean>(false);
const [isCreditsActive, setIsCreditsActive] = useState<boolean>(false);


  useEffect(() => {
    setIsDashboardActive(false);
    setIsInvestmentsActive(false);
    setIsTransactionsActive(false);
    setIsAccountsActive(false);
    setIsLoansActive(false);
    setIsServicesActive(false);
    setIsSettingsActive(false);
    setIsPrevilegesActive(false);
    setIsCreditsActive(false);


    if (pathname === "/dashboard") {
      setIsDashboardActive(true);
    }else if (pathname === "/investments") {
      setIsInvestmentsActive(true);
    }
else if (pathname === "/transactions") {
      setIsTransactionsActive(true);
    }
else if (pathname === "/accounts") {
      setIsAccountsActive(true);
    }
else if (pathname === "/loans") {
      setIsLoansActive(true);
    }
else if (pathname === "/services") {
      setIsServicesActive(true);
    }
else if (pathname === "/settings") {
      setIsSettingsActive(true);
    }
else if (pathname === "/previleges") {
      setIsPrevilegesActive(true);
    }
else if (pathname === "/credits") {
      setIsCreditsActive(true);
    }

  }, [pathname]);

  return {
    isDashboardActive,
    isInvestmentsActive,
    isTransactionsActive,
    isAccountsActive,
    isLoansActive,
    isServicesActive,
    isSettingsActive,
    isPrevilegesActive,
    isCreditsActive,
  };
};

export default useNavigation;
