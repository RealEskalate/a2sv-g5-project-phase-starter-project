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

export const creditcardstyles = [{
  iconwhite : "w-6 h-6",
  icongray : "w-6 h-6"

}]

export const colors = {
  blue: 'bg-[#1814F3]',
  white: 'bg -[#ffffff]',
  navbartext:' text-[#343C6A]',
  black: 'bg-[#000000]',
  textblack: 'text-[#232323]',
  textred: 'text-[#FF4B4A]',
  textgreen: 'text-[#41D4A8]',
  green :'bg-[#16DBCC]',
  red : 'bg-[#FF82AC]',
  textgray : 'text-[#718EBF]',
  gradientcard:'linear-gradient(to bottom, #4C49ED,#0A06F4)',
  gradientchart:'linear-gradient(to bottom, #2D60FF,#2D60FF)',
  textblue : "text-[#1814F3]",
  lightblue : "bg-gray-100"
};

export const logo = {
  icon : "/icons/logo.png",
}

export const textColors = {
  textWhite: 'text-white',
  textDimmed: 'text-[#718EBF]',
}