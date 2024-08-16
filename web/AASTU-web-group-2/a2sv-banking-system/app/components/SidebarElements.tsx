import React from 'react';
import { IconType } from 'react-icons';

type ElementType = {
  id: number;
  text: string;
  destination: string;
  icon: IconType;
};

interface Props {
  handleNav: (s: string) => void;
  handleActive: (s: string) => void
  elements: ElementType[];
  active: string;
}

const SidebarElements = ({handleActive, handleNav, elements, active }: Props) => {
  return (
    <div className="flex flex-col gap-5 mb-5">
      {elements.map((el) => (
        <div
          key={el.id}
          className={`${
            active === el.text
              ? "text-[#2D60FF] border-l-2"
              : "text-[#B1B1B1]"
          } flex gap-3 items-center font-semibold text-l`}
        >
          <button
            onClick={() => {handleActive(el.text)
              handleNav(el.destination)
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
    </div>
  );
};

export default SidebarElements;
