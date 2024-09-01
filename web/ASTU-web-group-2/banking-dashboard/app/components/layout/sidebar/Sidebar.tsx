import React, { useState, useEffect } from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";

import { sidebarListItems } from "./sidebarListItems";
import { signOut } from "next-auth/react";
import Image from "next/image";

interface SidebarProps {
  show: boolean;
  setter: React.Dispatch<React.SetStateAction<boolean>>;
}

const Sidebar: React.FC<SidebarProps> = ({ show, setter }) => {
  //   const router = useRouter();
  const pathname = usePathname();
  const [activeIndex, setActiveIndex] = useState(0);

  useEffect(() => {
    const index = sidebarListItems.findIndex((item) => item.path === pathname);
    setActiveIndex(index !== -1 ? index : 0);
  }, [pathname]);

  // Define our base class
  const className =
    "bg-white border-r border-[#E6EFF5] w-[250px] transition-[margin-left] ease-in-out duration-500  fixed top-0 bottom-0 left-0 z-50";
  // Append class based on state of sidebar visibility
  const appendClass = show ? " ml-0 drop-shadow-md" : " ml-[-250px] md:ml-0";

  // Clickable menu items
  const MenuItem: React.FC<{
    icon: React.ReactNode;
    name: string;
    route: string;
  }> = ({ icon, name, route }) => {
    // Highlight menu item based on currently displayed route
    const colorClass =
      pathname === route ? "text-white" : "text-white/50 hover:text-white";

    return (
      <Link
        href={route}
        onClick={() => setter((prev) => !prev)}
        className={`flex gap-1 [&>*]:my-auto text-md pl-6 py-3 border-b-[1px] border-b-white/10 ${colorClass}`}
      >
        <div className="text-xl flex [&>*]:mx-auto w-[30px]">{icon}</div>
        <div>{name}</div>
      </Link>
    );
  };

  // Overlay to prevent clicks in background, also serves as our close button
  const ModalOverlay: React.FC = () => (
    <div
      className={`flex md:hidden fixed top-0 right-0 bottom-0 left-0 bg-black/40 z-30`}
      onClick={() => setter((prev) => !prev)}
    />
  );

  return (
    <>
      <div className={`${className}${appendClass}`}>
        <div className="flex items-center p-4">
          <img
            src={"/assets/navbar/credit-card.svg"}
            width={36}
            height={36}
            alt="bankDash logo"
            className="mr-3"
          />
          <p className="font-black text-[25px] text-[#343C6A]">BankDash.</p>
        </div>
        <div className="relative flex flex-col flex-1 overflow-y-auto pt-2">
          {/* Active indicator */}
          <div
            className="absolute left-0 h-[60px] w-[6px] bg-[#2D60FF] rounded-r-[10px] transition-transform duration-300"
            style={{
              transform: `translateY(${activeIndex * 60}px)`,
            }}
          ></div>
          <div className="flex flex-col flex-1">
            {sidebarListItems.map((item, index) => (
              <Link
                data-id={`side-link-${item.path.slice(1)}`}
                key={index}
                href={item.path}
                onClick={() => {
                  if (show) {
                    setter((prev) => !prev);
                  }
                }}
              >
                <div
                  className={`flex items-center gap-8 h-[60px] ${
                    pathname === item.path ? "text-[#2D60FF]" : "text-[#B1B1B1]"
                  }`}
                >
                  <div className="flex gap-8 items-center pl-6">
                    <img
                      data-id={`side-image-${item.path.slice(1)}`}
                      src={pathname === item.path ? item.activeIcon : item.icon}
                      alt={item.name}
                      width={20}
                      height={20}
                    />
                    <h1
                      className={`font-medium text-[16px] ${
                        pathname === item.path
                          ? "text-[#2D60FF]"
                          : "text-[#B1B1B1]"
                      }`}
                    >
                      {item.name}
                    </h1>
                  </div>
                </div>
              </Link>
            ))}
           
          </div>
        </div>
      </div>
      {show ? <ModalOverlay /> : null}
    </>
  );
};

export default Sidebar;
