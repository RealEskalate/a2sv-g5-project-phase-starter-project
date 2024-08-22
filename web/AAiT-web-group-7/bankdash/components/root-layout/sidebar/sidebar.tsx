import Image from "next/image";
import { HiX } from "react-icons/hi";
import { FiChevronsLeft } from "react-icons/fi";
import useNavigation from "@/hooks/use-navigation";
import SidebarLink from "./sidebaritems/sidebarLinks";
import logo from '@/public/banklogo.svg';

import HomeLogo from '@/public/sidebarlogo/home.svg'
import TransactionsLogo from '@/public/sidebarlogo/transactions.svg'
import AccountsLogo from '@/public/sidebarlogo/accounts.svg'
import InvestmentsLogo from '@/public/sidebarlogo/investment.svg'
import CreditCardsLogo from '@/public/sidebarlogo/credit.svg'
import LoansLogo from '@/public/sidebarlogo/loan.svg'
import ServicesLogo from '@/public/sidebarlogo/services.svg'
import MyPrevilegesLogo from '@/public/sidebarlogo/previleges.svg'
import SettingsLogo from '@/public/sidebarlogo/settings.svg'

interface Props {
  open: boolean;
  onClose: React.MouseEventHandler<HTMLSpanElement>;
  collapse: boolean;
  onCollapse: React.MouseEventHandler<HTMLSpanElement>;
}
interface ILink {
  title: string;
  icon: string;
  link: string;
  isActive: boolean;
  collapsed: boolean;
}

const Sidebar: React.FC<Props> = ({ open, onClose, collapse, onCollapse }) => {
  const {
    isDashboardActive,
    isLoansActive,
    isServicesActive,
    isAccountsActive,
    isInvestmentsActive,
    isTransactionsActive,
    isSettingsActive,
    isPrevilegesActive,
    isCreditsActive,
  } = useNavigation();

  const links: ILink[] = [
    {
      title: "Dashboard",
      icon: HomeLogo,
      link: "/dashboard",
      isActive: isDashboardActive,
      collapsed: collapse,
    },
    {
      title: "Transactions",
      icon: TransactionsLogo,
      link: "/transactions",
      isActive: isTransactionsActive,
      collapsed: collapse,
    },
    {
      title: "Accounts",
      icon: AccountsLogo,
      link: "/accounts",
      isActive: isAccountsActive,
      collapsed: collapse,
    },
    {
      title: "Investments",
      icon: InvestmentsLogo,
      link: "/investments",
      isActive: isInvestmentsActive,
      collapsed: collapse,
    },
    {
      title: "Credit Cards",
      icon: CreditCardsLogo,
      link: "/credits",
      isActive: isCreditsActive,
      collapsed: collapse,
    },
    {
      title: "Loans",
      icon: LoansLogo,
      link: "/loans",
      isActive: isLoansActive,
      collapsed: collapse,
    },
    {
        title: "Services",
        icon: ServicesLogo,
        link: "/services",
        isActive: isServicesActive,
        collapsed: collapse,
      },
      {
        title: "My Previleges",
        icon: MyPrevilegesLogo,
        link: "/previleges",
        isActive: isPrevilegesActive,
        collapsed: collapse,
      },
      {
        title: "Settings",
        icon: SettingsLogo,
        link: "/settings",
        isActive: isSettingsActive,
        collapsed: collapse,
      },
  ];

  return (
    <div
      className={`fixed top-0 border-r-2 left-0 z-50 flex min-h-full flex-col bg-surfContainer pb-10 shadow-white/5 transition-all duration-175 ${
        open ? "translate-x-0" : "-translate-x-full"
      } ${collapse ? "w-[90px]" : "w-[260px]"} ${
        open ? "bg-opacity-100 backdrop-blur-md" : "bg-opacity-0"
      }`}
      style={{ zIndex: 9999 }}  // Ensures it stays on top
    >
      <span
        className="absolute top-5 right-5 block cursor-pointer md:hidden"
        onClick={onClose}
        aria-label="Sidebar-Close-Icon"
      >
        <HiX className="h-[30px] w-[30px]" size={20} />
      </span>

      <div className={` h-[90px] w-full flex items-center gap-3 px-6 overflow-hidden`}>
        <Image src={logo} alt= "Logo" />
        <div
          className={`text-[#343c6a] text-2xl font-bold text-navy-700 ${collapse ? "opacity-0" : "opacity-100"}`}
        >
          BankDash
        </div>
      </div>

      {links.map((link: ILink, index: number) => {
        return (
          <SidebarLink
            title={link.title}
            Icon={link.icon}
            link={link.link}
            isActive={link.isActive}
            collapsed={link.collapsed}
            key={index}
          />
        );
      })}

      <span
        className="my-5 mx-auto cursor-pointer hidden md:block"
        onClick={onCollapse}
        aria-label="Collapse-Icon"
      >
        <FiChevronsLeft className={`h-[40px] w-[40px] ${collapse ? "rotate-180" : ""}`} />
      </span>
    </div>
  );
};

export default Sidebar;
