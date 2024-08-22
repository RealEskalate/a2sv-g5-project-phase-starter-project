import { MdHome, MdAttachMoney, MdPayment, MdSettings } from "react-icons/md";
import { RiHandCoinLine } from "react-icons/ri";

const pages = [
  {
    text: "Dashboard",
    destination: "/dashboard",
    icon: MdHome,
  },
  {
    text: "Transactions",
    destination: "/transactions",
    icon: MdAttachMoney,
  },
  {
    text: "Loans",
    destination: "/loans",
    icon: RiHandCoinLine,
  },
  {
    text: "Services",
    destination: "/services",
    icon: MdPayment,
  },
  {
    text: "Settings",
    destination: "/settings",
    icon: MdSettings,
  },
];

export { pages };
