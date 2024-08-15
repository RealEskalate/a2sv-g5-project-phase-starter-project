import {
  FaHome,
  FaExchangeAlt,
  FaWallet,
  FaChartLine,
  FaCreditCard,
  FaMoneyCheckAlt,
  FaCog,
  FaGift,
  FaUserShield,
} from "react-icons/fa";

export const sidebarLinks = [
  { id:1, route: "/", label: "Dashboard", Icon: FaHome },
  { id:2, route: "transaction", label: "Transaction", Icon: FaExchangeAlt },
  { id:3, route: "accounts", label: "Accounts", Icon: FaWallet },
  { id:4, route: "investments", label: "Investments", Icon: FaChartLine },
  { id:4, route: "credit-card", label: "Credit Card", Icon: FaCreditCard },
  { id:6, route: "loans", label: "Loans", Icon: FaMoneyCheckAlt },
  { id:7, route: "services", label: "Services", Icon: FaCog },
  { id:8, route: "transfer", label: "Transfer", Icon: FaGift },
  { id:9, route: "settings", label: "Settings", Icon: FaUserShield },
];

export const user = {
  name: "John Doe",
  email: "john.doe@example.com",
  profileImage: "/path/to/profile-image.jpg", // Use a valid path for the profile image
};
