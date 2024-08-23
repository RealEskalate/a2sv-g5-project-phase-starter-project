import { useRouter } from "next/navigation";
import React from "react";
import { IconType } from "react-icons";
import { IoLogOutOutline, IoMoonOutline, IoSettingsOutline } from "react-icons/io5";
import { useDarkMode } from "./Context/DarkModeContext";
import { signOut } from "next-auth/react";

type ElementType = {
  id: number;
  text: string;
  destination: string;
  icon: IconType;
};

interface Props {
  handleNav: (s: string) => void;
  handleActive: (s: string) => void;
  elements: ElementType[];
  active: string;
}

const SidebarElements = ({
  handleActive,
  handleNav,
  elements,
  active,
}: Props) => {
  const route = useRouter();
  const { darkMode, toggleDarkMode } = useDarkMode(); // Use dark mode context

  return (
    <div className="flex flex-col gap-5 mb-5 ">
      {elements.map((el) => (
        <div
          key={el.id}
          className={`${
            active === el.text ? "text-[#2D60FF] border-l-2" : "text-[#B1B1B1]"
          } flex gap-3 items-center font-semibold text-l`}
        >
          <button
            onClick={() => {
              handleActive(el.text);
              handleNav(el.destination);
            }}
            className={`flex items-center w-full`}
          >
            <span
              className={`${
                active === el.text ? "bg-[#2D60FF]" : "hidden"
              } rounded-r-lg w-1 h-10`}
            ></span>
            <div className="px-5 flex items-center gap-6">
              <span className="text-2xl">
                <el.icon />
              </span>
              {el.text.charAt(0).toUpperCase() + el.text.slice(1)}
            </div>
          </button>
        </div>
      ))}
      <div className="flex gap-5 text-xl md:items-center px-10 py-10 md:hidden">
        <div
          className="cursor-pointer text-xl bg-[#F5F7FA] rounded-full px-2 py-2 dark:bg-[#050914] dark:border dark:border-[#333B69]"
          onClick={() => route.push("./bankingSettings")}
        >
          <IoSettingsOutline />
        </div>
        {/* <div className="cursor-pointer text-xl bg-[#F5F7FA] rounded-full px-2 py-2 dark:bg-[#050914] dark:border dark:border-[#333B69]">
              <IoMdNotificationsOutline />
            </div> */}
        <div
          className="cursor-pointer text-xl bg-[#F5F7FA] rounded-full px-2 py-2 dark:bg-[#050914] dark:border dark:border-[#333B69]"
          onClick={toggleDarkMode}
        >
          <IoMoonOutline />
        </div>
        <div
          className="cursor-pointer text-xl bg-[#F5F7FA] rounded-lg px-2 py-2 dark:bg-[#050914] dark:border dark:border-[#333B69]"
          onClick={() => {
            signOut();
          }}
        >
          <IoLogOutOutline />
        </div>
      </div>
    </div>
  );
};

export default SidebarElements;
