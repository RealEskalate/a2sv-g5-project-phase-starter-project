"use client";
import { useUser } from "@/contexts/UserContext";
import { getQuickTransfer } from "@/lib/api";
import { QuickTransferData } from "@/types";
import React, { useEffect, useRef, useState } from "react";
import {
  IoChevronBackCircleOutline,
  IoChevronForwardCircleOutline,
} from "react-icons/io5";
import { Profile } from "./Profile";
import { ModalTrans } from "./ModalTrans";
import { PiTelegramLogoLight } from "react-icons/pi";

// Shimmer component for skeleton loading effect
const Shimmer = () => {
  return (
    <div className="animate-pulse h-20 w-20 bg-gray-200 rounded-full"></div>
  );
};

export const QuickTransfer = ({ onLoadingComplete }: { onLoadingComplete: any }) => {
  const { isDarkMode } = useUser();
  const QuickTransferSection = useRef<HTMLDivElement | null>(null);
  const [quickTransfer, setQuickTransfer] = useState<QuickTransferData[]>([]);
  const [selectedProfile, setSelectedProfile] = useState<QuickTransferData>(
    {} as QuickTransferData
  );
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const accounts = await getQuickTransfer(10);
        setQuickTransfer(accounts || []);
          onLoadingComplete(false);
           setLoading(false);
      } finally {
       
      
      }
    };
    fetchData();
  }, [onLoadingComplete]);

  const handleProfileSelect = (account: QuickTransferData) => {
    setSelectedProfile(account);
  };

  const handleModalToggle = () => {
    setIsModalOpen(!isModalOpen);
  };

  const scrollCards = (scrollOffset: number) => {
    if (QuickTransferSection.current) {
      QuickTransferSection.current.scrollLeft += scrollOffset;
    }
  };

  return (
    <div className="space-y-5">
      <div className="font-inter text-[16px] font-semibold">
        <h4>Quick Transfer</h4>
      </div>
      <div
        className={`
        ${isDarkMode ? "bg-gray-800 text-white" : "bg-white text-black"}
        rounded-xl
        md:shadow-lg
        p-5
        space-y-5
      `}
      >
        <div>
          <button
            className={`
            float-right
            hover:bg-blue-500
            rounded-xl
            ${isDarkMode ? "hover:bg-blue-600" : "hover:bg-blue-500"}
          `}
            onClick={() => scrollCards(200)}
          >
            <IoChevronForwardCircleOutline size={20} />
          </button>
          <button
            className={`
            float-left
            hover:bg-blue-500
            rounded-xl
            ${isDarkMode ? "hover:bg-blue-600" : "hover:bg-blue-500"}
          `}
            onClick={() => scrollCards(-200)}
          >
            <IoChevronBackCircleOutline size={20} />
          </button>
        </div>

        <div
          ref={QuickTransferSection}
          className={`
          flex
          max-w-[300px]
          space-x-5
          overflow-x-auto
          [&::-webkit-scrollbar]:hidden
          [-ms-overflow-style:none]
          [scrollbar-width:none]
        `}
        >
          {loading
            ? [1, 2, 3, 4].map((index) => <Shimmer key={index} />)
            : quickTransfer.map((account) => (
                <Profile
                  key={account.id}
                  image="/images/avatar2.svg"
                  name={account.name}
                  job="Director"
                  isSelected={selectedProfile?.id === account.id}
                  onClick={() => handleProfileSelect(account)}
                />
              ))}
        </div>

        <div className="flex space-x-10 h-[40px] items-center">
          <button
            className={`
              ${
                isDarkMode
                  ? "bg-[#3B6EE2] hover:bg-[#2A56B8]"
                  : "bg-[#1814F3] hover:bg-[#0F0DC7]"
              }
              text-white
              rounded-full
              px-4
              h-[40px]
              ml-2
              flex
              items-center
              space-x-2
              transition-all duration-300 ease-in-out
            `}
            onClick={handleModalToggle}
            disabled={!selectedProfile}
          >
            <p>Transfer</p>
            <PiTelegramLogoLight />
          </button>
          {isModalOpen && (
            <div onClick={handleModalToggle}>
              <div
                className="relative bg-white p-6 rounded-lg shadow-lg max-w-md w-full"
                onClick={(e) => e.stopPropagation()}
              >
                <ModalTrans
                  isOpen={isModalOpen}
                  onClose={handleModalToggle}
                  reciverUserName={selectedProfile.username}
                />
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};
