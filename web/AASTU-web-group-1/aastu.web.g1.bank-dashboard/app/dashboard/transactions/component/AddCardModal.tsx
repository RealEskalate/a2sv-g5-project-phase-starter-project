import React from "react";

import { IoClose } from "react-icons/io5";
import InputForm from "../../credit-cards/components/InputForm";
import { useUser } from "@/contexts/UserContext";

interface Props {
  isOpen: boolean;
  onClose: () => void;
}

export const AddCardModal = ({ isOpen, onClose }: Props) => {
    const {isDarkMode} = useUser();
  if (!isOpen) return null;

  return (
    <div
      className={`fixed inset-0 z-50 flex items-center justify-center ${
        isDarkMode
          ? "bg-gray-900/70 backdrop-blur-md"
          : "bg-black bg-opacity-50 "
      } `}
    >
      /
      <div
        className={`relative w-full max-w-lg p-8 rounded-3xl shadow-xl transition-transform transform ${
          isDarkMode
            ? "bg-black bg-opacity-50"
            : "bg-gradient-to-r from-white via-gray-100 to-white"
        } `}
      >
        <button
          onClick={onClose}
          className="absolute top-4 right-4 text-gray-600 hover:text-gray-900"
        >
          <IoClose size={24} />
        </button>

        <div className="text-center">
          {" "}
          <h2 className="text-2xl font-bold tracking-wide text-transparent bg-clip-text bg-gradient-to-r from-sky-500 to-indigo-500">
            Add New Card
          </h2>
        </div>

        <InputForm />
      </div>
    </div>
  );
};
