import useNavigation from "@/hooks/use-navigation";
import React, { useEffect, useState } from "react";
import { BsJustify } from "react-icons/bs";
import { MdSettings, MdNotifications } from "react-icons/md";
import { IoPersonCircleSharp } from "react-icons/io5";
import { useRouter } from "next/navigation";
import { signOut, useSession } from "next-auth/react";

interface Props {
  openSidebar: () => void;
}

const NavBar: React.FC<Props> = ({ openSidebar }) => {
  const [currentPage, setCurrentPage] = useState<string>("");
  const { data: session } = useSession();
  const router = useRouter();

  console.log("session", session);

  const {
    isDashboardActive,
    isInvestmentsActive,
    isTransactionsActive,
    isAccountsActive,
    isLoansActive,
    isServicesActive,
    isSettingsActive,
    isPrevilegesActive,
    isCreditsActive,
  } = useNavigation();


  useEffect(() => {
    let page = "";
    switch (true) {
      case isDashboardActive:
        page = "Dashboard";
        break;
      case isInvestmentsActive:
        page = "Investments";
        break;
      case isTransactionsActive:
        page = "Transactions";
        break;
      case isAccountsActive:
        page = "Accounts";
        break;
      case isLoansActive:
        page = "Loans";
        break;
      case isServicesActive:
        page = "Services";
        break;
      case isSettingsActive:
        page = "Settings";
        break;
      case isPrevilegesActive:
        page = "Previleges";
        break;
      case isCreditsActive:
        page = "Credits";
        break;
      default:
        page = "Overview";
    }
    setCurrentPage(page);
  }, [
    isDashboardActive,
    isInvestmentsActive,
    isTransactionsActive,
    isAccountsActive,
    isLoansActive,
    isServicesActive,
    isSettingsActive,
    isPrevilegesActive,
    isCreditsActive,
  ]);

  const handleLogout = () => {
    localStorage.removeItem("token");
    signOut({ callbackUrl: "/auth/login" });
  };

  return (
    <div className="shadow-md flex justify-between items-center py-3 px-4 rounded-md bg-white sticky top-0 z-40">
      <p className="text-lg font-semibold">{currentPage}</p>
      <div className="flex items-center space-x-4">
        <MdSettings
          className="cursor-pointer text-[#2b37e0]"
          onClick={() => router.push("/settings")}
          aria-label="Settings-Button"
          size={20}
        />
        <MdNotifications
          className="cursor-pointer text-[#e02b2b]"
          onClick={() => alert("You have no new notifications.")}
          aria-label="Notification-Bell-Button"
          size={20}
        />
        {session ? (
          <div className="flex items-center space-x-2">
            <span className="font-medium">{session.user?.name}</span>
            <button
              className="bg-[#2b37e0] text-white rounded px-3 py-1"
              onClick={handleLogout}
              aria-label="Logout-Button"
            >
              Logout
            </button>
          </div>
        ) : (
          <IoPersonCircleSharp
            className="cursor-pointer"
            onClick={() => router.push("/auth/login")}
            aria-label="Profile-Avatar-Button"
            size={40}
          />
        )}
        <button className="md:hidden" onClick={openSidebar} aria-label="Open-Sidebar-Button">
          <BsJustify className="h-6 w-6" />
        </button>
      </div>
    </div>
  );
};

export default NavBar;
