"use client";
import Image from "next/image";
import Link from "next/link";
import { useEffect } from "react";
import { usePathname } from "next/navigation";
import { useSelector, useDispatch } from "react-redux";
import { RootState } from "@/lib/redux/store"; // Update the path if necessary
import { toggleSidebar, setActiveItem } from "@/lib/redux/slices/layoutSlice"; // Update the path if necessary
import { menuItems, logo } from "@/../../public/Icons";
import { signOut } from "next-auth/react";
import { AiOutlineLogout } from "react-icons/ai";

const Sidebar = () => {
  const dispatch = useDispatch();
  const pathname = usePathname();
  const { ishidden, activeItem } = useSelector(
    (state: RootState) => state.layout
  );
  const currentPath = pathname;

  useEffect(() => {
    const activeMenuItem = menuItems.find((item) =>
      currentPath.startsWith(item.href)
    );
    if (activeMenuItem) {
      dispatch(setActiveItem(activeMenuItem!.title));
    } else if (!activeMenuItem && currentPath === "/") {
      dispatch(setActiveItem("Dashboard"));
    }
  }, [currentPath, dispatch]);

  return (
    <aside
      className={`fixed top-0 left-0 z-40 min-w-44 w-44 h-full bg-white dark:bg-darkBackground shadow-md transition-transform transform ${
        ishidden ? "translate-x-0" : "-translate-x-full"
      } md:translate-x-0 md:relative md:w-[30%] lg:w-[20%] md:flex md:flex-col md:gap-6 md:py-1 md:border md:border-[#E6EFF5] dark:md:border-gray-700`}
    >
      <div className="ml-5 mt-3 flex items-center gap-8">
        <div className="flex justify-between items-center">
          <Image src={logo} alt="Logo" width={36} height={36} />
          <div className="text-[#343C6A] dark:text-darkText pl-2 md:text-xl md:pl-1 lg:pl-2 lg:text-2xl text-base xl:text-4xl md:text-[21px] font-[800] font-mont">
            BankDash.
          </div>
        </div>
        <div
          onClick={() => dispatch(toggleSidebar())}
          className="md:hidden rounded-full text-end py-3 px-4 text-blue-600 dark:text-darkAccent font-bold bg-[#F5F7FA] dark:bg-gray-800 dark:hover:bg-gray-600 "
        >
          X
        </div>
      </div>
      <div className="flex flex-col justify-between h-full">
        <ul className="flex flex-col mb-16">
          {menuItems.map((item) => (
            <Link href={item.href} key={item.title}>
              <li
                onClick={() => dispatch(setActiveItem(item.title))}
                className={`flex gap-3 items-center px-8 py-3 text-md md:text-lg ${
                  activeItem === item.title
                    ? "border-l-4 border-l-[#2D60FF] text-[#2D60FF] font-bold dark:border-l-darkAccent dark:text-darkAccent"
                    : "text-[#a59d9d] dark:text-darkText"
                }`}
              >
                <Image
                  src={item.icon}
                  alt={item.title}
                  width={24}
                  height={24}
                  className={`w-6 h-6 ${
                    activeItem === item.title
                      ? "filter-active"
                      : "filter-inactive"
                  }`}
                />
                <div>{item.title}</div>
              </li>
            </Link>
          ))}
        </ul>
        <button
          className="ml-5 w-1/2 mb-4 flex gap-3 items-center rounded-xl hover:text-red-600 text-xl font-semibold hover:bg-[#F5F7FA] dark:hover:bg-gray-700 p-2"
          onClick={() => {
            signOut({ callbackUrl: "/auth/signin" });
          }}
        >
          <AiOutlineLogout />
          Logout
        </button>
      </div>
    </aside>
  );
};

export default Sidebar;
