import {
  Homeicon,
  Transfericon,
  Investicon,
  CreditCardicon,
  Loanicon,
  Serviceicon,
  Settingsicon,
} from "@/components/dashboard-layout/icons";

export const sidebarLinks = [
  {
    title: "Dashboard",
    route: "/dashboard",
    icon: Homeicon,
  },
  {
    title: "Transactions",
    route: "/dashboard/transactions",
    icon: Transfericon,
  },
  {
    title: "Accounts",
    route: "/dashboard/accounts",
    icon: Investicon,
  },

  {
    title: "Investments",
    route: "/dashboard/investments",
    icon: Investicon,
  },
  {
    title: "Credit Cards",
    route: "/dashboard/credit-cards",
    icon: CreditCardicon,
  },
  {
    title: "Loans",
    route: "/dashboard/loans",
    icon: Loanicon,
  },
  {
    title: "Services",
    route: "/dashboard/services",
    icon: Serviceicon,
  },
  {
    title: "Setting",
    route: "/dashboard/settings",
    icon: Settingsicon,
  },
];

export const HeadersTitle = (path: string) => {
  switch (path) {
    case "/dashboard":
      return "Dashboard";
    case "/dashboard/transactions":
      return "Transactions";
    case "/dashboard/accounts":
      return "Accounts";
    case "/dashboard/investments":
      return "Investments";
    case "/dashboard/credit-cards":
      return "Credit Cards";
    case "/dashboard/loans":
      return "Loans";
    case "/dashboard/services":
      return "Services";
    case "/dashboard/setting":
      return "Setting";
    default:
      return "Dashboard";
  }
};
